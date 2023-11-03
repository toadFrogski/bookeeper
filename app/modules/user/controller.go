package user

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	svc domain.UserService
}

func (uc UserControllerImpl) Register(c *gin.Context) {
	uc.svc.Register(c)
}

func (uc UserControllerImpl) Login(c *gin.Context) {
	uc.svc.Login(c)
}
