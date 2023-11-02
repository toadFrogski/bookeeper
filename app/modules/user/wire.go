//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(db *gorm.DB) *UserControllerImpl {
	panic(wire.Build(UserProviderSet))
}
