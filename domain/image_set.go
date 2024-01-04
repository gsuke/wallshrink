package domain

type ImageSet struct {
	Path                   string
	BaseNameToImageFileMap map[string]ImageFile // [Basename of the file] -> [ImageFile]
}
