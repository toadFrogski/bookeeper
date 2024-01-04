package book

import "mime/multipart"

type (
	SaveBookForm struct {
		Name        string                `form:"name" binding:"required"`
		Author      string                `form:"author" binding:"required"`
		Description string                `form:"description" binding:"required"`
		Photo       *multipart.FileHeader `form:"image" binding:"required"`
	}
)
