package user

import (
	"bookeeper/domain"
	"bookeeper/utils/constants"
	"bookeeper/utils/dto"
	_ "bookeeper/utils/dto"
	"bookeeper/utils/panic"
	"bookeeper/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserSvc domain.IUserService
}

// Register godoc
// @Summary Register new  user
// @Tags auth
// @Accept  json
// @Produce json
// @Param   register body RegisterUserForm true "Register form"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} NamedValidationErrorsResponse
// @Failure 500 {object} AnyResponse
// @Router /register [post]
func (uc UserController) Register(c *gin.Context) {
	uc.UserSvc.Register(c)
}

// Login godoc
// @Summary Login user
// @Tags auth
// @Accept  json
// @Produce json
// @Param   login body LoginUserForm true "Login form"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} NamedValidationErrorsResponse
// @Failure 500 {object} AnyResponse
// @Router /login [post]
func (uc UserController) Login(c *gin.Context) {
	uc.UserSvc.Login(c)
}

// GetUserInfo godoc
// @Summary Get user info
// @Tags user
// @Produce json
// @Success 200 {object} UserResponse
// @Failure 400 {object} AnyResponse
// @Failure 401 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /profile [get]
func (uc UserController) GetUserInfo(c *gin.Context) {
	var user *domain.User
	vars, exist := c.Get("user")
	if !exist {
		panic.PanicException(constants.InternalError)
	}
	user = vars.(*domain.User)

	c.JSON(http.StatusOK, dto.BuildResponse[domain.User](constants.Success, *user))
}

// GetUserInfoByID godoc
// @Summary Get user info by ID
// @Tags user
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} AnyResponse
// @Failure 401 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /profile/{user_id} [get]
func (uc UserController) GetUserInfoByID(c *gin.Context) {
	var user *domain.User

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	user = uc.UserSvc.GetUserInfo(uint(userID))
	c.JSON(http.StatusOK, dto.BuildResponse[domain.User](constants.Success, *user))
}

// Refresh godoc
// @Summary Refresh user session
// @Tags auth
// @Produce json
// @Success 200 {object} AuthResponse
// @Failure 400 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /refresh [post]
func (uc UserController) Refresh(c *gin.Context) {
	var user *domain.User
	var newToken *token.Token

	claims, err := token.ExtractTokenClaims(c)
	if err != nil {
		panic.PanicException(constants.InvalidRequest)
	}
	sub, err := claims.GetSubject()
	if err != nil {
		panic.PanicException(constants.InvalidRequest)
	}

	userID, err := strconv.Atoi(sub)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	user = uc.UserSvc.GetUserInfo(uint(userID))
	newToken, err = token.GenerateToken(user)
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	c.JSON(http.StatusOK, dto.BuildResponse[Auth](constants.Success, Auth(*newToken)))
}
