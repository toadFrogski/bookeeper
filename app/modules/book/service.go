package book

import (
	"gg/domain"
	"gg/utils/constants"
	"gg/utils/dto"
	"gg/utils/panic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookServiceImpl struct {
	bookRepo domain.BookRepository
}

func (bs BookServiceImpl) GetAllBooks(c *gin.Context) {
	defer panic.PanicHandler(c)
	data, err := bs.bookRepo.GetAllBooks()
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[[]domain.Book](constants.Success, data))
}

func (bs BookServiceImpl) SaveBook(c *gin.Context) {
	defer panic.PanicHandler(c)
}
