package domain

import (
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	ADMIN UserRole = "admin"
	USER  UserRole = "user"
)

type (
	UserRole string

	User struct {
		ID       uint     `gorm:"primarykey"`
		Username string   `gorm:"column:username" json:"username"`
		Password string   `gorm:"password" json:"-"`
		Email    string   `gorm:"column:email" json:"email"`
		Role     UserRole `gorm:"column:role" json:"-"`
		Books    []Book   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"books"`
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

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))

	return nil
}
