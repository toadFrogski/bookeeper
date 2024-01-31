package handlers

import (
	"bookeeper/database"
	"bookeeper/middlewares"
	"bookeeper/modules/user"
	"bookeeper/utils/constants"

	"github.com/gin-gonic/gin"
)

func GetUserRoutes(r gin.IRouter) gin.IRouter {
	userAPI := user.Wire(database.DB)
	r.GET("/profile",
		middlewares.RoleAccessMiddleware([]constants.Role{constants.Admin, constants.User}),
		userAPI.GetUserInfo)
	r.GET("/profile/:userID",
		middlewares.RoleAccessMiddleware([]constants.Role{constants.Admin}),
		userAPI.GetUserInfoByID)

	return r
}
