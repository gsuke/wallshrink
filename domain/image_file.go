package domain

type ImageFile struct {
	size      int
	width     int
	height    int
	baseName  string
	extension string // includes "."
}

type ImageFileRepository interface {
	LoadImageFile(path string) (ImageFile, error)
}
