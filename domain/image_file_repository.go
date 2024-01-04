package domain

type ImageFileRepository interface {
	// LoadImageFile loads a image file.
	// Returned ImageFile has no ParentImageSet.
	LoadImageFile(filePath string) (ImageFile, error)
	Compress(srcImageFile ImageFile, destImageFile ImageFile, quality int) (ImageFile, error)
	SSIM(imageFile1 ImageFile, imageFile2 ImageFile) (float64, error)
	RemoveImageFile(filePath string) error
}
