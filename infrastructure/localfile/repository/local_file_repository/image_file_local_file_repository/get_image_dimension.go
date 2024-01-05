package image_file_local_file_repository

import (
	"context"
	"fmt"
	"time"
	"wompressor/domain"

	"gopkg.in/vansante/go-ffprobe.v2"
)

func getImageDimension(path string) (domain.Dimension, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil {
		return domain.Dimension{}, fmt.Errorf("%w: %s", domain.ErrImageInfoLoadFailed, err)
	}

	dimension := domain.Dimension{
		Width:  data.Streams[0].Width,
		Height: data.Streams[0].Height,
	}

	return dimension, nil
}
