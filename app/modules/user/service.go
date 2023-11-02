package user

import (
	"gg/domain"
	"gg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserServiceImpl struct {
	repo domain.UserRepository
}

type CreateUserForm struct {
	Email    string `form:"email" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (us UserServiceImpl) CreateUser(c *gin.Context) {
	defer utils.PanicHandler(c)

	var user domain.User
	var createUserForm CreateUserForm

	if err := c.ShouldBind(&createUserForm); err != nil {
		utils.PanicException(utils.InvalidRequest)
	}

	user.Email = createUserForm.Email
	user.Username = createUserForm.Username
	user.Password = createUserForm.Password

	if err := us.repo.CreateUser(&user); err != nil {
		utils.PanicException(utils.InternalError)
	}

	c.JSON(http.StatusCreated, utils.BuildResponse[any](utils.Success, nil))
}
