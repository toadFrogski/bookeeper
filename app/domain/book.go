package domain

import (
	"github.com/gin-gonic/gin"
)

type (
	Book struct {
		ID          uint   `gorm:"primarykey"`
		Name        string `gorm:"column:name" json:"name"`
		Author      string `gorm:"column:author" json:"author"`
		Description string `gorm:"colunm:description; text" json:"description"`
		Photo       string `gorm:"column:photo" json:"photo"`
		UserID      uint
	}

	IBookController interface {
		GetAllBooks(c *gin.Context)
		SaveBook(c *gin.Context)
	}

	IBookService interface {
		GetAllBooks(c *gin.Context)
		SaveBook(c *gin.Context)
	}

	IBookRepository interface {
		GetAllBooks() ([]Book, error)
		SaveBook(book *Book) error
	}
)
