package usecase

import (
	"fmt"
	"math"
	"strings"
	"wompressor/domain"

	humanize "github.com/dustin/go-humanize"
	"github.com/samber/do"
)

func CompressImageSetUseCase(sourcePath string, destinationPath string, scaleDownDimension domain.Dimension) error {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)
	imageFileRepository := do.MustInvoke[domain.ImageFileRepository](nil)

	tempImageSet := imageSetRepository.PrepareTempImageSet()
	defer imageSetRepository.RemoveTempImageSet(tempImageSet) // TODO: delete temp files one at a time

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

	// detect duplicate stem
	if err = detectDuplicateStem(sourceImageSet); err != nil {
		return err
	}
	if err = detectDuplicateStem(destinationImageSet); err != nil {
		return err
	}

	// Remove files from destination image set that are not in the source image set
	targetImageFiles := []domain.ImageFile{}
	for baseName, destImageFile := range destinationImageSet.BaseNameToImageFileMap {
		srcImageFiles := sourceImageSet.GetImageFilesByStem(baseName.Stem)
		if len(srcImageFiles) == 0 {
			targetImageFiles = append(targetImageFiles, destImageFile)
		}
	}
	for _, imageFile := range targetImageFiles {
		delete(destinationImageSet.BaseNameToImageFileMap, imageFile.BaseName)
		fmt.Printf("Deleted %s\n", imageFile.FullPath())
		imageFileRepository.RemoveImageFile(imageFile)
	}

	// Compress all image files
	i := 0
	for _, imageFile := range sourceImageSet.BaseNameToImageFileMap {
		fmt.Printf("[%d/%d]\n", i+1, len(sourceImageSet.BaseNameToImageFileMap))

		// Attempted compression
		// 0: lossless
		// -1: error
		for _, quality := range []int{85, 100, 0, -1} {
			if quality == -1 {
				return ErrSSIMShortage
			}

			// Display message
			if quality == 0 {
				fmt.Printf(
					" [lossless]: %s\n",
					imageFile.BaseName.String(),
				)
			} else {
				fmt.Printf(
					" [quality=%d] Attempting compression: %s\n",
					quality,
					imageFile.BaseName.String(),
				)
			}

			// Compress temporarily
			compressedImageFile, err := imageFile.CompressTemp(tempImageSet, scaleDownDimension, quality)
			if err != nil {
				return err
			}
			tempImageSet.BaseNameToImageFileMap[compressedImageFile.BaseName] = compressedImageFile

			// Calculate SSIM
			ssim, _ := imageFileRepository.SSIM(imageFile, compressedImageFile)
			fmt.Printf(" SSIM: %f\n", ssim)
			if ssim < 0.99 {
				continue
			}

			// Calculate compression ratio: (compressed file size / src file size)
			compressionRatio := int(math.Round(
				float64(compressedImageFile.Size) / float64(imageFile.Size) * 100.0,
			))
			fmt.Printf(
				" SSIM OK! %s[100%%] -> %s[%d%%]\n",
				humanize.IBytes(uint64(imageFile.Size)),
				humanize.IBytes(uint64(compressedImageFile.Size)),
				compressionRatio,
			)

			// replace if ratio < 100
			// no replace if ratio >= 100
			var resultImageFile domain.ImageFile
			var destinationImageFileBaseName domain.BaseName
			if compressionRatio < 100 {
				resultImageFile = compressedImageFile
				destinationImageFileBaseName = domain.BaseName{
					Stem:      imageFile.BaseName.Stem,
					Extension: ".webp",
				}
			} else {
				resultImageFile = imageFile
				destinationImageFileBaseName = imageFile.BaseName
				fmt.Println(" Not compressed because the ratio exceeded 100%")
			}

			// Compare: src image <-> result image
			// (check only if there is a same stem file in dest)
			if len(destinationImageSet.GetImageFilesByStem(imageFile.BaseName.Stem)) > 0 {
				isFilesSame, err := imageFileRepository.IsFilesSame(
					resultImageFile.FullPath(),
					destinationImageSet.GetImageFilesByStem(imageFile.BaseName.Stem)[0].FullPath(),
				)
				if err != nil {
					return err
				}
				if isFilesSame {
					fmt.Println(" Skip: no file replacement required")
					break
				}
			}

			// Copy compressed image file to destination directory
			destinationImageSet, _, err = imageSetRepository.CopyImageFile(
				resultImageFile,
				destinationImageSet,
				destinationImageFileBaseName,
			)
			if err != nil {
				return err
			}

			break // SSIM OK
		}

		i++
	}

	return nil
}

func detectDuplicateStem(imageSet domain.ImageSet) error {
	stemToImageFilesMap := imageSet.GetDuplicateStemFiles()
	if len(stemToImageFilesMap) == 0 {
		return nil
	}

	// Prepare error message
	var sb strings.Builder
	for stem, imageFiles := range stemToImageFilesMap {
		sb.WriteString("\n")
		sb.WriteString(stem)
		sb.WriteString(": ")
		for _, imageFile := range imageFiles {
			sb.WriteString(imageFile.BaseName.String())
			sb.WriteString(", ")
		}
	}
	return fmt.Errorf("%w%s", ErrThereAreDuplicateStemFiles, sb.String())
}
