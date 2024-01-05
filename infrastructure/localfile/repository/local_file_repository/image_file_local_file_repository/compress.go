package image_file_local_file_repository

import (
	"fmt"
	"wompressor/domain"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// Compress compresses image file.
// If quality = 0, it will be lossless.
func (r *imageFileLocalFileRepository) Compress(
	srcImageFile domain.ImageFile,
	destImageFile domain.ImageFile,
	quality int,
) (domain.ImageFile, error) {

	// ffmpeg.KwArgs
	kwArgs := ffmpeg.KwArgs{
		"s": fmt.Sprintf("%dx%d", destImageFile.Dimension.Width, destImageFile.Dimension.Height),
	}
	if quality == 0 {
		kwArgs["lossless"] = 1
	} else {
		kwArgs["quality"] = quality
	}

	// Compress
	err := ffmpeg.
		Input(srcImageFile.FullPath()).
		Output(
			destImageFile.FullPath(),
			kwArgs,
		).
		Run()
	if err != nil {
		return domain.ImageFile{}, err
	}

	// Get compressed file size
	size, err := getFileSize(destImageFile.FullPath())
	if err != nil {
		return domain.ImageFile{}, err
	}

	destImageFile.Size = size
	return destImageFile, err
}
