package image_file_local_file_repository

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"wallshrink/domain"

	"github.com/google/uuid"
)

func (r *imageFileLocalFileRepository) SSIM(imageFile1 domain.ImageFile, imageFile2 domain.ImageFile) (float64, error) {

	tempFileName := strings.Replace(uuid.NewString(), "-", "", -1) + ".txt"
	defer os.Remove(tempFileName)

	err := exec.Command(
		"ffmpeg",
		"-i",
		imageFile1.FullPath(),
		"-i",
		imageFile2.FullPath(),
		"-filter_complex",
		fmt.Sprintf("scale2ref,ssim=f=%s", tempFileName),
		"-an",
		"-f",
		"null",
		"-",
	).Run()
	if err != nil {
		return 0, fmt.Errorf("%w: %s", domain.ErrSSIMCalculateFailed, err)
	}

	ssimOutputText, err := os.ReadFile(tempFileName)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", domain.ErrSSIMCalculateFailed, err)
	}

	ssim, err := parseFFmpegSSIMText(string(ssimOutputText))
	if err != nil {
		return 0, err
	}

	return ssim, nil
}

// parseFFmpegSSIMText parses SSIM output from ffmpeg.
// e.g. "n:1 Y:0.963481 U:0.963360 V:0.931346 All:0.958104 (13.778228)"
func parseFFmpegSSIMText(ssimText string) (float64, error) {
	errorTemplate := fmt.Errorf("%w: SSIM output from ffmpeg is wrong: \"%s\"", domain.ErrSSIMCalculateFailed, ssimText)

	for _, keyValueText := range strings.Split(ssimText, " ") {
		splitKeyValueText := strings.Split(keyValueText, ":")
		if len(splitKeyValueText) != 2 {
			return 0, errorTemplate
		}

		key := splitKeyValueText[0]
		value := splitKeyValueText[1]

		if key == "All" {
			ssim, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return 0, errorTemplate
			}
			return ssim, nil
		}
	}

	return 0, errorTemplate
}
