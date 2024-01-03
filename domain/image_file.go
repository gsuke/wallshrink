package domain

import "github.com/google/uuid"

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

func (f *ImageFile) CompressTemp(tempImageSet ImageSet, scaleDownDimension Dimension) (ImageFile, error) {
	compressedImageFile := ImageFile{
		Dimension:      f.Dimension.ScaleDown(scaleDownDimension),
		Stem:           uuid.NewString(),
		Extension:      ".webp",
		ParentImageSet: tempImageSet,
	}
	return compressedImageFile, nil
}
