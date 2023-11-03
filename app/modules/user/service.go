package user

import (
	"crypto/sha256"
	"fmt"
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

type RegisterUserForm struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (us UserServiceImpl) Register(c *gin.Context) {
	defer panic.PanicHandler(c)

	var user domain.User
	var RegisterUserForm RegisterUserForm

	if err := c.ShouldBind(&RegisterUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user.Email = RegisterUserForm.Email
	user.Username = RegisterUserForm.Username
	user.Password = RegisterUserForm.Password

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
	defer panic.PanicHandler(c)

	var user domain.User
	var loginUserForm LoginUserForm

	if err := c.ShouldBind(&loginUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user, err := us.repo.GetUserByEmail(loginUserForm.Email)
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	password := fmt.Sprintf("%x", sha256.Sum256([]byte(loginUserForm.Password)))
	if password != user.Password {
		panic.PanicException(constants.InvalidRequest)
	}

	accessToken, err := token.GenerateToken(uint(user.ID))
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusOK,
		dto.BuildResponse[map[string]string](constants.Success, map[string]string{"token": accessToken}),
	)
}
