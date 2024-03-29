package handlers

import (
	"bookeeper/database"
	"bookeeper/middlewares"
	"bookeeper/modules/book"
	"bookeeper/utils/constants"

	"github.com/gin-gonic/gin"
)

func GetBooksRoutes(r gin.IRouter) gin.IRouter {
	bookAPI := book.Wire(database.DB)

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", bookAPI.GetBookList)
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
