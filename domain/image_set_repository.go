package domain

type ImageSetRepository interface {
	LoadImageSet(path string) (imageSet ImageSet, warnings []error, err error)
	LoadImageFile(imageSet ImageSet, fileBaseName string) (ImageSet, error)
	PrepareTempImageSet() ImageSet
}
