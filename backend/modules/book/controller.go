package book

import (
	"bookeeper/domain"
	"bookeeper/utils/constants"
	_ "bookeeper/utils/dto"
	_ "bookeeper/utils/paginator"
	"bookeeper/utils/panic"
	"strconv"

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
// @Param photo formData file true "Image to be uploaded"
// @Param name formData string true "Name of book"
// @Param author formData string true "Book author"
// @Param description formData string true "Description of book"
// @Success 200 {object} AnyResponse
// @Failure 400 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /book/save [post]
func (bc BookController) SaveBook(c *gin.Context) {
	bc.BookSvc.SaveBook(c)
}

// UpdateBook godoc
//
// @Summary Update book
// @Tags book
// @Accept mpfd
// @Produce json
// @Param book_id path int true "Book ID"
// @Param photo formData file false "Image to be uploaded"
// @Param name formData string false "Name of book"
// @Param author formData string false "Book author"
// @Param description formData string false "Description of book"
// @Success 200 {object} AnyResponse
// @Failure 400 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /book/{book_id} [post]
func (bc BookController) UpdateBook(c *gin.Context) {
	bc.BookSvc.UpdateBook(c)
}

// GetBook godoc
// @Summary Get book by ID
// @Tags book
// @Param book_id path int true "Book ID"
// @Success 200 {object} BookResponse
// @Failure 400 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /book/{book_id} [get]
func (bc BookController) GetBook(c *gin.Context) {
	bc.BookSvc.GetBook(c)
}

// DeleteBook godoc
// @Summary Delete book by ID
// @Tags book
// @Param book_id path int true "Book ID"
// @Success 200 {object} AnyResponse
// @Failure 400 {object} AnyResponse
// @Router /book/{book_id} [delete]
func (bc BookController) DeleteBookByID(c *gin.Context) {
	bc.BookSvc.DeleteBookByID(c)
}

// GetAllUserBooks godoc
// @Summary Get all user books by user ID
// @Tags book
// @Param user_id path int true "User ID"
// @Success 200 {object} BooksResponse
// @Failure 400 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /book/user/{user_id} [get]
func (bc BookController) GetBooksByUserID(c *gin.Context) {
	param, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	userID := uint(param)

	bc.BookSvc.GetBooksByUserID(c, userID)
}

// GetAllUserBooks godoc
// @Summary Get all user books by user ID
// @Tags book
// @Success 200 {object} BooksResponse
// @Failure 400 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /book/user/me [get]
func (bc BookController) GetBooksBySelf(c *gin.Context) {
	var user *domain.User

	vars, exist := c.Get("user")
	if !exist {
		panic.PanicException(constants.InternalError)
	}
	user = vars.(*domain.User)

	bc.BookSvc.GetBooksByUserID(c, user.ID)
}
