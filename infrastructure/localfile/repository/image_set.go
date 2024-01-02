package repository

import (
	"fmt"
	"os"
	"wallshrink/domain"

	"github.com/samber/do"
)

type imageSetLocalFileRepository struct{}

func NewImageSetLocalFileRepository(i *do.Injector) (domain.ImageSetRepository, error) {
	return &imageSetLocalFileRepository{}, nil
}

func (r *imageSetLocalFileRepository) LoadImageSet(path string) (imageSet domain.ImageSet, warnings []error, err error) {
	imageFileRepository := do.MustInvoke[domain.ImageFileRepository](nil)

	files, err := os.ReadDir(path)
	if err != nil {
		return domain.ImageSet{}, nil, fmt.Errorf("%w: %s", domain.ErrDirectoryLoadFailed, err)
	}

	var imageFiles []domain.ImageFile
	for _, f := range files {
		// TODO: handle error
		fullPath := path + "/" + f.Name()
		imageFile, _ := imageFileRepository.LoadImageFile(fullPath)
		imageFiles = append(imageFiles, imageFile)
	}

	return domain.ImageSet{
		Path:       path,
		ImageFiles: imageFiles,
	}, []error{}, nil
}
