package domain

type ImageSetRepository interface {
	LoadImageSet(path string) (*ImageSet, []error)
}
