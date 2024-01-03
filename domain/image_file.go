package domain

import "path/filepath"

type ImageFile struct {
	Size           int
	Dimension      Dimension
	Stem           string
	Extension      string // includes "."
	ParentImageSet ImageSet
}

func (f *ImageFile) FullPath() string {
	return filepath.Join(f.ParentImageSet.Path, f.BaseName())
}

func (f *ImageFile) BaseName() string {
	return f.Stem + f.Extension
}
