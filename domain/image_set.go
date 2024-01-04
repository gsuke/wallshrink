package domain

type ImageSet struct {
	Path                   string
	BaseNameToImageFileMap map[BaseName]ImageFile // [Basename of the file] -> [ImageFile]
}
