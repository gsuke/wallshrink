package image_set_local_file_repository

import (
	"fmt"
	"os"
	"path/filepath"
	"wallshrink/domain"
)

func (r *imageSetLocalFileRepository) LoadImageSet(path string) (imageSet domain.ImageSet, warnings []error, err error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return domain.ImageSet{}, nil, fmt.Errorf("%w: %s", domain.ErrImageSetLoadFailed, err)
	}

	imageSet = domain.ImageSet{
		Path:                   path,
		BaseNameToImageFileMap: map[string]domain.ImageFile{},
	}

	// Load all ImageFiles in ImageSet
	for i, f := range files {
		NewImageSet, err := r.LoadImageFile(imageSet, f.Name())

		fmt.Printf("%d/%d: %s\n", i+1, len(files), filepath.Join(path, f.Name()))

		if err != nil {
			warnings = append(warnings, err)
			fmt.Println("  ↑ [!] Failed to load image information. The directory should contain only image files.")
			continue
		}

		imageSet = NewImageSet

	}

	return imageSet, warnings, nil
}
