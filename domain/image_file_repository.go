package domain

type ImageFileRepository interface {
	// LoadImageFile loads a image file.
	// Returned ImageFile has no ParentImageSet.
	LoadImageFile(fileBaseName string) (ImageFile, error)
	Compress(srcImageFile ImageFile, destImageFile ImageFile) (ImageFile, error)
}
