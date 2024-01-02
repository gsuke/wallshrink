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

func (r *imageFileLocalFileRepository) LoadImageFile(path string) (domain.ImageFile, error) {
	stem, extension := splitFileName(path)

	return domain.ImageFile{
		// TODO
		Size:      0,
		Width:     0,
		Height:    0,
		Stem:      stem,
		Extension: extension,
	}, nil
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
