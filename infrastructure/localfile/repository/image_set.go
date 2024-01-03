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

type imageSetLocalFileRepository struct{}

func NewImageSetLocalFileRepository(i *do.Injector) (domain.ImageSetRepository, error) {
	if !isFFProbeAvailable() {
		fmt.Println(ErrFFProbeIsNotAvailable)
		os.Exit(1)
	}

	return &imageSetLocalFileRepository{}, nil
}

func (r *imageSetLocalFileRepository) LoadImageSet(path string) (imageSet domain.ImageSet, warnings []error, err error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return domain.ImageSet{}, nil, fmt.Errorf("%w: %s", domain.ErrImageSetLoadFailed, err)
	}

	imageSet = domain.ImageSet{
		Path:                   path,
		BaseNameToImageFileMap: map[string]domain.ImageFile{},
	}

	// Load all ImageFiles in ImageSet
	for i, f := range files {
		NewImageSet, err := r.LoadImageFile(imageSet, f.Name())

		fmt.Printf("%d/%d: %s\n", i+1, len(files), filepath.Join(path, f.Name()))

		if err != nil {
			warnings = append(warnings, err)
			fmt.Println("  ↑ [!] Failed to load image information. The directory should contain only image files.")
			continue
		}

		imageSet = NewImageSet

	}

	return imageSet, warnings, nil
}

func (r *imageSetLocalFileRepository) LoadImageFile(imageSet domain.ImageSet, fileBaseName string) (domain.ImageSet, error) {
	stem, extension := splitFileName(fileBaseName)

	imageFile := domain.ImageFile{
		Stem:           stem,
		Extension:      extension,
		ParentImageSet: imageSet,
	}

	// Get file size
	size, err := getFileSize(imageFile.FullPath())
	if err != nil {
		return domain.ImageSet{}, err
	}
	imageFile.Size = size

	// Get image dimension
	width, height, err := getImageDimension(imageFile.FullPath())
	if err != nil {
		return domain.ImageSet{}, err
	}
	imageFile.Width = width
	imageFile.Height = height

	imageSet.BaseNameToImageFileMap[imageFile.BaseName()] = imageFile

	return imageSet, nil
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
