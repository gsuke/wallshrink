package usecase

import (
	"fmt"
	"wallshrink/domain"

	"github.com/samber/do"
)

func CompressImageSetUseCase(sourcePath string, destinationPath string, scaleDownDimension domain.Dimension) error {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)

	tempImageSet := imageSetRepository.PrepareTempImageSet()

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

	for _, imageFile := range sourceImageSet.BaseNameToImageFileMap {
		compressedImageFile, _ := imageFile.CompressTemp(tempImageSet, scaleDownDimension)
		fmt.Println(compressedImageFile)
	}

	return nil
}
