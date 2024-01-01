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

func (r *imageSetLocalFileRepository) LoadImageSet(path string) (*domain.ImageSet, []error) {
	files, _ := os.ReadDir(path)

	// var imageFiles []domain.ImageFile
	for _, f := range files {
		fmt.Println(f.Name())
	}

	return nil, []error{}
}
