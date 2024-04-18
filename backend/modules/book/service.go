package book

import (
	"bookeeper/domain"
	"bookeeper/utils/constants"
	"bookeeper/utils/dto"
	p "bookeeper/utils/paginator"
	"bookeeper/utils/panic"
	"errors"
	"net/http"
	"os"
	"path/filepath"
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

func (bs BookService) UpdateBook(c *gin.Context) {
	var form OptionalBookForm
	var book *domain.Book

	bookID := c.Param("book_id")

	if err := c.ShouldBind(&form); err != nil {
		panic.PanicWithMessage(constants.InvalidRequest, err.Error())
	}

	book, err := bs.BookRepo.GetBookByID(bookID)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	if form.Name != book.Name {
		book.Name = form.Name
	}
	if form.Description != book.Description {
		book.Description = form.Description
	}
	if form.Author != book.Author {
		book.Author = form.Author
	}
	if form.Photo != nil {
		media := os.Getenv("MEDIA_DIRECTORY")
		fpath := filepath.Join(media, book.Photo)
		c.SaveUploadedFile(form.Photo, fpath)
	}

	if err := bs.BookRepo.UpdateBook(book); err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[any](constants.Success, nil))
}

func (bs BookService) SaveBook(c *gin.Context) {
	var saveBookForm BookForm
	var user *domain.User

	if err := c.ShouldBind(&saveBookForm); err != nil {
		panic.PanicWithMessage(constants.InvalidRequest, err.Error())
	}

	if err := validateBookForm(&saveBookForm); err != nil {
		panic.PanicWithMessage(constants.InvalidRequest, err.Error())
	}

	claims, exist := c.Get("user")
	if !exist {
		panic.PanicException(constants.InternalError)
	}
	user = claims.(*domain.User)

	fpath := filepath.Join(time.Now().Format("20060102"),
		strconv.FormatUint(uint64(user.ID), 10),
		strconv.FormatInt(time.Now().Unix(), 10),
	)

	book := domain.Book{
		Name:        saveBookForm.Name,
		Author:      saveBookForm.Author,
		Description: saveBookForm.Description,
		UserID:      user.ID,
		Photo:       fpath,
	}

	if err := bs.BookRepo.CreateBook(&book); err != nil {
		panic.PanicException(constants.InternalError)
	}

	media := os.Getenv("MEDIA_DIRECTORY")

	fpath = filepath.Join(media, fpath)
	c.SaveUploadedFile(saveBookForm.Photo, fpath)
}

func (bs BookService) GetBook(c *gin.Context) {
	var book *domain.Book

	book, err := bs.BookRepo.GetBookByID(c.Param("bookID"))
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[domain.Book](constants.Success, *book))
}

func (bs BookService) DeleteBookByID(c *gin.Context) {
	var book *domain.Book
	var user *domain.User

	media := os.Getenv("MEDIA_DIRECTORY")

	vars, exist := c.Get("user")
	if !exist {
		panic.PanicException(constants.InternalError)
	}
	user = vars.(*domain.User)

	book, err := bs.BookRepo.GetBookByID(c.Param("bookID"))
	if err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	canDelete := user.ID == book.UserID
	for _, userRole := range user.Roles {
		canDelete = canDelete || userRole.Name == constants.Admin
	}

	if !canDelete {
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

func (bs BookService) GetBooksByUserID(c *gin.Context, userID uint) {
	var booksPointers []*domain.Book
	var books []domain.Book

	booksPointers, err := bs.BookRepo.GetAllUserBooks(userID)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	if len(booksPointers) == 0 {
		panic.PanicException(constants.DataNotFound)
	}

	for _, bookPtr := range booksPointers {
		books = append(books, *bookPtr)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[[]domain.Book](constants.Success, books))
}
