package domain

import "math"

type Dimension struct {
	Width  int
	Height int
}

func (d *Dimension) ScaleDown(scaleDownDimension Dimension) Dimension {
	if d.Width <= scaleDownDimension.Width || d.Height <= scaleDownDimension.Height {
		return *d
	}

	widthRatio := float64(scaleDownDimension.Width) / float64(d.Width)
	heightRatio := float64(scaleDownDimension.Height) / float64(d.Height)

	if widthRatio == heightRatio {
		return Dimension{
			Width:  int(math.Round(float64(d.Width) * widthRatio)),
			Height: int(math.Round(float64(d.Height) * heightRatio)),
		}
	}

	if widthRatio < heightRatio {
		return Dimension{
			Width:  int(math.Round(float64(d.Width) * heightRatio)),
			Height: scaleDownDimension.Height,
		}
	}

	return Dimension{
		Width:  scaleDownDimension.Width,
		Height: int(math.Round(float64(d.Height) * widthRatio)),
	}
}
