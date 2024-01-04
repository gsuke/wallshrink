package image_file_local_file_repository

import (
	"errors"
	"testing"
	"wallshrink/domain"
)

func Test_parseFFmpegSSIMText(t *testing.T) {
	tests := []struct {
		name       string
		ssimText   string
		wantResult float64
		wantErr    error
	}{
		// Success
		{"success1", "n:1 Y:0.963481 U:0.963360 V:0.931346 All:0.958104 (13.778228)", 0.958104, nil},
		// Fail
		{"fail1", "a:0.90000:b c:0.90000:d e:0.90000:f", 0, domain.ErrSSIMCalculateFailed},
		{"fail2", "n:1 Y:0.963481 U:0.963360 V:0.931346 All:foo (13.778228)", 0, domain.ErrSSIMCalculateFailed},
		{"fail3", "n:1 Y:0.963481 U:0.963360 V:0.931346 (13.778228)", 0, domain.ErrSSIMCalculateFailed},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseFFmpegSSIMText(test.ssimText)
			if err != nil && !errors.Is(err, test.wantErr) {
				t.Errorf("\"%s\": error = %v, wantErr %v", test.ssimText, err, test.wantErr)
				return
			}
			if got != test.wantResult {
				t.Errorf("\"%s\": got %v, want %v", test.ssimText, got, test.wantResult)
			}
		})
	}
}
