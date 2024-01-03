package domain

import "github.com/google/uuid"

type ImageSet struct {
	Path                   string
	BaseNameToImageFileMap map[string]ImageFile // [Basename of the file] -> [ImageFile]
}

func (s *ImageSet) CreateTempCompressedImageFile(
	srcImageFile ImageFile,
	scaleDownDimension Dimension,
) (
	tempImageSet ImageSet,
	createdTempImageFile ImageFile,
	err error,
) {
	createdTempImageFile = ImageFile{
		Dimension:      srcImageFile.Dimension.ScaleDown(scaleDownDimension),
		Stem:           uuid.New().String(),
		Extension:      ".webp",
		ParentImageSet: *s,
	}
	return *s, createdTempImageFile, nil
}
