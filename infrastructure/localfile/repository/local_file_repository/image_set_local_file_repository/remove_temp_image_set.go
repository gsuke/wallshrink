package image_set_local_file_repository

import (
	"fmt"
	"os"
	"wallshrink/domain"

	"github.com/samber/do"
)

func (r *imageSetLocalFileRepository) RemoveTempImageSet(tempImageSet domain.ImageSet) []error {
	if tempImageSet.Path != os.TempDir() {
		return []error{fmt.Errorf("%w: \"%s\"", domain.ErrIsNotTemporaryImageSet, tempImageSet.Path)}
	}

	imageFileRepository := do.MustInvoke[domain.ImageFileRepository](nil)

	errs := []error{}
	for _, f := range tempImageSet.BaseNameToImageFileMap {
		err := imageFileRepository.RemoveImageFile(f)
		if err != nil {
			errs = append(errs, err)
			fmt.Printf("[!] %s\n", err)
		}
	}

	if len(errs) != 0 {
		return errs
	}
	return nil
}
