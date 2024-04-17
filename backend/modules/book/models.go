package book

import (
	"bookeeper/domain"
	"bookeeper/utils/dto"
	"bookeeper/utils/paginator"
	"mime/multipart"
)

type (
	SaveBookForm struct {
		Name        string                `form:"name" binding:"required"`
		Author      string                `form:"author" binding:"required"`
		Description string                `form:"description" binding:"required"`
		Photo       *multipart.FileHeader `form:"photo" binding:"required"`
	} // @name SaveBookForm

	BookPaginator    = paginator.Paginator[[]domain.Book] // @name BookPaginator
	BookListResponse = dto.Response[BookPaginator]        // @name BookListResponse
	BookResponse     = dto.Response[domain.Book]          // @name BookResponse
	BooksResponse    = dto.Response[[]domain.Book]        // @name BooksResponse
)
