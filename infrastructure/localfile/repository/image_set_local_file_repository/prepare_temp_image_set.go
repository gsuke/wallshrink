package image_set_local_file_repository

import (
	"os"
	"wallshrink/domain"
)

func (r *imageSetLocalFileRepository) PrepareTempImageSet() domain.ImageSet {
	return domain.ImageSet{
		Path:                   os.TempDir(),
		BaseNameToImageFileMap: map[string]domain.ImageFile{},
	}
}
