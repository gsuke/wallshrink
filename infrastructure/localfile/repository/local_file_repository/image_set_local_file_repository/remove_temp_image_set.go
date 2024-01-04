package image_set_local_file_repository

import (
	"fmt"
	"os"
	"wallshrink/domain"

	"github.com/samber/do"
)

func (r *imageSetLocalFileRepository) RemoveTempImageSet(tempImageSet domain.ImageSet) error {
	if tempImageSet.Path != os.TempDir() {
		return fmt.Errorf("%w: \"%s\"", domain.ErrIsNotTemporaryImageSet, tempImageSet.Path)
	}

	imageFileRepository := do.MustInvoke[domain.ImageFileRepository](nil)

	for _, f := range tempImageSet.BaseNameToImageFileMap {
		err := imageFileRepository.RemoveImageFile(f)
		// TODO: Handle error in detail
		if err != nil {
			fmt.Printf("[!] %s\n", err)
		}
	}

	return nil
}
