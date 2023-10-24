package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooksRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"aa": "b"})
		})
	}

	return r
}
