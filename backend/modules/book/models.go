package book

import (
	"gg/domain"
	"gg/utils/dto"
	"gg/utils/paginator"
	"mime/multipart"
)

type (
	SaveBookForm struct {
		Name        string                `form:"name" binding:"required"`
		Author      string                `form:"author" binding:"required"`
		Description string                `form:"description" binding:"required"`
		UserID      uint                  `form:"user_id" binding:"required"`
		Photo       *multipart.FileHeader `form:"photo" binding:"required"`
	} // @name SaveBookForm

	BookPaginator    = paginator.Paginator[[]domain.Book] // @name BookPaginator
	BookListResponse = dto.Response[BookPaginator]        // @name BookListResponse
	BookResponse     = dto.Response[domain.Book]          // @name BookResponse
)
