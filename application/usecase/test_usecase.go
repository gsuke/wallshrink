package usecase

import (
	"wallshrink/domain"

	"github.com/samber/do"
)

func TestUseCase() {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)

	imageSetRepository.LoadImageSet("foo")
}
