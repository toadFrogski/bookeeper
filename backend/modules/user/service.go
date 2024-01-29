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
	UserRepo domain.IUserRepository
}

func (us UserService) Register(c *gin.Context) {
	var user *domain.User
	var assertUser *domain.User
	var RegisterUserForm RegisterUserForm

	if err := c.ShouldBind(&RegisterUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	if err := userRepo.db.Model(domain.User{}).Where("email = ?", RegisterUserForm.Email).First(&assertUser).Error; err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	if assertUser != nil {
		panic.PanicException(constants.RegistredEmail)
	}

	user.Email = RegisterUserForm.Email
	user.Username = RegisterUserForm.Username
	user.Password = RegisterUserForm.Password

	if err := us.UserRepo.CreateUser(user); err != nil {
		panic.PanicException(constants.InternalError)
	}

	accessToken, err := token.GenerateToken(*user)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusCreated,
		dto.BuildResponse[AuthResponse](constants.Success, AuthResponse{Token: accessToken}),
	)
}

func (us UserService) Login(c *gin.Context) {
	var user *domain.User
	var loginUserForm LoginUserForm

	if err := c.ShouldBind(&loginUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user, err := us.UserRepo.GetUserByEmail(loginUserForm.Email)
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	if err := user.ValidatePassword(loginUserForm.Password); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	accessToken, err := token.GenerateToken(*user)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusOK,
		dto.BuildResponse[AuthResponse](constants.Success, AuthResponse{Token: accessToken}),
	)
}

func (us UserService) GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, dto.BuildResponse[string](constants.Success, "test"))
}
