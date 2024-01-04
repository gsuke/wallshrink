package domain

import (
	"path/filepath"
	"strings"
)

type BaseName struct {
	Stem      string
	Extension string // includes "."
}

func (b *BaseName) String() string {
	return b.Stem + b.Extension
}

func NewBaseName(filePath string) BaseName {
	basename := filepath.Base(filePath)
	extension := filepath.Ext(filePath)
	stem := strings.TrimSuffix(basename, extension)

	return BaseName{
		Stem:      stem,
		Extension: extension,
	}
}
