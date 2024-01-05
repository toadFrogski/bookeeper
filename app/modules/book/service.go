package book

import (
	"fmt"
	"gg/domain"
	"gg/utils/constants"
	"gg/utils/dto"
	"gg/utils/panic"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type BookService struct {
	bookRepo domain.IBookRepository
}

func (bs BookService) GetAllBooks(c *gin.Context) {
	data, err := bs.bookRepo.GetAllBooks()
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[[]domain.Book](constants.Success, data))
}

func (bs BookService) SaveBook(c *gin.Context) {
	cwd := os.Getenv("MEDIA_DIRECTORY")

	var saveBookForm SaveBookForm

	if err := c.ShouldBind(&saveBookForm); err != nil {
		panic.PanicWithMessage(constants.InvalidRequest, err.Error())
	}

	if err := validateSaveBookForm(&saveBookForm); err != nil {
		panic.PanicWithMessage(constants.InvalidRequest, err.Error())
	}

	fpath := filepath.Join(time.Now().Format("20060102"),
		string(saveBookForm.UserID)+string(time.Now().Unix())+saveBookForm.Photo.Filename,
	)

	book := domain.Book{
		Name:        saveBookForm.Name,
		Author:      saveBookForm.Author,
		Description: saveBookForm.Description,
		UserID:      c.Keys["user_id"].(uint),
		Photo:       fpath,
	}

	if err := bs.bookRepo.CreateBook(&book); err != nil {
		panic.PanicException(constants.InternalError)
	}

	fpath = filepath.Join(cwd, fpath)
	c.SaveUploadedFile(saveBookForm.Photo, fpath)
}

func (bs BookService) GetBook(c *gin.Context) {
	book, err := bs.bookRepo.GetUserBookByID(c.Param("bookID"))
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}
	c.JSON(http.StatusOK, dto.BuildResponse[domain.Book](constants.Success, *book))
}

func (bs BookService) DeleteBookByID(c *gin.Context) {
	bs.bookRepo.DeleteBookByID(c.Param("bookID"))
}

func validateSaveBookForm(form *SaveBookForm) error {
	test := form.Photo.Header.Get("Content-Type")
	if ok, _ := regexp.MatchString("image/*", test); !ok {
		return fmt.Errorf("Uploaded file is not image")
	}

	return nil
}
