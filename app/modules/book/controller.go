package book

import (
	"gg/domain"
	_ "gg/utils/dto"

	"github.com/gin-gonic/gin"
)

type BookControllerImpl struct {
	bookSvc domain.BookService
}

// GetAllBooks godoc
//
// @Summary Get all books
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response[domain.Book]
// @Failure 400 {object} dto.Response[any]
// @Router /book/ [get]
func (bc BookControllerImpl) GetAllBooks(c *gin.Context) {
	bc.bookSvc.GetAllBooks(c)
}

func (bc BookControllerImpl) SaveBook(c *gin.Context) {
	bc.bookSvc.SaveBook(c)
}
