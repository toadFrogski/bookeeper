package user

import (
	"bookeeper/domain"
	"bookeeper/utils/constants"
	_ "bookeeper/utils/dto"
	"bookeeper/utils/panic"
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
// @Failure 400 {object} AnyResponse
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
// @Failure 400 {object} AnyResponse
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
// @Failure 500 {object} AnyResponse
// @Router /profile [get]
func (uc UserController) GetUserInfo(c *gin.Context) {
	var user *domain.User
	vars, exist := c.Get("user")
	if !exist {
		panic.PanicException(constants.InternalError)
	}
	user = vars.(*domain.User)
	uc.UserSvc.GetUserInfo(c, user.ID)
}

// GetUserInfoByID godoc
// @Summary Get user info by ID
// @Tags user
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} AnyResponse
// @Failure 500 {object} AnyResponse
// @Router /profile/{user_id} [get]
func (uc UserController) GetUserInfoByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		panic.PanicException(constants.InternalError)
	}

	uc.UserSvc.GetUserInfo(c, uint(userID))
}
