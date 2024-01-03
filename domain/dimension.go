package domain

import "math"

type Dimension struct {
	Width  int
	Height int
}

func (d *Dimension) ScaleDown(dimensionToScaleDown Dimension) Dimension {
	if d.Width <= dimensionToScaleDown.Width || d.Height <= dimensionToScaleDown.Height {
		return *d
	}

	widthRatio := float64(dimensionToScaleDown.Width) / float64(d.Width)
	heightRatio := float64(dimensionToScaleDown.Height) / float64(d.Height)

	if widthRatio == heightRatio {
		return Dimension{
			Width:  int(math.Round(float64(d.Width) * widthRatio)),
			Height: int(math.Round(float64(d.Height) * heightRatio)),
		}
	}

	if widthRatio < heightRatio {
		return Dimension{
			Width:  int(math.Round(float64(d.Width) * heightRatio)),
			Height: dimensionToScaleDown.Height,
		}
	}

	return Dimension{
		Width:  dimensionToScaleDown.Width,
		Height: int(math.Round(float64(d.Height) * widthRatio)),
	}
}
