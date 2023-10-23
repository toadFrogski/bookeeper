package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooksRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	// bookController := NewBookController()

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"aa": "b"})
		})
	}

	return r
}
