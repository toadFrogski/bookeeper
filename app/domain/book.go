package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Book struct {
		gorm.Model
		Name        string `gorm:"column:name" json:"name"`
		Author      string `gorm:"column:author" json:"author"`
		Description string `gorm:"colunm:description; text" json:"description"`
		Photo       string `gorm:"column:photo" json:"photo"`
		UserID      uint
	}

	BookController interface {
		GetAllBooks(c *gin.Context)
		SaveBook(c *gin.Context)
	}

	BookService interface {
		GetAllBooks(c *gin.Context)
		SaveBook(c *gin.Context)
	}

	BookRepository interface {
		GetAllBooks() ([]Book, error)
		SaveBook(book *Book) error
	}
)
