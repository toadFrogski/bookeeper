package user

import (
	"gg/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (ur *UserRepositoryImpl) CreateUser(u *domain.User) error {
	if err := ur.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}
