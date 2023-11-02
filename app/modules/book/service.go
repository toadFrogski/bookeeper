package book

import (
	"gg/domain"
	"gg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookServiceImpl struct {
	repo domain.BookRepository
}

func (bs BookServiceImpl) GetAllBooks(c *gin.Context) {
	defer utils.PanicHandler(c)
	data, err := bs.repo.GetAllBooks()
	if err != nil {
		utils.PanicException(utils.DataNotFound)
	}

	c.JSON(http.StatusOK, utils.BuildResponse[[]domain.Book](utils.Success, data))
}
