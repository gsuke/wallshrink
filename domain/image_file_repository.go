package domain

type ImageFileRepository interface {
	LoadImageFile(filePath string) (ImageFile, error)
	Compress(srcImageFile ImageFile, destImageFile ImageFile, quality int) (ImageFile, error)
	SSIM(ImageFile, ImageFile) (float64, error)
	RemoveImageFile(ImageFile) error

	// IsFilesSame compares whether 2 files are identical.
	// Returns false if the specified file does not exist.
	IsFilesSame(filePath1 string, filePath2 string) (bool, error)
}
