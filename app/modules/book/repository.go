package book

import (
	"gg/domain"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func (br BookRepositoryImpl) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book
	if err := br.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
