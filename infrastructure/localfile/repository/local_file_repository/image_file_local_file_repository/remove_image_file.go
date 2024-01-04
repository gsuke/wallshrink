package image_file_local_file_repository

import (
	"os"
	"wallshrink/domain"
)

func (r *imageFileLocalFileRepository) RemoveImageFile(imageFile domain.ImageFile) error {
	err := os.Remove(imageFile.FullPath())
	if err != nil {
		return err
	}

	return nil
}
