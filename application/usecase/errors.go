package usecase

import "errors"

var (
	ErrSrcAndDestAreSame          = errors.New("source and destination image sets must be different")
	ErrSSIMShortage               = errors.New("SSIM did not reach the pass line")
	ErrThereAreDuplicateStemFiles = errors.New("there are duplicate stem files")
)
