package book

import (
	"fmt"
	"regexp"
)

func validateSaveBookForm(form *SaveBookForm) error {
	fileType := form.Photo.Header.Get("Content-Type")
	if ok, _ := regexp.MatchString("image/*", fileType); !ok {
		return fmt.Errorf("Uploaded file is not image")
	}

	return nil
}
