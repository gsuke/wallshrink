package domain

import (
	"path/filepath"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type ImageFileParentless struct {
	Size      int
	Dimension Dimension
	Stem      string
	Extension string // includes "."
}

type ImageFile struct {
	ImageFileParentless
	ParentImageSet ImageSet
}

func (f *ImageFileParentless) BaseName() string {
	return f.Stem + f.Extension
}

func (f *ImageFile) FullPath() string {
	return filepath.Join(f.ParentImageSet.Path, f.Stem+f.Extension)
}

func (f *ImageFile) CompressTemp(tempImageSet ImageSet, scaleDownDimension Dimension, quality int) (ImageFile, error) {
	imageFileRepository := do.MustInvoke[ImageFileRepository](nil)

	compressedImageFile := ImageFile{}
	compressedImageFile.Dimension = f.Dimension.ScaleDown(scaleDownDimension)
	compressedImageFile.Stem = uuid.NewString()
	compressedImageFile.Extension = ".webp"
	compressedImageFile.ParentImageSet = tempImageSet

	compressedImageFile, err := imageFileRepository.Compress(*f, compressedImageFile, quality)
	if err != nil {
		return ImageFile{}, err
	}

	return compressedImageFile, nil
}
