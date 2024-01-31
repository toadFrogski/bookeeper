package book

import (
	"bookeeper/domain"
	_ "bookeeper/utils/dto"
	_ "bookeeper/utils/paginator"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookSvc domain.IBookService
}

// GetBookList godoc
//
// @Summary Get book list
// @Tags book
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit books per page"
// @Param sort query string false "Sorting"
// @Success 200 {object} BookListResponse
// @Failure 400 {object} AnyResponse
// @Router /book/ [get]
func (bc BookController) GetBookList(c *gin.Context) {
	bc.BookSvc.GetBookList(c)
}

// SaveBook godoc
//
// @Summary Save book
// @Tags book
// @Accept mpfd
// @Produce json
// @Param image formData file true "Image to be uploaded"
// @Param name formData string true "Name of book"
// @Success 200 {object} AnyResponse
// @Failude 400 {object} AnyResponse
// @Failude 500 {object} AnyResponse
// @Router /book/save [post]
func (bc BookController) SaveBook(c *gin.Context) {
	bc.BookSvc.SaveBook(c)
}

// GetBook godoc
// @Summary Get book by ID
// @Tags book
// @Param book_id path int true "Book ID"
// @Success 200 {object} BookResponse
// @Failude 400 {object} AnyResponse
// @Failude 500 {object} AnyResponse
// @Router /book/{book_id} [get]
func (bc BookController) GetBook(c *gin.Context) {
	bc.BookSvc.GetBook(c)
}

// DeleteBook godoc
// @Summmary Delete book by ID
// @Tags book
// @Param book_id path int true "Book ID"
// @Success 200 {object} AnyResponse
// @Failure 400 {object} AnyResponse
// @Router /book/{book_id} [delete]
func (bc BookController) DeleteBookByID(c *gin.Context) {
	bc.BookSvc.DeleteBookByID(c)
}
