package user

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func GetUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	userAPI := Wire(db)

	r.POST("register", userAPI.Register)
	r.POST("login", userAPI.Login)

	// userRouter := r.Group("user")
	// userRouter.Use(middlewares.JwtAuthMiddleware())
	// {
	// 	userRouter.POST("create", userAPI.CreateUser)
	// }
}
