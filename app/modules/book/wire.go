//go:build wireinject
// +build wireinject

package book

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(db *gorm.DB) *BookControllerAPI {
	panic(wire.Build(BookProviderSet))
}
