package book

import (
	"gg/domain"
	"gg/utils/constants"
	"gg/utils/dto"
	"gg/utils/panic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookService struct {
	bookRepo domain.IBookRepository
}

func (bs BookService) GetAllBooks(c *gin.Context) {
	defer panic.PanicHandler(c)
	data, err := bs.bookRepo.GetAllBooks()
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[[]domain.Book](constants.Success, data))
}

func (bs BookService) SaveBook(c *gin.Context) {
	defer panic.PanicHandler(c)
}
