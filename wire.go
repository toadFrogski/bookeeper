package main

import (
	controller "gg/controllers"
	repository "gg/repositories"
	service "gg/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func initBookAPI(db *gorm.DB) controller.BookControllerAPI {
	wire.Build(service.ProvideBookService, controller.ProvideBookController, repository.ProvideBookRepository)

	return controller.BookControllerAPI{}
}
