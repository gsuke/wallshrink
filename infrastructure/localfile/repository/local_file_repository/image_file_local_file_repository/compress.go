package image_file_local_file_repository

import (
	"fmt"
	"wallshrink/domain"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func (r *imageFileLocalFileRepository) Compress(
	srcImageFile domain.ImageFile,
	destImageFile domain.ImageFile,
) (domain.ImageFile, error) {

	// Compress
	err := ffmpeg.
		Input(srcImageFile.FullPath()).
		Output(
			destImageFile.FullPath(),
			ffmpeg.KwArgs{
				"s":       fmt.Sprintf("%dx%d", destImageFile.Dimension.Width, destImageFile.Dimension.Height),
				"quality": 65,
			},
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
