package image_file_local_file_repository

import (
	"fmt"
	"os"
	"wallshrink/domain"
	repository "wallshrink/infrastructure/localfile/repository/local_file_repository"

	"github.com/samber/do"
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

	return &imageFileLocalFileRepository{}, nil
}
