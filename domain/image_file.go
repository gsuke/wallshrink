package domain

import (
	"path/filepath"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type ImageFileParentless struct {
	Size      int
	Dimension Dimension
	BaseName  BaseName
}

type ImageFile struct {
	ImageFileParentless
	ImageSetPath string
}

func (f *ImageFile) FullPath() string {
	return filepath.Join(f.ImageSetPath, f.BaseName.String())
}

func (f *ImageFile) CompressTemp(tempImageSet ImageSet, scaleDownDimension Dimension, quality int) (ImageFile, error) {
	imageFileRepository := do.MustInvoke[ImageFileRepository](nil)

	compressedImageFile := ImageFile{}
	compressedImageFile.Dimension = f.Dimension.ScaleDown(scaleDownDimension)
	compressedImageFile.BaseName.Stem = uuid.NewString()
	compressedImageFile.BaseName.Extension = ".webp"
	compressedImageFile.ImageSetPath = tempImageSet.Path

	compressedImageFile, err := imageFileRepository.Compress(*f, compressedImageFile, quality)
	if err != nil {
		return ImageFile{}, err
	}

	return compressedImageFile, nil
}
