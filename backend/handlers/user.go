package handlers

import (
	"gg/database"
	"gg/modules/user"

	"github.com/gin-gonic/gin"
)

func GetUserRoutes(r gin.IRouter) gin.IRouter {
	userAPI := user.Wire(database.DB)
	r.GET("/profile", userAPI.GetUserInfo)
	return r
}
