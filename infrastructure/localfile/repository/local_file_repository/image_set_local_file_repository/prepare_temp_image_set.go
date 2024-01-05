package image_set_local_file_repository

import (
	"os"
	"wompressor/domain"
)

func (r *imageSetLocalFileRepository) PrepareTempImageSet() domain.ImageSet {
	return domain.ImageSet{
		Path:                   os.TempDir(),
		BaseNameToImageFileMap: map[domain.BaseName]domain.ImageFile{},
	}
}
