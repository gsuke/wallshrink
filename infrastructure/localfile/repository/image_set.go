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
	files, err := os.ReadDir(path)
	if err != nil {
		return domain.ImageSet{}, nil, fmt.Errorf("%w: %s", domain.ErrDirectoryLoadFailed, err)
	}

	// Load Image Files
	var imageFiles []domain.ImageFile
	for _, f := range files {
		imageFiles = append(imageFiles, domain.ImageFile{
			// TODO
			Size:      0,
			Width:     0,
			Height:    0,
			BaseName:  f.Name(),
			Extension: "",
		})
	}

	return domain.ImageSet{
		Path:       path,
		ImageFiles: imageFiles,
	}, []error{}, nil
}
