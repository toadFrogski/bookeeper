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
	repo domain.BookRepository
}

func (bs BookServiceImpl) GetAllBooks(c *gin.Context) {
	defer panic.PanicHandler(c)
	data, err := bs.repo.GetAllBooks()
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[[]domain.Book](constants.Success, data))
}
