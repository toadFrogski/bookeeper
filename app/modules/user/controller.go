package user

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userSvc domain.UserService
}

func (uc UserControllerImpl) Register(c *gin.Context) {
	uc.userSvc.Register(c)
}

func (uc UserControllerImpl) Login(c *gin.Context) {
	uc.userSvc.Login(c)
}
