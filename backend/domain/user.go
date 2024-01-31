package domain

import (
	"bookeeper/utils/constants"
	"html"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		ID          uint    `gorm:"primarykey"`
		Username    string  `gorm:"column:username" json:"username"`
		Password    string  `gorm:"password;" json:"-"`
		Email       string  `gorm:"column:email;uniqueIndex" json:"email"`
		Description string  `gorm:"column:description; text" json:"description"`
		Avatar      string  `gorm:"column:avatar" json:"avatar"`
		Books       []*Book `json:"books,omitempty"`
		Roles       []*Role `gorm:"many2many:user_roles" json:"-"`
	} // @name User

	IUserController interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
		GetUserInfo(c *gin.Context)
		GetUserInfoByID(c *gin.Context)
	}

	IUserService interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
		GetUserInfo(c *gin.Context, userID uint)
	}

	IUserRepository interface {
		CreateUser(u *User) error
		GetUserByEmail(email string) (*User, error)
		GetUserByID(ID uint) (*User, error)
		GetUserInfoByID(ID uint) (*User, error)
		IsUserExist(email string) bool
	}
)

func (user *User) BeforeSave(tx *gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Description = html.EscapeString(user.Description)

	var userRole Role
	if err := tx.Where("name = ?", constants.User).First(&userRole).Error; err != nil {
		return err
	}
	user.Roles = append(user.Roles, &userRole)
	return nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
