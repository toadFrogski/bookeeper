package user

import (
	"gg/domain"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	userCtl     *UserControllerImpl
	userCtlOnce sync.Once

	userSvc     *UserServiceImpl
	userSvcOnce sync.Once

	userRepo     *UserRepositoryImpl
	userRepoOnce sync.Once

	UserProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserController,
		ProvideUserRepository,
		ProvideUserService,

		wire.Bind(new(domain.UserController), new(*UserControllerImpl)),
		wire.Bind(new(domain.UserService), new(*UserServiceImpl)),
		wire.Bind(new(domain.UserRepository), new(*UserRepositoryImpl)),
	)
)

func ProvideUserController(userSvc domain.UserService) *UserControllerImpl {
	userCtlOnce.Do(func() {
		userCtl = &UserControllerImpl{userSvc: userSvc}
	})
	return userCtl
}

func ProvideUserService(userRepo domain.UserRepository) *UserServiceImpl {
	userSvcOnce.Do(func() {
		userSvc = &UserServiceImpl{userRepo: userRepo}
	})
	return userSvc
}

func ProvideUserRepository(db *gorm.DB) *UserRepositoryImpl {
	userRepoOnce.Do(func() {
		userRepo = &UserRepositoryImpl{db: db}
	})
	return userRepo
}
