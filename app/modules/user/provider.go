package user

import (
	"gg/domain"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	ctl     *UserControllerImpl
	ctlOnce sync.Once

	svc     *UserServiceImpl
	svcOnce sync.Once

	repo     *UserRepositoryImpl
	repoOnce sync.Once

	UserProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserController,
		ProvideUserRepository,
		ProvideUserService,

		wire.Bind(new(domain.UserController), new(*UserControllerImpl)),
		wire.Bind(new(domain.UserService), new(*UserServiceImpl)),
		wire.Bind(new(domain.UserRepository), new(*UserRepositoryImpl)),
	)
)

func ProvideUserController(svc domain.UserService) *UserControllerImpl {
	ctlOnce.Do(func() {
		ctl = &UserControllerImpl{svc: svc}
	})
	return ctl
}

func ProvideUserService(repo domain.UserRepository) *UserServiceImpl {
	svcOnce.Do(func() {
		svc = &UserServiceImpl{repo: repo}
	})
	return svc
}

func ProvideUserRepository(db *gorm.DB) *UserRepositoryImpl {
	repoOnce.Do(func() {
		repo = &UserRepositoryImpl{db: db}
	})
	return repo
}
