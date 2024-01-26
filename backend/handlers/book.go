package handlers

import (
	"gg/database"
	"gg/middlewares"
	"gg/modules/book"
	"gg/utils/constants"

	"github.com/gin-gonic/gin"
)

func GetBooksRoutes(r gin.IRouter) gin.IRouter {
	bookAPI := book.Wire(database.DB)

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", bookAPI.GetAllBooks)
		bookRouter.GET(":bookID", bookAPI.GetBook)
		bookRouter.POST("save",
			middlewares.RoleAccessMiddleware([]constants.Role{constants.Admin, constants.User}),
			bookAPI.SaveBook)
		bookRouter.DELETE(":bookID",
			middlewares.RoleAccessMiddleware([]constants.Role{constants.Admin, constants.User}),
			bookAPI.DeleteBookByID)
	}

	return r
}
