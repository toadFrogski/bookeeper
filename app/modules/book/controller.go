package book

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type BookControllerImpl struct {
	svc domain.BookService
}

func (bc BookControllerImpl) GetAllBooks(c *gin.Context) {
	bc.svc.GetAllBooks(c)
}
