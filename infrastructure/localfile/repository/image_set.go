package repository

import (
	"wallshrink/domain"

	"github.com/samber/do"
)

type imageSetLocalFileRepository struct{}

func NewImageSetLocalFileRepository(i *do.Injector) (domain.ImageSetRepository, error) {
	return &imageSetLocalFileRepository{}, nil
}

func (r *imageSetLocalFileRepository) LoadImageSet(path string) (*domain.ImageSet, []error) {
	return nil, []error{}
}
