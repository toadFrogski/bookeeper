package domain

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Username string `gorm:"column:username" json:"username"`
		Password string `gorm:"password" json:"-"`
		Email    string `gorm:"column:email" json:"email"`
		Books    []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"books"`
	}

	UserRepository interface {
		CreateUser(u *User) error
		GetUserByEmail(email string) (User, error)
	}

	UserService interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
	}

	UserController interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))
	u.CreatedAt = time.Now()

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()

	return nil
}
