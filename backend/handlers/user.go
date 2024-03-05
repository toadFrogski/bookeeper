package handlers

import (
	"bookeeper/database"
	m "bookeeper/middlewares"
	"bookeeper/modules/user"
	c "bookeeper/utils/constants"

	"github.com/gin-gonic/gin"
)

func GetUserRoutes(r gin.IRouter) gin.IRouter {
	userAPI := user.Wire(database.DB)
	r.GET("/profile",
		m.RoleAccessMiddleware([]c.Role{c.Admin, c.User}),
		userAPI.GetUserInfo)
	r.GET("/profile/:userID",
		m.RoleAccessMiddleware([]c.Role{c.Admin}),
		userAPI.GetUserInfoByID)

	return r
}
