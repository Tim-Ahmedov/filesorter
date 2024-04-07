package sorter

import (
	"filesorter/domain"
	"fmt"
	"os"
	"path/filepath"
)

type Sorter struct {
}

func (s *Sorter) SortFiles(unsortedFiles []*string) []*domain.SortedFileData {
	var SortedFileDataResult []*domain.SortedFileData

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

		var sortedFile = domain.SortedFileData{
			ModificationTime: fileStat.ModTime().Format("2006-01-02"),
			Data: domain.FileData{
				ModificationTime: fileStat.ModTime(),
				FullPath:         fullPath,
				Extension:        filepath.Ext(*unsortedFiles[i]),
				Size:             fileStat.Size(),
			},
		}

		SortedFileDataResult = append(SortedFileDataResult, &sortedFile)
	}

	fmt.Printf("Sorted %d files from %d\n", len(SortedFileDataResult), len(unsortedFiles))

	return SortedFileDataResult
}
