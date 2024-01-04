package usecase

import (
	"errors"
	"fmt"
	"path/filepath"
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

	// Compress all image files
	i := 0
	for _, imageFile := range sourceImageSet.BaseNameToImageFileMap {
		fmt.Printf("[%d/%d]\n", i+1, len(sourceImageSet.BaseNameToImageFileMap))

		// Attempted compression: quality 50 -> 100
		for quality := 50; quality <= 100; quality += 10 {
			if quality > 100 {
				// TODO: define error
				return errors.New("failed to set quality")
			}

			fmt.Printf(
				" [quality=%d] Attempting compression: %s\n",
				quality,
				imageFile.BaseName.String(),
			)

			// Compress temporarily
			compressedImageFile, err := imageFile.CompressTemp(tempImageSet, scaleDownDimension, quality)
			if err != nil {
				return err
			}
			tempImageSet.BaseNameToImageFileMap[compressedImageFile.BaseName] = compressedImageFile

			// Calculate SSIM
			ssim, _ := imageFileRepository.SSIM(imageFile, compressedImageFile)
			fmt.Printf(" SSIM: %f\n", ssim)

			if ssim > 0.98 {
				fmt.Println(" SSIM OK!")

				// Compare 2 images
				isFilesSame, err := imageFileRepository.IsFilesSame(
					compressedImageFile.FullPath(),
					filepath.Join(destinationImageSet.Path, imageFile.BaseName.String()),
				)
				if err != nil {
					return err
				}

				// Copy compressed image file to destination directory
				if isFilesSame {
					fmt.Println(" Skip: no file replacement required")
				} else {
					destinationImageSet, _, err = imageSetRepository.CopyImageFile(
						compressedImageFile,
						destinationImageSet,
						imageFile.BaseName,
					)
					if err != nil {
						return err
					}
				}

				break
			}
		}

		i++
	}

	// TODO: delete later
	fmt.Println(len(sourceImageSet.BaseNameToImageFileMap), len(destinationImageSet.BaseNameToImageFileMap))

	return nil
}
