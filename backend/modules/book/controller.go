package book

import (
	"gg/domain"
	_ "gg/utils/dto"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookSvc domain.IBookService
}

// GetAllBooks godoc
//
// @Summary Get all books
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response[domain.Book]
// @Failure 400 {object} dto.Response[any]
// @Router /book/ [get]
func (bc BookController) GetAllBooks(c *gin.Context) {
	bc.BookSvc.GetAllBooks(c)
}

// SaveBook godoc
//
// @Summary Save book
// @Accept mpfd
// @Produce json
// @Param image formData file true "Image to be uploaded"
// @Param name formData string true "Name of book"
// @Success 200 {object} dto.Response[any]
// @Failude 400 {object} dto.Response[any]
// @Failude 500 {object} dto.Response[any]
// @Router /book/save [post]
func (bc BookController) SaveBook(c *gin.Context) {
	bc.BookSvc.SaveBook(c)
}

// GetBook godoc
// @Summary Get book by ID
// @Param book_id path int true "Book ID"
// @Success 200 {object} dto.Response[domain.Book]
// @Failude 400 {object} dto.Response[any]
// @Failude 500 {object} dto.Response[any]
// @Router /book/{book_id} [get]
func (bc BookController) GetBook(c *gin.Context) {
	bc.BookSvc.GetBook(c)
}

// DeleteBook godoc
// @Summmary Delete book by ID
// @Param book_id path int true "Book ID"
// @Success 200 {object} dto.Response[any]
// @Failure 400 {object} dto.Response[any]
// @Router /book/{book_id} [delete]
func (bc BookController) DeleteBookByID(c *gin.Context) {
	bc.BookSvc.DeleteBookByID(c)
}
