package image_file_local_file_repository

import (
	"os"
)

func (r *imageFileLocalFileRepository) RemoveImageFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}
