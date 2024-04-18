package domain

import (
	"bookeeper/utils/paginator"

	"github.com/gin-gonic/gin"
)

type (
	Book struct {
		ID          uint   `gorm:"primarykey" json:"ID"`
		Name        string `gorm:"column:name" json:"name"`
		Author      string `gorm:"column:author" json:"author"`
		Description string `gorm:"column:description; text" json:"description"`
		Photo       string `gorm:"column:photo" json:"photo"`
		UserID      uint   `json:"-"`
		User        *User  `json:"user,omitempty"`
	} // @name Book

	IBookController interface {
		GetBook(c *gin.Context)
		GetBookList(c *gin.Context)
		SaveBook(c *gin.Context)
		DeleteBookByID(c *gin.Context)
		GetBooksByUserID(c *gin.Context)
		GetBooksBySelf(c *gin.Context)
	}

	IBookService interface {
		GetBook(c *gin.Context)
		GetBookList(c *gin.Context)
		SaveBook(c *gin.Context)
		UpdateBook(c *gin.Context)
		DeleteBookByID(c *gin.Context)
		GetBooksByUserID(c *gin.Context, userID uint)
	}

	IBookRepository interface {
		CreateBook(book *Book) error
		UpdateBook(book *Book) error
		GetBookList(paginator paginator.Paginator[[]*Book]) ([]*Book, error)
		DeleteBook(book *Book) error
		DeleteBookByID(ID string) error
		GetBookByID(ID string) (*Book, error)
		GetAllUserBooks(ID uint) ([]*Book, error)
	}
)
