package image_file_local_file_repository

import (
	"fmt"
	"os"
	"wompressor/domain"
	repository "wompressor/infrastructure/localfile/repository/local_file_repository"

	"github.com/samber/do"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type imageFileLocalFileRepository struct{}

func NewImageFileLocalFileRepository(i *do.Injector) (domain.ImageFileRepository, error) {
	if !isFFmpegAvailable() {
		fmt.Println(repository.ErrFFmpegIsNotAvailable)
		os.Exit(1)
	}

	if !isFFProbeAvailable() {
		fmt.Println(repository.ErrFFProbeIsNotAvailable)
		os.Exit(1)
	}

	ffmpeg.LogCompiledCommand = false

	return &imageFileLocalFileRepository{}, nil
}
