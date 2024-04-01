package sorter

import (
	"filesorter/domain"
	"fmt"
	"os"
	"path/filepath"
)

type Sorter struct {
}

func (s *Sorter) SortFiles(unsortedFiles []*string) *domain.SortedFilesData {
	var result = &domain.SortedFieldsResult
	var sortedFileSetter = domain.SortedFilesData{}

	for i := 0; i < len(unsortedFiles); i++ {
		var fileStat, errGetFileStat = os.Stat(*unsortedFiles[i])
		if errGetFileStat != nil {
			fmt.Printf("Error getting data for file %s\n", *unsortedFiles[i])
			continue
		}

		var fullPath, errGetFilePath = filepath.Abs(*unsortedFiles[i])
		if errGetFilePath != nil {
			fmt.Printf("Error getting file path for %s\n", *unsortedFiles[i])
			continue
		}

		var sortedFile = sortedFileSetter.SetData(
			fullPath,
			fileStat.ModTime(),
			filepath.Ext(*unsortedFiles[i]),
			fileStat.Size(),
		)

		domain.SortedFieldsResult = append(domain.SortedFieldsResult, sortedFile)
	}

	fmt.Printf("Sorted %d files from %d\n", len(domain.SortedFieldsResult), len(unsortedFiles))

	return domain.SortedFieldsResult
}
