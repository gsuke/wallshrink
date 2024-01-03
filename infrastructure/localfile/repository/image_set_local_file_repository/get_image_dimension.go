package image_set_local_file_repository

import (
	"context"
	"fmt"
	"time"
	"wallshrink/domain"

	"gopkg.in/vansante/go-ffprobe.v2"
)

func getImageDimension(path string) (width int, height int, err error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: %s", domain.ErrImageInfoLoadFailed, err)
	}

	return data.Streams[0].Width, data.Streams[0].Height, nil
}
