package cli

import (
	"wallshrink/infrastructure/localfile/repository/image_set_local_file_repository"

	"github.com/samber/do"
)

func inject() {
	// Default injector
	do.Provide(nil, func(*do.Injector) (int, error) {
		return 42, nil
	})

	do.Provide(nil, image_set_local_file_repository.NewImageSetLocalFileRepository)
}
