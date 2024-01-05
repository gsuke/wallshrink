package domain

import (
	"path/filepath"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type ImageFile struct {
	Size               int
	Dimension          Dimension
	BaseName           BaseName
	ParentImageSetPath string
}

func (f *ImageFile) FullPath() string {
	return filepath.Join(f.ParentImageSetPath, f.BaseName.String())
}

func (f *ImageFile) CompressTemp(tempImageSet ImageSet, scaleDownDimension Dimension, quality int) (ImageFile, error) {
	imageFileRepository := do.MustInvoke[ImageFileRepository](nil)

	compressedImageFile := ImageFile{
		Dimension: f.Dimension.ScaleDown(scaleDownDimension),
		BaseName: BaseName{
			Stem:      uuid.NewString(),
			Extension: ".webp",
		},
		ParentImageSetPath: tempImageSet.Path,
	}

	compressedImageFile, err := imageFileRepository.Compress(*f, compressedImageFile, quality)
	if err != nil {
		return ImageFile{}, err
	}

	return compressedImageFile, nil
}
