package usecase

import (
	"fmt"
	"math"
	"path/filepath"
	"wompressor/domain"

	humanize "github.com/dustin/go-humanize"
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

	// Src and dest must be different
	isSameImageSets, err := imageSetRepository.IsSameImageSets(sourceImageSet, destinationImageSet)
	if err != nil {
		return err
	}
	if isSameImageSets {
		return ErrSrcAndDestAreSame
	}

	// Remove files from Destination ImageSet that are not in the Source ImageSet
	targetImageFiles := []domain.ImageFile{}
	for baseName, imageFile := range destinationImageSet.BaseNameToImageFileMap {
		if _, ok := sourceImageSet.BaseNameToImageFileMap[baseName]; !ok {
			targetImageFiles = append(targetImageFiles, imageFile)
		}
	}
	for _, imageFile := range targetImageFiles {
		delete(destinationImageSet.BaseNameToImageFileMap, imageFile.BaseName)
		fmt.Printf("Deleted %s.\n", imageFile.FullPath())
		imageFileRepository.RemoveImageFile(imageFile)
	}

	// Compress all image files
	i := 0
	for _, imageFile := range sourceImageSet.BaseNameToImageFileMap {
		fmt.Printf("[%d/%d]\n", i+1, len(sourceImageSet.BaseNameToImageFileMap))

		// Attempted compression: quality 50 -> 100
		for _, quality := range []int{50, 70, 80, 90, 95, 100, -1} {
			if quality == -1 {
				return ErrSSIMShortage
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
				fmt.Printf(
					" SSIM OK! %s[100%%] -> %s[%d%%]\n",
					humanize.IBytes(uint64(imageFile.Size)),
					humanize.IBytes(uint64(compressedImageFile.Size)),
					int(math.Round(float64(compressedImageFile.Size)/float64(imageFile.Size)*100.0)),
				)

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

	return nil
}
