package repository

import (
	"fmt"
	"os"
	"path/filepath"
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

	imageSet = domain.ImageSet{
		Path:       path,
		ImageFiles: []domain.ImageFile{},
	}

	for _, f := range files {
		imageFile, _ := imageFileRepository.LoadImageFile(imageSet, filepath.Base(f.Name())) // TODO: handle error
		imageSet.ImageFiles = append(imageSet.ImageFiles, imageFile)
	}

	return imageSet, []error{}, nil
}
