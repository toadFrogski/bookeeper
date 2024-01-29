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
	var user domain.User
	var RegisterUserForm RegisterUserForm
	var userExist bool

	if err := c.ShouldBind(&RegisterUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	if err := userRepo.db.Model(domain.User{}).
		Select("count(*) > 0").Where("email = ?", RegisterUserForm.Email).
		Find(&userExist).Error; err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	if userExist {
		panic.PanicException(constants.RegistredEmail)
	}

	user.Email = RegisterUserForm.Email
	user.Username = RegisterUserForm.Username
	user.Password = RegisterUserForm.Password

	if err := us.UserRepo.CreateUser(&user); err != nil {
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
