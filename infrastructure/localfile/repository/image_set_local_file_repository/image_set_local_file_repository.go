package image_set_local_file_repository

import (
	"fmt"
	"os"
	"wallshrink/domain"
	"wallshrink/infrastructure/localfile/repository"

	"github.com/samber/do"
)

type imageSetLocalFileRepository struct{}

func NewImageSetLocalFileRepository(i *do.Injector) (domain.ImageSetRepository, error) {
	if !isFFProbeAvailable() {
		fmt.Println(repository.ErrFFProbeIsNotAvailable)
		os.Exit(1)
	}

	return &imageSetLocalFileRepository{}, nil
}
