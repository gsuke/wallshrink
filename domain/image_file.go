package domain

type ImageFile struct {
	Size           int
	Dimension      Dimension
	Stem           string
	Extension      string // includes "."
	ParentImageSet ImageSet
}

func (f *ImageFile) BaseName() string {
	return f.Stem + f.Extension
}
