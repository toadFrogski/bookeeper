package handlers

import (
	"gg/database"
	"gg/middlewares"
	"gg/modules/book"

	"github.com/gin-gonic/gin"
)

func GetBooksRoutes(r gin.IRouter) gin.IRouter {
	bookAPI := book.Wire(database.DB)

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", bookAPI.GetAllBooks)
		bookRouter.GET(":bookID", bookAPI.GetBook)
		bookRouter.POST("save", middlewares.JwtAuthMiddleware(), bookAPI.SaveBook)
		bookRouter.DELETE(":bookID", middlewares.JwtAuthMiddleware(), bookAPI.DeleteBookByID)
	}

	return r
}
