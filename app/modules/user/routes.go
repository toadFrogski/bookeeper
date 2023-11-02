package user

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func GetUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	userAPI := Wire(db)

	userRouter := r.Group("user")
	{
		userRouter.POST("create", userAPI.CreateUser)
	}
}
