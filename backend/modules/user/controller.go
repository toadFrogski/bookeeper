package user

import (
	"gg/domain"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserSvc domain.IUserService
}

// Register godoc
// @Summary Register new  user
// @Accept  json
// @Produce json
// @Param   register body RegisterUserForm true "Register form"
// @Success 200 {object} dto.Response[AuthResponse]
// @Failure 400 {object} dto.Response[any]
// @Router /register [post]
func (uc UserController) Register(c *gin.Context) {
	uc.UserSvc.Register(c)
}

func (uc UserController) Login(c *gin.Context) {
	uc.UserSvc.Login(c)
}

func (uc UserController) GetUserInfo(c *gin.Context) {
	uc.UserSvc.GetUserInfo(c)
}
