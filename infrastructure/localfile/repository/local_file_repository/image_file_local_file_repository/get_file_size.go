package image_file_local_file_repository

import (
	"fmt"
	"os"
	"wallshrink/domain"
)

func getFileSize(path string) (int, error) {
	// Open file
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", domain.ErrFileInfoLoadFailed, err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("%w: %s", domain.ErrFileInfoLoadFailed, err)
	}

	return int(fileInfo.Size()), nil
}
