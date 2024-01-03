package domain_test

import (
	"testing"
	"wallshrink/domain"
)

func TestScaleDown(t *testing.T) {
	dimensionToScaleDown := domain.Dimension{
		Width:  600,
		Height: 400,
	}

	tests := []struct {
		srcWidth   int
		srcHeight  int
		wantWidth  int
		wantHeight int
	}{
		// No change
		{599, 399, 599, 399},
		{600, 400, 600, 400},
		{601, 400, 601, 400},
		{600, 401, 600, 401},
		// Scale down
		{1200, 800, 600, 400},
		{1200, 500, 960, 400},
		{800, 800, 600, 600},
	}

	for i, test := range tests {
		srcDimension := domain.Dimension{
			Width:  test.srcWidth,
			Height: test.srcHeight,
		}
		wantDimension := domain.Dimension{
			Width:  test.wantWidth,
			Height: test.wantHeight,
		}

		gotDimension := srcDimension.ScaleDown(dimensionToScaleDown)
		if gotDimension != wantDimension {
			t.Errorf(
				"No.%d: ScaleDown (%d, %d) -> (%d, %d): Expected (%d, %d), got (%d, %d)\n",
				i+1,
				srcDimension.Width,
				srcDimension.Height,
				dimensionToScaleDown.Width,
				dimensionToScaleDown.Height,
				wantDimension.Width,
				wantDimension.Height,
				gotDimension.Width,
				gotDimension.Height,
			)
		}
	}
}
