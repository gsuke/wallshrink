package cli

import (
	"wompressor/infrastructure/localfile/repository/local_file_repository/image_file_local_file_repository"
	"wompressor/infrastructure/localfile/repository/local_file_repository/image_set_local_file_repository"

	"github.com/samber/do"
)

func inject() {
	// Default injector
	do.Provide(nil, func(*do.Injector) (int, error) {
		return 42, nil
	})

	do.Provide(nil, image_set_local_file_repository.NewImageSetLocalFileRepository)
	do.Provide(nil, image_file_local_file_repository.NewImageFileLocalFileRepository)
}
