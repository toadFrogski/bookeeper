package book

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type BookControllerAPI struct {
	svc domain.BookService
}

func (b BookControllerAPI) GetAllBooks(c *gin.Context) {
	b.svc.GetAllBooks(c)
}
