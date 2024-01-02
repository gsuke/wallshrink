package domain

type ImageFile struct {
	Size      int
	Width     int
	Height    int
	Stem      string
	Extension string // includes "."
}
