package user

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc domain.IUserService
}

func (uc UserController) Register(c *gin.Context) {
	uc.userSvc.Register(c)
}

func (uc UserController) Login(c *gin.Context) {
	uc.userSvc.Login(c)
}
