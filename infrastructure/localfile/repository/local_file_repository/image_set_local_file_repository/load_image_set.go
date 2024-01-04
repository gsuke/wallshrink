package image_set_local_file_repository

import (
	"fmt"
	"os"
	"path/filepath"
	"wallshrink/domain"

	"github.com/samber/do"
)

func (r *imageSetLocalFileRepository) LoadImageSet(dirPath string) (imageSet domain.ImageSet, warnings []error, err error) {
	imageFileRepository := do.MustInvoke[domain.ImageFileRepository](nil)

	dirFiles, err := os.ReadDir(dirPath)
	if err != nil {
		return domain.ImageSet{}, nil, fmt.Errorf("%w: %s", domain.ErrImageSetLoadFailed, err)
	}

	imageSet = domain.ImageSet{
		Path:                   dirPath,
		BaseNameToImageFileMap: map[domain.BaseName]domain.ImageFile{},
	}

	// Load all ImageFiles in ImageSet
	for i, f := range dirFiles {
		baseName := domain.NewBaseName(f.Name())
		imageFilePath := filepath.Join(dirPath, baseName.String())
		fmt.Printf("[%d/%d] %s\n", i+1, len(dirFiles), imageFilePath)

		imageFileParentless, err := imageFileRepository.LoadImageFile(imageFilePath)
		if err != nil {
			warnings = append(warnings, err)
			fmt.Println("  ↑ [!] Failed to load image information. The directory should contain only image files.")
			continue
		}

		imageFile := domain.ImageFile{
			ImageFileParentless: imageFileParentless,
			ImageSetPath:        imageSet.Path,
		}

		imageSet.BaseNameToImageFileMap[imageFile.BaseName] = imageFile
	}

	return imageSet, warnings, nil
}
