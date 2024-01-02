package usecase

import (
	"fmt"
	"wallshrink/domain"

	"github.com/samber/do"
)

func TestUseCase(srcPath string) error {
	imageSetRepository := do.MustInvoke[domain.ImageSetRepository](nil)

	imageSet, _, err := imageSetRepository.LoadImageSet(srcPath)
	if err != nil {
		return err
	}

	for _, f := range imageSet.ImageFiles {
		fmt.Println(f.FullPath())
	}
	return nil
}
