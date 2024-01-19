package domain

import (
	"html"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		ID       uint    `gorm:"primarykey"`
		Username string  `gorm:"column:username" json:"username"`
		Password string  `gorm:"password" json:"-"`
		Email    string  `gorm:"column:email" json:"email"`
		Books    []*Book `json:"-"`
		Roles    []*Role `gorm:"many2many:user_roles" json:"-"`
	}

	IUserController interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
	}

	IUserService interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
	}

	IUserRepository interface {
		CreateUser(u *User) error
		GetUserByEmail(email string) (User, error)
	}
)

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
