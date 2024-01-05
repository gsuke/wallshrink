package image_set_local_file_repository

import (
	"io"
	"os"
	"path/filepath"
	"wompressor/domain"
)

func (r *imageSetLocalFileRepository) CopyImageFile(
	srcImageFile domain.ImageFile,
	destImageSet domain.ImageSet,
	destImageFileBaseName domain.BaseName,
) (
	destImageSetUpdated domain.ImageSet,
	imageFileUpdated domain.ImageFile,
	err error,
) {

	// Copy file
	err = copyFile(
		srcImageFile.FullPath(),
		filepath.Join(destImageSet.Path, destImageFileBaseName.String()),
	)
	if err != nil {
		return domain.ImageSet{}, domain.ImageFile{}, err
	}

	// Update ImageFile
	imageFileUpdated = srcImageFile
	imageFileUpdated.ParentImageSetPath = destImageSet.Path
	imageFileUpdated.BaseName = destImageFileBaseName

	// Update destImageSet
	destImageSetUpdated = destImageSet.DeepCopy()
	destImageSetUpdated.BaseNameToImageFileMap[imageFileUpdated.BaseName] = imageFileUpdated

	return destImageSetUpdated, imageFileUpdated, nil
}

func copyFile(srcFilePath string, destFilePath string) error {
	src, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
