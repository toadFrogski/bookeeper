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

type UserService struct {
	userRepo domain.IUserRepository
}

func (us UserService) Register(c *gin.Context) {
	var user domain.User
	var RegisterUserForm RegisterUserForm

	if err := c.ShouldBind(&RegisterUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user.Email = RegisterUserForm.Email
	user.Username = RegisterUserForm.Username
	user.Password = RegisterUserForm.Password
	user.Roles = append(user.Roles, &domain.Role{Name: string(constants.User)})

	if err := us.userRepo.CreateUser(&user); err != nil {
		panic.PanicException(constants.InternalError)
	}

	accessToken, err := token.GenerateToken(user)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusCreated,
		dto.BuildResponse[map[string]string](constants.Success, map[string]string{"token": accessToken}),
	)
}

func (us UserService) Login(c *gin.Context) {
	var user domain.User
	var loginUserForm LoginUserForm

	if err := c.ShouldBind(&loginUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user, err := us.userRepo.GetUserByEmail(loginUserForm.Email)
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	if err := user.ValidatePassword(loginUserForm.Password); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	accessToken, err := token.GenerateToken(user)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusOK,
		dto.BuildResponse[map[string]string](constants.Success, map[string]string{"token": accessToken}),
	)
}
