package user

import (
	"gg/domain"
	"sync"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var (
	userCtl     *UserController
	userCtlOnce sync.Once

	userSvc     *UserService
	userSvcOnce sync.Once

	userRepo     *UserRepository
	userRepoOnce sync.Once

	UserProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserController,
		ProvideUserRepository,
		ProvideUserService,

		wire.Bind(new(domain.IUserController), new(*UserController)),
		wire.Bind(new(domain.IUserService), new(*UserService)),
		wire.Bind(new(domain.IUserRepository), new(*UserRepository)),
	)
)

func ProvideUserController(userSvc domain.IUserService) *UserController {
	userCtlOnce.Do(func() {
		userCtl = &UserController{userSvc: userSvc}
	})
	return userCtl
}

func ProvideUserService(userRepo domain.IUserRepository) *UserService {
	userSvcOnce.Do(func() {
		userSvc = &UserService{userRepo: userRepo}
	})
	return userSvc
}

func ProvideUserRepository(db *gorm.DB) *UserRepository {
	userRepoOnce.Do(func() {
		userRepo = &UserRepository{db: db}
	})
	return userRepo
}
