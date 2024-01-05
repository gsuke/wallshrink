package image_file_local_file_repository

import (
	"path/filepath"
	"wompressor/domain"
)

func (r *imageFileLocalFileRepository) LoadImageFile(filePath string) (domain.ImageFile, error) {
	imageFile := domain.ImageFile{
		BaseName:           domain.NewBaseName(filePath),
		ParentImageSetPath: filepath.Dir(filePath),
	}

	// Get file size
	size, err := getFileSize(filePath)
	if err != nil {
		return domain.ImageFile{}, err
	}
	imageFile.Size = size

	// Get image dimension
	dimension, err := getImageDimension(filePath)
	if err != nil {
		return domain.ImageFile{}, err
	}
	imageFile.Dimension = dimension

	return imageFile, nil
}
