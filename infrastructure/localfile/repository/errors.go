package repository

import "errors"

var (
	ErrFFProbeIsNotAvailable = errors.New("`ffprobe` may be missing in $PATH")
)
