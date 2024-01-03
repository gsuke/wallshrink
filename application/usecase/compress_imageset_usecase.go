package usecase

import (
	"fmt"
	"wallshrink/domain"

	"github.com/samber/do"
)

func CompressImageSetUseCase(sourcePath string, destinationPath string, width int, height int) error {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)

	// Source ImageSet
	sourceImageSet, _, err := imageSetRepository.LoadImageSet(sourcePath)
	if err != nil {
		return err
	}

	// Destination ImageSet
	destinationImageSet, _, err := imageSetRepository.LoadImageSet(destinationPath)
	if err != nil {
		return err
	}

	fmt.Println(sourceImageSet, destinationImageSet)
	return nil
}
