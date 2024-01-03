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
		return domain.ImageSet{}, nil, fmt.Errorf("%w: %s", domain.ErrImageSetLoadFailed, err)
	}

	imageSet = domain.ImageSet{
		Path:                   path,
		BaseNameToImageFileMap: map[string]domain.ImageFile{},
	}

	for i, f := range files {
		imageFile, err := imageFileRepository.LoadImageFile(imageSet, filepath.Base(f.Name()))

		if err != nil {
			warnings = append(warnings, err)
			fmt.Println("[!] Failed to image information. The directory should contain only image files.")
			fmt.Println(err)
			continue
		}

		imageSet.BaseNameToImageFileMap[imageFile.BaseName()] = imageFile
		fmt.Printf("%d/%d: %s\n", i, len(files), imageFile.FullPath())
	}

	return imageSet, warnings, nil
}
