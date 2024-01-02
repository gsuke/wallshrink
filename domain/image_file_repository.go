package domain

type ImageFileRepository interface {
	LoadImageFile(path string) (ImageFile, error)
}
