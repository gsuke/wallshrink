package image_file_local_file_repository

import "wallshrink/domain"

func (r *imageFileLocalFileRepository) LoadImageFile(filePath string) (domain.ImageFileParentless, error) {
	stem, extension := splitFileName(filePath)

	imageFile := domain.ImageFileParentless{
		Stem:      stem,
		Extension: extension,
	}

	// Get file size
	size, err := getFileSize(filePath)
	if err != nil {
		return domain.ImageFileParentless{}, err
	}
	imageFile.Size = size

	// Get image dimension
	dimension, err := getImageDimension(filePath)
	if err != nil {
		return domain.ImageFileParentless{}, err
	}
	imageFile.Dimension = dimension

	return imageFile, nil
}
