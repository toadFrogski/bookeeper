package book

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooksRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	bookAPI := Wire(db)

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", bookAPI.GetAllBooks)
	}

	return r
}
