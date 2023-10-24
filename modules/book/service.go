package book

import (
	"gg/domain"
	"gg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookServiceImpl struct {
	bookRepository domain.BookRepository
}

func (b BookServiceImpl) GetAllBooks(c *gin.Context) {
	defer utils.PanicHandler(c)
	data, err := b.bookRepository.GetAllBooks()
	if err != nil {
		utils.PanicException(utils.DataNotFound)
	}

	c.JSON(http.StatusOK, utils.BuildResponse[[]domain.Book](utils.Success, data))
}
