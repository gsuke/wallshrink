package repository

import (
	"path/filepath"
	"strings"
	"wallshrink/domain"

	"github.com/samber/do"
)

type imageFileLocalFileRepository struct{}

func NewImageFileLocalFileRepository(i *do.Injector) (domain.ImageFileRepository, error) {
	return &imageFileLocalFileRepository{}, nil
}

func (r *imageFileLocalFileRepository) LoadImageFile(path string) (domain.ImageFile, error) {
	stem, extension := splitFileName(path)

	return domain.ImageFile{
		// TODO
		Size:      0,
		Width:     0,
		Height:    0,
		Stem:      stem,
		Extension: extension,
	}, nil
}

func splitFileName(path string) (stem string, extension string) {
	basename := filepath.Base(path)
	extension = filepath.Ext(path)
	stem = strings.TrimSuffix(basename, extension)
	return
}
