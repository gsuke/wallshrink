package domain

import "errors"

var (
	ErrDirectoryLoadFailed = errors.New("failed to load directory")
	ErrFileInfoLoadFailed  = errors.New("failed to load file information")
	ErrImageInfoLoadFailed = errors.New("failed to load image information")
)
