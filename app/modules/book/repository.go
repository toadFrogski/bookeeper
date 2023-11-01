package book

import (
	"gg/domain"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func (b BookRepositoryImpl) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
