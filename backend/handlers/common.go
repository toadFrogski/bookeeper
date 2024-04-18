package handlers

import (
	"bookeeper/database"
	"bookeeper/modules/user"

	"github.com/gin-gonic/gin"
)

func GetAuthRoutes(r gin.IRouter) gin.IRouter {
	userAPI := user.Wire(database.DB)

	r.POST("/register", userAPI.Register)
	r.POST("/login", userAPI.Login)
	r.POST("/refresh", userAPI.Refresh)

	return r
}
