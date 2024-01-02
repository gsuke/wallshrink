package domain

import "path/filepath"

type ImageFile struct {
	Size           int
	Width          int
	Height         int
	Stem           string
	Extension      string // includes "."
	ParentImageSet ImageSet
}

func (f *ImageFile) FullPath() string {
	return filepath.Join(f.ParentImageSet.Path, f.Stem+f.Extension)
}
