package book

import (
	"gg/domain"
	"gg/utils/paginator"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func (br BookRepository) GetAllBooks() ([]*domain.Book, error) {
	var books []*domain.Book
	if err := br.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br BookRepository) GetBookList(p paginator.Paginator[[]*domain.Book]) ([]*domain.Book, error) {
	var books []*domain.Book
	if err := br.db.Scopes(p.Paginate(books, br.db)).
		Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("ID", "Email", "Avatar")
		}).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br BookRepository) CreateBook(book *domain.Book) error {
	if err := br.db.Create(book).Error; err != nil {
		return err
	}

	return nil
}

func (br BookRepository) DeleteBook(book *domain.Book) error {
	if err := br.db.Delete(&book).Error; err != nil {
		return err
	}

	return nil
}

func (br BookRepository) DeleteBookByID(ID string) error {
	if err := br.db.Delete(&domain.Book{}, ID).Error; err != nil {
		return err
	}

	return nil
}

func (br BookRepository) GetUserBookByID(ID string) (*domain.Book, error) {
	var book domain.Book

	if err := br.db.Model(domain.Book{}).First(&book, ID).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
