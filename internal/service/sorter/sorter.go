package sorter

import (
	"os"
	"path/filepath"
)

type fileData struct {
	modificationTime string
	fullPath         string
	extension        string
	size             int64
}

func SortFiles(unsortedFiles []*string) ([]fileData, error) {
	var sortedFieldsResult []fileData

	for i := 0; i < len(unsortedFiles); i++ {
		var fileStat, err = os.Stat(*unsortedFiles[i])
		if err != nil {
			continue
		}

		var fullPath, _ = filepath.Abs(*unsortedFiles[i])

		var fileData = fileData{
			modificationTime: fileStat.ModTime().String(),
			fullPath:         fullPath,
			extension:        filepath.Ext(*unsortedFiles[i]),
			size:             fileStat.Size(),
		}

		sortedFieldsResult = append(sortedFieldsResult, fileData)
	}

	return sortedFieldsResult, nil
}
