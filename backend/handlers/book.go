package handlers

import (
	"bookeeper/database"
	m "bookeeper/middlewares"
	"bookeeper/modules/book"
	c "bookeeper/utils/constants"

	"github.com/gin-gonic/gin"
)

func GetBooksRoutes(r gin.IRouter) gin.IRouter {
	bookAPI := book.Wire(database.DB)

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", bookAPI.GetBookList)
		bookRouter.GET(":bookID", bookAPI.GetBook)
		bookRouter.POST("save",
			m.RoleAccessMiddleware([]c.Role{c.Admin, c.User}),
			bookAPI.SaveBook)
		bookRouter.DELETE(":bookID",
			m.RoleAccessMiddleware([]c.Role{c.Admin, c.User}),
			bookAPI.DeleteBookByID)
	}

	return r
}
