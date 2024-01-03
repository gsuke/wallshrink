package domain

type ImageSetRepository interface {
	LoadImageSet(path string) (imageSet ImageSet, warnings []error, err error)
	PrepareTempImageSet() ImageSet
}
