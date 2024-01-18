package book

import "mime/multipart"

type (
	SaveBookForm struct {
		Name        string                `form:"name" binding:"required"`
		Author      string                `form:"author" binding:"required"`
		Description string                `form:"description" binding:"required"`
		UserID      uint                  `form:"user_id" binding:"required"`
		Photo       *multipart.FileHeader `form:"photo" binding:"required"`
	}
)
