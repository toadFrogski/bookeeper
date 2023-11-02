package domain

import (
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	User struct {
		Username string `gorm:"column:username" json:"username"`
		Password string `gorm:"password" json:"-"`
		Email    string `gorm:"column:email" json:"email"`
		BaseModel
	}

	UserRepository interface {
		CreateUser(u *User) error
	}

	UserService interface {
		CreateUser(c *gin.Context)
	}

	UserController interface {
		CreateUser(c *gin.Context)
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))

	return nil
}
