package image_file_local_file_repository

import (
	"context"
	"errors"
	"os/exec"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

// isFFProbeAvailable Checks if `ffprobe` is in $PATH.
func isFFProbeAvailable() bool {
	// ffprobe.SetFFProbeBinPath("foo") // For Testing

	ctx, cancelFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFn()

	_, err := ffprobe.ProbeURL(ctx, "")
	return !errors.Is(err, exec.ErrNotFound)
}
