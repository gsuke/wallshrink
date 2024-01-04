package image_file_local_file_repository

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"wallshrink/domain"
)

func (r *imageFileLocalFileRepository) IsFilesSame(filePath1 string, filePath2 string) (bool, error) {
	result, err := isFilesSame(filePath1, filePath2)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrFileNotFound):
			return false, nil
		default:
			return false, err
		}
	}
	return result, nil
}

const chunkSize = 64000

func isFilesSame(filePath1 string, filePath2 string) (bool, error) {
	file1, err := os.Open(filePath1)
	if err != nil {
		return false, fmt.Errorf("%w: %s", domain.ErrFileNotFound, filePath1)
	}
	defer file1.Close()

	file2, err := os.Open(filePath2)
	if err != nil {
		return false, fmt.Errorf("%w: %s", domain.ErrFileNotFound, filePath2)
	}
	defer file2.Close()

	for {
		bytes1 := make([]byte, chunkSize)
		_, err1 := file1.Read(bytes1)

		bytes2 := make([]byte, chunkSize)
		_, err2 := file2.Read(bytes2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true, nil
			}
			if err1 == io.EOF || err2 == io.EOF {
				return false, nil
			}
			return false, fmt.Errorf("%w: comparing \"%s\" and \"%s\"", domain.ErrUnexpected, filePath1, filePath2)
		}

		if !bytes.Equal(bytes1, bytes2) {
			return false, nil
		}
	}

}
