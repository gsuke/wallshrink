package image_set_local_file_repository

import "wallshrink/domain"

func (r *imageSetLocalFileRepository) LoadImageFile(imageSet domain.ImageSet, fileBaseName string) (domain.ImageSet, error) {
	stem, extension := splitFileName(fileBaseName)

	imageFile := domain.ImageFile{
		Stem:           stem,
		Extension:      extension,
		ParentImageSet: imageSet,
	}

	// Get file size
	size, err := getFileSize(imageFile.FullPath())
	if err != nil {
		return domain.ImageSet{}, err
	}
	imageFile.Size = size

	// Get image dimension
	dimension, err := getImageDimension(imageFile.FullPath())
	if err != nil {
		return domain.ImageSet{}, err
	}
	imageFile.Dimension = dimension

	imageSet.BaseNameToImageFileMap[imageFile.BaseName()] = imageFile

	return imageSet, nil
}
