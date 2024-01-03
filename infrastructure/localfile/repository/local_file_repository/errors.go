package repository

import "errors"

var (
	ErrFFProbeIsNotAvailable = errors.New("`ffprobe` may be missing in $PATH")
	ErrFFmpegIsNotAvailable  = errors.New("`ffmpeg` may be missing in $PATH")
)
