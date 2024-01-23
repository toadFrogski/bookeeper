package user

import (
	"gg/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (ur *UserRepository) CreateUser(u *domain.User) error {
	if err := ur.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	user := domain.User{Email: email}
	if err := ur.db.Preload("Roles").First(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}
