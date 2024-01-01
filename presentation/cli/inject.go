package cli

import (
	"wallshrink/infrastructure/localfile/repository"

	"github.com/samber/do"
)

func inject() {
	injecter := do.New()

	do.Provide(injecter, repository.NewImageSetLocalFileRepository)
}
