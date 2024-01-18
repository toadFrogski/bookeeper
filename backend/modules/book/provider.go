package book

import (
	"gg/domain"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	bookCtl     *BookController
	bookCtlOnce sync.Once

	bookSvc     *BookService
	bookSvcOnce sync.Once

	bookRepo     *BookRepository
	bookRepoOnce sync.Once

	BookProviderSet wire.ProviderSet = wire.NewSet(
		ProvideBookController,
		ProvideBookRepository,
		ProvideBookService,

		wire.Bind(new(domain.IBookController), new(*BookController)),
		wire.Bind(new(domain.IBookService), new(*BookService)),
		wire.Bind(new(domain.IBookRepository), new(*BookRepository)),
	)
)

func ProvideBookController(bookSvc domain.IBookService, userSvc domain.IUserService) *BookController {
	bookCtlOnce.Do(func() {
		bookCtl = &BookController{bookSvc: bookSvc}
	})
	return bookCtl
}

func ProvideBookService(bookRepo domain.IBookRepository) *BookService {
	bookSvcOnce.Do(func() {
		bookSvc = &BookService{bookRepo: bookRepo}
	})
	return bookSvc
}

func ProvideBookRepository(db *gorm.DB) *BookRepository {
	bookRepoOnce.Do(func() {
		bookRepo = &BookRepository{db: db}
	})
	return bookRepo
}
