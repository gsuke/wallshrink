package domain

import (
	"path/filepath"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type ImageFile struct {
	Size           int
	Dimension      Dimension
	Stem           string
	Extension      string // includes "."
	ParentImageSet ImageSet
}

func (f *ImageFile) BaseName() string {
	return f.Stem + f.Extension
}

func (f *ImageFile) FullPath() string {
	return filepath.Join(f.ParentImageSet.Path, f.Stem+f.Extension)
}

func (f *ImageFile) CompressTemp(tempImageSet ImageSet, scaleDownDimension Dimension) (ImageFile, error) {
	imageFileRepository := do.MustInvoke[ImageFileRepository](nil)

	compressedImageFile := ImageFile{
		Dimension:      f.Dimension.ScaleDown(scaleDownDimension),
		Stem:           uuid.NewString(),
		Extension:      ".webp",
		ParentImageSet: tempImageSet,
	}

	compressedImageFile, err := imageFileRepository.Compress(*f, compressedImageFile)
	if err != nil {
		return ImageFile{}, err
	}

	return compressedImageFile, nil
}
