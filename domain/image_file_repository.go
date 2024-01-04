package domain

type ImageFileRepository interface {
	LoadImageFile(filePath string) (ImageFileParentless, error)
	Compress(srcImageFile ImageFile, destImageFile ImageFile, quality int) (ImageFile, error)
	SSIM(ImageFile, ImageFile) (float64, error)
	RemoveImageFile(ImageFile) error
}
