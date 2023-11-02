package book

import (
	"gg/domain"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	ctl     *BookControllerImpl
	ctlOnce sync.Once

	svc     *BookServiceImpl
	svcOnce sync.Once

	repo     *BookRepositoryImpl
	repoOnce sync.Once

	BookProviderSet wire.ProviderSet = wire.NewSet(
		ProvideBookController,
		ProvideBookRepository,
		ProvideBookService,

		wire.Bind(new(domain.BookController), new(*BookControllerImpl)),
		wire.Bind(new(domain.BookService), new(*BookServiceImpl)),
		wire.Bind(new(domain.BookRepository), new(*BookRepositoryImpl)),
	)
)

func ProvideBookController(svc domain.BookService) *BookControllerImpl {
	ctlOnce.Do(func() {
		ctl = &BookControllerImpl{svc: svc}
	})
	return ctl
}

func ProvideBookService(repo domain.BookRepository) *BookServiceImpl {
	svcOnce.Do(func() {
		svc = &BookServiceImpl{repo: repo}
	})
	return svc
}

func ProvideBookRepository(db *gorm.DB) *BookRepositoryImpl {
	repoOnce.Do(func() {
		repo = &BookRepositoryImpl{db: db}
	})
	return repo
}
