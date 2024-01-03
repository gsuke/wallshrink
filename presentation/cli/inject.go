package cli

import (
	"wallshrink/infrastructure/localfile/repository"

	"github.com/samber/do"
)

func inject() {
	// Default injector
	do.Provide(nil, func(*do.Injector) (int, error) {
		return 42, nil
	})

	do.Provide(nil, repository.NewImageSetLocalFileRepository)
}
