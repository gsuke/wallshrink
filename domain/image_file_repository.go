package domain

type ImageFileRepository interface {
	LoadImageFile(imageSet ImageSet, filename string) (ImageFile, error)
}
