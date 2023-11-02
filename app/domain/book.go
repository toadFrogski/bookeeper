package domain

import (
	"github.com/gin-gonic/gin"
)

type (
	Book struct {
		Name        string `gorm:"column:name" json:"name"`
		Author      string `gorm:"column:author" json:"author"`
		Description string `gorm:"colunm:description; text" json:"description"`
		BaseModel
	}

	BookRepository interface {
		GetAllBooks() ([]Book, error)
	}

	BookService interface {
		GetAllBooks(c *gin.Context)
	}

	BookController interface {
		GetAllBooks(c *gin.Context)
	}
)
