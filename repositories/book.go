package repository

import (
	model "gg/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBooks() ([]model.Book, error)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

func ProvideBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

func (b BookRepositoryImpl) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
