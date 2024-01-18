package domain

import (
	"github.com/gin-gonic/gin"
)

type (
	Book struct {
		ID          uint   `gorm:"primarykey"`
		Name        string `gorm:"column:name" json:"name"`
		Author      string `gorm:"column:author" json:"author"`
		Description string `gorm:"column:description; text" json:"description"`
		Photo       string `gorm:"column:photo" json:"photo"`
		UserID      uint
		User        User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	}

	IBookController interface {
		GetAllBooks(c *gin.Context)
		GetBook(c *gin.Context)
		SaveBook(c *gin.Context)
		DeleteBookByID(c *gin.Context)
	}

	IBookService interface {
		GetAllBooks(c *gin.Context)
		GetBook(c *gin.Context)
		SaveBook(c *gin.Context)
		DeleteBookByID(c *gin.Context)
	}

	IBookRepository interface {
		GetAllBooks() ([]Book, error)
		CreateBook(book *Book) error
		DeleteBook(book *Book) error
		DeleteBookByID(ID string) error
		GetUserBookByID(ID string) (*Book, error)
	}
)
