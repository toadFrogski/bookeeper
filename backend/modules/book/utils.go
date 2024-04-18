package book

import (
	"fmt"
	"regexp"
)

func validateBookForm(form *BookForm) error {
	fileType := form.Photo.Header.Get("Content-Type")
	if ok, _ := regexp.MatchString("image/*", fileType); !ok {
		return fmt.Errorf("Uploaded file is not image")
	}

	return nil
}
