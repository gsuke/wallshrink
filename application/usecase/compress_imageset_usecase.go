package usecase

import (
	"fmt"
	"wallshrink/domain"

	"github.com/samber/do"
)

func CompressImageSetUseCase(sourcePath string, destinationPath string, dimensionToScaleDown domain.Dimension) error {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)

	// Source ImageSet
	fmt.Printf("Load Directory: %s\n", sourcePath)
	sourceImageSet, _, err := imageSetRepository.LoadImageSet(sourcePath)
	if err != nil {
		return err
	}

	// Destination ImageSet
	fmt.Printf("Load Directory: %s\n", destinationPath)
	destinationImageSet, _, err := imageSetRepository.LoadImageSet(destinationPath)
	if err != nil {
		return err
	}

	fmt.Println(len(sourceImageSet.BaseNameToImageFileMap), len(destinationImageSet.BaseNameToImageFileMap))
	return nil
}
