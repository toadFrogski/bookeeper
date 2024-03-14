package user

import (
	"bookeeper/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (ur UserRepository) CreateUser(u *domain.User) error {
	if err := ur.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user *domain.User
	if err := ur.db.Preload("Roles").Where("email = ? ", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user *domain.User
	if err := ur.db.Preload("Roles").Where("username = ? ", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) GetUserByID(ID uint) (*domain.User, error) {
	var user *domain.User
	if err := ur.db.Preload("Roles").First(&user, ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) GetUserInfoByID(ID uint) (*domain.User, error) {
	var user *domain.User
	if err := ur.db.Model(&domain.User{}).
		Preload("Books", func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("User")
		}).First(&user, ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) IsUserAttributeExist(attr string, value string) bool {
	var userExist bool
	ur.db.Model(domain.User{}).Select("count(*) > 0").Where("? = ?", attr, value).Find(&userExist)

	return userExist
}
