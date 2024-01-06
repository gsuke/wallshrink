package domain

type ImageSet struct {
	Path                   string
	BaseNameToImageFileMap map[BaseName]ImageFile // [Basename of the file] -> [ImageFile]
}

// GetDuplicateStemFiles extracts duplicate stem files.
func (s *ImageSet) GetDuplicateStemFiles() (stemToImageFilesMap map[string][]ImageFile) {
	stemToBaseNamesMap := map[string][]BaseName{}
	stemToImageFilesMap = map[string][]ImageFile{}

	for baseName := range s.BaseNameToImageFileMap {
		stemToBaseNamesMap[baseName.Stem] = append(stemToBaseNamesMap[baseName.Stem], baseName)
	}

	for stem, baseNames := range stemToBaseNamesMap {
		// There are duplicate stem files
		if len(baseNames) > 1 {
			for _, baseName := range baseNames {
				stemToImageFilesMap[stem] = append(stemToImageFilesMap[stem], s.BaseNameToImageFileMap[baseName])
			}
		}
	}

	return stemToImageFilesMap
}
