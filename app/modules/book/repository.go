package book

import (
	"gg/domain"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func (br BookRepository) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book
	if err := br.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br BookRepository) SaveBook(book *domain.Book) error {
	return nil
}
