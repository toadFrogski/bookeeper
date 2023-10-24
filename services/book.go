package service

import (
	model "gg/models"
	repository "gg/repositories"
	"gg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookService interface {
	GetAllBooks(c *gin.Context)
}

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func ProvideBookService(bookRepository repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{bookRepository: bookRepository}
}

func (b BookServiceImpl) GetAllBooks(c *gin.Context) {
	defer utils.PanicHandler(c)
	data, err := b.bookRepository.GetAllBooks()
	if err != nil {
		utils.PanicException(utils.DataNotFound)
	}

	c.JSON(http.StatusOK, utils.BuildResponse[[]model.Book](utils.Success, data))
}
