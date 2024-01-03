package image_file_local_file_repository

import (
	"path/filepath"
	"strings"
)

func splitFileName(path string) (stem string, extension string) {
	basename := filepath.Base(path)
	extension = filepath.Ext(path)
	stem = strings.TrimSuffix(basename, extension)
	return
}
