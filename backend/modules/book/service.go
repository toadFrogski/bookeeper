package book

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gg/domain"
	"gg/utils/constants"
	"gg/utils/dto"
	p "gg/utils/paginator"
	"gg/utils/panic"
	"gg/utils/token"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookService struct {
	BookRepo domain.IBookRepository
}

func (bs BookService) GetBookList(c *gin.Context) {
	var books []*domain.Book
	var paginator p.Paginator[[]*domain.Book]

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	sort, ok := c.GetQuery("sort")
	if !ok {
		sort = "id desc"
	}

	paginator.Page = page
	paginator.Limit = limit
	paginator.Sort = sort

	books, err = bs.BookRepo.GetBookList(paginator)
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}
	paginator.Rows = books

	c.JSON(http.StatusOK, dto.BuildResponse[p.Paginator[[]*domain.Book]](constants.Success, paginator))
}

func (bs BookService) SaveBook(c *gin.Context) {
	var saveBookForm SaveBookForm
	var user token.Claims

	media := os.Getenv("MEDIA_DIRECTORY")

	if err := c.ShouldBind(&saveBookForm); err != nil {
		panic.PanicWithMessage(constants.InvalidRequest, err.Error())
	}

	if err := validateSaveBookForm(&saveBookForm); err != nil {
		panic.PanicWithMessage(constants.InvalidRequest, err.Error())
	}

	claims, exist := c.Get("user")
	if !exist {
		panic.PanicException(constants.InternalError)
	}
	user = claims.(token.Claims)

	fpath := filepath.Join(time.Now().Format("20060102"),
		string(saveBookForm.UserID)+
			string(time.Now().Unix())+
			base64.StdEncoding.EncodeToString([]byte(saveBookForm.Photo.Filename)),
	)

	book := domain.Book{
		Name:        saveBookForm.Name,
		Author:      saveBookForm.Author,
		Description: saveBookForm.Description,
		UserID:      uint(user.UserID),
		Photo:       fpath,
	}

	if err := bs.BookRepo.CreateBook(&book); err != nil {
		panic.PanicException(constants.InternalError)
	}

	fpath = filepath.Join(media, fpath)
	c.SaveUploadedFile(saveBookForm.Photo, fpath)
}

func (bs BookService) GetBook(c *gin.Context) {
	var book *domain.Book

	book, err := bs.BookRepo.GetUserBookByID(c.Param("bookID"))
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[domain.Book](constants.Success, *book))
}

func (bs BookService) DeleteBookByID(c *gin.Context) {
	var book *domain.Book
	var user token.Claims

	media := os.Getenv("MEDIA_DIRECTORY")

	claims, exist := c.Get("user")
	if !exist {
		panic.PanicException(constants.InternalError)
	}
	user = claims.(token.Claims)

	book, err := bs.BookRepo.GetUserBookByID(c.Param("bookID"))
	if err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	if canDelete := uint(user.UserID) == book.UserID || slices.Contains(user.Roles, constants.Admin); !canDelete {
		panic.PanicException(constants.PermissionDenied)
	}

	bs.BookRepo.DeleteBook(book)

	fpath := filepath.Join(media, book.Photo)
	if _, err := os.Stat(fpath); errors.Is(err, os.ErrNotExist) {
		// do nothing if file does not exist
	} else if err := os.Remove(fpath); err != nil {
		panic.PanicException(constants.InternalError)
	}
}

func validateSaveBookForm(form *SaveBookForm) error {
	fileType := form.Photo.Header.Get("Content-Type")
	if ok, _ := regexp.MatchString("image/*", fileType); !ok {
		return fmt.Errorf("Uploaded file is not image")
	}

	return nil
}
