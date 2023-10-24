package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	BaseModel struct {
		CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
		UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
		DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
	}

	Book struct {
		ID          int    `gorm:"column:id; primary_key; not null; autoIncrement" json:"id"`
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
