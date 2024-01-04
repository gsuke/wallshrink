package domain

type ImageSetRepository interface {
	LoadImageSet(path string) (imageSet ImageSet, warnings []error, err error)
	PrepareTempImageSet() ImageSet
	CopyImageFile(
		srcImageFile ImageFile,
		destImageSet ImageSet,
		destImageFileBaseName BaseName,
	) (
		destImageSetUpdated ImageSet,
		imageFileUpdated ImageFile,
		err error,
	)

	// RemoveTempImageSet removes temporary ImageSet.
	// The arg "tempImageSet" must be a temporary ImageSet.
	RemoveTempImageSet(ImageSet) error
}
