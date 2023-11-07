package book

import (
	"gg/domain"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	bookCtl     *BookControllerImpl
	bookCtlOnce sync.Once

	bookSvc     *BookServiceImpl
	bookSvcOnce sync.Once

	bookRepo     *BookRepositoryImpl
	bookRepoOnce sync.Once

	BookProviderSet wire.ProviderSet = wire.NewSet(
		ProvideBookController,
		ProvideBookRepository,
		ProvideBookService,

		wire.Bind(new(domain.BookController), new(*BookControllerImpl)),
		wire.Bind(new(domain.BookService), new(*BookServiceImpl)),
		wire.Bind(new(domain.BookRepository), new(*BookRepositoryImpl)),
	)
)

func ProvideBookController(bookSvc domain.BookService) *BookControllerImpl {
	bookCtlOnce.Do(func() {
		bookCtl = &BookControllerImpl{bookSvc: bookSvc}
	})
	return bookCtl
}

func ProvideBookService(bookRepo domain.BookRepository) *BookServiceImpl {
	bookSvcOnce.Do(func() {
		bookSvc = &BookServiceImpl{bookRepo: bookRepo}
	})
	return bookSvc
}

func ProvideBookRepository(db *gorm.DB) *BookRepositoryImpl {
	bookRepoOnce.Do(func() {
		bookRepo = &BookRepositoryImpl{db: db}
	})
	return bookRepo
}
