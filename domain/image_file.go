package domain

type ImageFile struct {
	Size      int
	Width     int
	Height    int
	BaseName  string
	Extension string // includes "."
}
