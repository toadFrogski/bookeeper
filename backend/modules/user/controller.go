package user

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserSvc domain.IUserService
}

func (uc UserController) Register(c *gin.Context) {
	uc.UserSvc.Register(c)
}

func (uc UserController) Login(c *gin.Context) {
	uc.UserSvc.Login(c)
}
