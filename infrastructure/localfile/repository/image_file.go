package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	"wallshrink/domain"

	"github.com/samber/do"
	"gopkg.in/vansante/go-ffprobe.v2"
)

type imageFileLocalFileRepository struct{}

func NewImageFileLocalFileRepository(i *do.Injector) (domain.ImageFileRepository, error) {
	if !isFFProbeAvailable() {
		fmt.Println(ErrFFProbeIsNotAvailable)
		os.Exit(1)
	}
	return &imageFileLocalFileRepository{}, nil
}

func (r *imageFileLocalFileRepository) LoadImageFile(imageSet domain.ImageSet, filename string) (domain.ImageFile, error) {
	stem, extension := splitFileName(filename)

	imageFile := domain.ImageFile{
		Stem:           stem,
		Extension:      extension,
		ParentImageSet: imageSet,
	}

	// Get file size
	size, err := getFileSize(imageFile.FullPath())
	if err != nil {
		return domain.ImageFile{}, err
	}
	imageFile.Size = size

	// Get image dimension
	width, height, err := getImageDimension(imageFile.FullPath())
	if err != nil {
		return domain.ImageFile{}, err
	}
	imageFile.Width = width
	imageFile.Height = height

	return imageFile, nil
}

func splitFileName(path string) (stem string, extension string) {
	basename := filepath.Base(path)
	extension = filepath.Ext(path)
	stem = strings.TrimSuffix(basename, extension)
	return
}

// isFFProbeAvailable Checks if `ffprobe` is in $PATH.
func isFFProbeAvailable() bool {
	// ffprobe.SetFFProbeBinPath("foo") // For Testing

	ctx, cancelFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFn()

	_, err := ffprobe.ProbeURL(ctx, "")
	return !errors.Is(err, exec.ErrNotFound)
}

func getFileSize(path string) (int, error) {
	// Open file
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", domain.ErrFileInfoLoadFailed, err)
	}
	defer file.Close()

	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("%w: %s", domain.ErrFileInfoLoadFailed, err)
	}

	return int(fileInfo.Size()), nil
}

func getImageDimension(path string) (width int, height int, err error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %s", domain.ErrImageInfoLoadFailed, err)
	}

	return data.Streams[0].Width, data.Streams[0].Height, nil
}
