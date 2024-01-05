package usecase

import "errors"

var (
	ErrSrcAndDestAreSame = errors.New("source and destination image sets must be different")
)
