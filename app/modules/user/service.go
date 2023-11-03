package user

import (
	"gg/domain"
	"gg/utils/constants"
	"gg/utils/dto"
	"gg/utils/panic"
	"gg/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserServiceImpl struct {
	repo domain.UserRepository
}

type CreateUserForm struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (us UserServiceImpl) Register(c *gin.Context) {
	defer panic.PanicHandler(c)

	var user domain.User
	var createUserForm CreateUserForm

	if err := c.ShouldBind(&createUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user.Email = createUserForm.Email
	user.Username = createUserForm.Username
	user.Password = createUserForm.Password

	if err := us.repo.CreateUser(&user); err != nil {
		panic.PanicException(constants.InternalError)
	}

	accessToken, err := token.GenerateToken(uint(user.ID))
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusCreated,
		dto.BuildResponse[map[string]string](constants.Success, map[string]string{"token": accessToken}),
	)
}

func (us UserServiceImpl) Login(c *gin.Context) {

}
