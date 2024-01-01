package usecase

import (
	"wallshrink/domain"

	"github.com/samber/do"
)

func TestUseCase(srcPath string) {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)

	imageSetRepository.LoadImageSet(srcPath)
}
