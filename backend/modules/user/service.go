package user

import (
	"bookeeper/domain"
	"bookeeper/utils/constants"
	"bookeeper/utils/dto"
	"bookeeper/utils/panic"
	"bookeeper/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserRepo domain.IUserRepository
}

func (us UserService) Register(c *gin.Context) {
	var user *domain.User
	var RegisterUserForm *RegisterUserForm

	if err := c.ShouldBind(&RegisterUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	if exist := us.UserRepo.IsUserExist(RegisterUserForm.Email); exist {
		panic.PanicException(constants.RegistredEmail)
	}

	user = &domain.User{
		Email:    RegisterUserForm.Email,
		Username: RegisterUserForm.Username,
		Password: RegisterUserForm.Password,
	}

	if err := us.UserRepo.CreateUser(user); err != nil {
		panic.PanicException(constants.InternalError)
	}

	accessToken, err := token.GenerateToken(user)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusCreated,
		dto.BuildResponse[Auth](constants.Success, Auth{Token: accessToken}),
	)
}

func (us UserService) Login(c *gin.Context) {
	var user *domain.User
	var loginUserForm *LoginUserForm

	if err := c.ShouldBind(&loginUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user, err := us.UserRepo.GetUserByEmail(loginUserForm.Email)
	if err != nil {
		panic.PanicException(constants.InvalidRequest)
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
		dto.BuildResponse[Auth](constants.Success, Auth{Token: accessToken}),
	)
}

func (us UserService) GetUserInfo(c *gin.Context, userID uint) {
	var user *domain.User

	user, err := us.UserRepo.GetUserInfoByID(userID)
	if err != nil {
		panic.PanicException(constants.DataNotFound)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[domain.User](constants.Success, *user))
}
