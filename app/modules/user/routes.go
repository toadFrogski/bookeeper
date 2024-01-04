package user

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func GetUserRoutes(r gin.IRouter, db *gorm.DB) gin.IRouter {
	// @TODO: Write user profile
	// userRouter := r.Group("user")
	// userRouter.Use(middlewares.JwtAuthMiddleware())
	// {
	// 	userRouter.POST("create", userAPI.CreateUser)
	// }
	return r
}

func GetAuthRoutes(r gin.IRouter, db *gorm.DB) gin.IRouter {
	userAPI := Wire(db)

	r.POST("register", userAPI.Register)
	r.POST("login", userAPI.Login)

	return r
}
