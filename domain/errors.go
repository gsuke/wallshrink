package domain

import "errors"

var (
	ErrImageSetLoadFailed     = errors.New("failed to load image set")
	ErrFileInfoLoadFailed     = errors.New("failed to load file information")
	ErrImageInfoLoadFailed    = errors.New("failed to load image information")
	ErrSSIMCalculateFailed    = errors.New("failed to calculate SSIM")
	ErrIsNotTemporaryImageSet = errors.New("the specified image set is not temporary")
	ErrUnexpected             = errors.New("something went wrong")
	ErrFileNotFound           = errors.New("the specified file could not be found")
)
