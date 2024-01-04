package usecase

import (
	"errors"
	"fmt"
	"wallshrink/domain"

	"github.com/samber/do"
)

func CompressImageSetUseCase(sourcePath string, destinationPath string, scaleDownDimension domain.Dimension) error {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)
	imageFileRepository := do.MustInvoke[domain.ImageFileRepository](nil)

	tempImageSet := imageSetRepository.PrepareTempImageSet()
	defer imageSetRepository.RemoveTempImageSet(tempImageSet)

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

	// TODO: delete later
	fmt.Println(len(sourceImageSet.BaseNameToImageFileMap), len(destinationImageSet.BaseNameToImageFileMap))

	// Compress all image files
	for _, imageFile := range sourceImageSet.BaseNameToImageFileMap {

		// Attempted compression: quality 50 -> 100
		for quality := 50; quality <= 100; quality += 10 {
			if quality > 100 {
				// TODO: define error
				return errors.New("failed to set quality")
			}

			fmt.Printf("Attempted compression (quality=%d): %s\n", quality, imageFile.BaseName())

			// Compress temporarily
			compressedImageFile, _ := imageFile.CompressTemp(tempImageSet, scaleDownDimension, quality)
			tempImageSet.BaseNameToImageFileMap[compressedImageFile.BaseName()] = compressedImageFile

			// Calculate SSIM
			ssim, _ := imageFileRepository.SSIM(imageFile, compressedImageFile)
			fmt.Printf("SSIM: %f\n", ssim)

			if ssim > 0.98 {
				fmt.Println("SSIM OK!")
				break
			}
		}

	}

	return nil
}
