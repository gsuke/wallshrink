package domain

type ImageSet struct {
	path       string
	imageFiles []ImageFile
}

type ImageSetRepository interface {
	LoadImageSet(path string) (*ImageSet, []error)
}
