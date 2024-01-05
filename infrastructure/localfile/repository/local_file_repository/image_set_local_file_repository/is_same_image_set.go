package image_set_local_file_repository

import (
	"os"
	"wompressor/domain"
)

func (r *imageSetLocalFileRepository) IsSameImageSets(
	imageSet1 domain.ImageSet,
	imageSet2 domain.ImageSet,
) (bool, error) {
	fileInfo1, err := os.Stat(imageSet1.Path)
	if err != nil {
		return false, err
	}
	fileInfo2, err := os.Stat(imageSet2.Path)
	if err != nil {
		return false, err
	}

	return os.SameFile(fileInfo1, fileInfo2), nil
}
