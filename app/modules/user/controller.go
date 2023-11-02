package user

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	svc domain.UserService
}

func (uc UserControllerImpl) CreateUser(c *gin.Context) {
	uc.svc.CreateUser(c)
}
