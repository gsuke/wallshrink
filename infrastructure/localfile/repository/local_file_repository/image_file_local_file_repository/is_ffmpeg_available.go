package image_file_local_file_repository

import (
	"errors"
	"os/exec"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// isFFmpegAvailable Checks if `ffmpeg` is in $PATH.
func isFFmpegAvailable() bool {
	// err := ffmpeg.Input("").SetFfmpegPath("foo").Run() // For Testing
	err := ffmpeg.Input("").Run()
	return !errors.Is(err, exec.ErrNotFound)
}
