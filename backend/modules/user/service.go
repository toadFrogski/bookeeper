package user

import (
	"bookeeper/domain"
	"bookeeper/utils/constants"
	"bookeeper/utils/dto"
	"bookeeper/utils/panic"
	"bookeeper/utils/token"
	"net/http"
	"net/mail"

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

	if ok := us.validateRegister(c, RegisterUserForm); !ok {
		return
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
		dto.BuildResponse[Auth](constants.Success, Auth(*accessToken)),
	)
}

func (us UserService) Login(c *gin.Context) {
	var user *domain.User
	var loginUserForm *LoginUserForm

	if err := c.ShouldBind(&loginUserForm); err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	user = us.validateLogin(c, loginUserForm)
	if user == nil {
		return
	}

	accessToken, err := token.GenerateToken(user)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(
		http.StatusOK,
		dto.BuildResponse[Auth](constants.Success, Auth(*accessToken)),
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

func (us UserService) validateRegister(c *gin.Context, registerForm *RegisterUserForm) bool {
	var emailErrors, usernameErrors *NamedValidationErrors

	emailErrors = &NamedValidationErrors{Name: "email"}
	usernameErrors = &NamedValidationErrors{Name: "username"}

	if _, err := mail.ParseAddress(registerForm.Email); err != nil {
		emailErrors.Errors = append(
			emailErrors.Errors,
			ValidationError{Type: "invalid_email", Description: "Email has invalid format"},
		)
	}

	if len(emailErrors.Errors) != 0 {
		c.JSON(http.StatusBadRequest, dto.BuildResponse[NamedValidationErrors](constants.InvalidRequest, *emailErrors))
		c.Abort()
		return false
	}

	if exist := us.UserRepo.IsUserAttributeExist("email", registerForm.Email); exist {
		emailErrors.Errors = append(
			emailErrors.Errors,
			ValidationError{Type: "existed_email", Description: "This email already registered"},
		)
		c.JSON(http.StatusBadRequest, dto.BuildResponse[NamedValidationErrors](constants.InvalidRequest, *emailErrors))
		c.Abort()
		return false
	}

	if exist := us.UserRepo.IsUserAttributeExist("username", registerForm.Email); exist {
		usernameErrors.Errors = append(
			usernameErrors.Errors,
			ValidationError{Type: "existed_username", Description: "This username already registered"},
		)
		c.JSON(http.StatusBadRequest, dto.BuildResponse[NamedValidationErrors](constants.InvalidRequest, *usernameErrors))
		c.Abort()
		return false
	}

	return true
}

func (us UserService) validateLogin(c *gin.Context, loginForm *LoginUserForm) *domain.User {
	var user *domain.User

	user, err := us.UserRepo.GetUserByEmail(loginForm.Email)
	if err != nil {
		user, err = us.UserRepo.GetUserByUsername(loginForm.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.BuildResponse[any](constants.UserNotFound, nil))
			c.Abort()
			return nil
		}
	}

	if err := user.ValidatePassword(loginForm.Password); err != nil {
		c.JSON(http.StatusBadRequest, dto.BuildResponse[any](constants.IncorrectPassword, nil))
		c.Abort()
		return nil
	}

	return user
}
