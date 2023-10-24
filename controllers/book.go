package controller

import (
	service "gg/services"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	GetAllBooks(c *gin.Context)
}

type BookControllerAPI struct {
	svc service.BookService
}

func ProvideBookController(svc service.BookService) *BookControllerAPI {
	return &BookControllerAPI{svc: svc}
}

func (b BookControllerAPI) GetAllBooks(c *gin.Context) {
	b.svc.GetAllBooks(c)
}
