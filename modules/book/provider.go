package book

import (
	"gg/domain"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	ctl     *BookControllerAPI
	ctlOnce sync.Once

	svc     *BookServiceImpl
	svcOnce sync.Once

	repo     *BookRepositoryImpl
	repoOnce sync.Once

	BookProviderSet wire.ProviderSet = wire.NewSet(
		ProvideBookController,
		ProvideBookRepository,
		ProvideBookService,

		wire.Bind(new(domain.BookController), new(*BookControllerAPI)),
		wire.Bind(new(domain.BookService), new(*BookServiceImpl)),
		wire.Bind(new(domain.BookRepository), new(*BookRepositoryImpl)),
	)
)

func ProvideBookController(svc domain.BookService) *BookControllerAPI {
	ctlOnce.Do(func() {
		ctl = &BookControllerAPI{svc: svc}
	})
	return ctl
}

func ProvideBookService(bookRepository domain.BookRepository) *BookServiceImpl {
	svcOnce.Do(func() {
		svc = &BookServiceImpl{bookRepository: bookRepository}
	})
	return svc
}

func ProvideBookRepository(db *gorm.DB) *BookRepositoryImpl {
	repoOnce.Do(func() {
		repo = &BookRepositoryImpl{db: db}
	})
	return repo
}
