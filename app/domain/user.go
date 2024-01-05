package domain

import (
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	User struct {
		ID       uint   `gorm:"primarykey"`
		Username string `gorm:"column:username" json:"username"`
		Password string `gorm:"password;->:false" json:"-"`
		Email    string `gorm:"column:email" json:"email"`
		Books    []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"books"`
	}

	IUserRepository interface {
		CreateUser(u *User) error
		GetUserByEmail(email string) (User, error)
	}

	IUserService interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
	}

	IUserController interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))

	return nil
}
