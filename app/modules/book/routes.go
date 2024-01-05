package book

import (
	"gg/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooksRoutes(r gin.IRouter, db *gorm.DB) gin.IRouter {

	bookAPI := Wire(db)

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", bookAPI.GetAllBooks)
		bookRouter.GET(":bookID", bookAPI.GetBook)
		bookRouter.POST("save", middlewares.JwtAuthMiddleware(), bookAPI.SaveBook)
	}

	return r
}
