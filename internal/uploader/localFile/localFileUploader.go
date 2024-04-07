package localFile

import (
	"filesorter/domain"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalFile struct {
}

func (lf *LocalFile) UploadFiles(destinationDir string, sortedFilesData []*domain.SortedFileData) bool {
	err := os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating destination dir!")
		return false
	}

	for i := 0; i < len(sortedFilesData); i++ {
		err := os.MkdirAll(
			filepath.Join(destinationDir, sortedFilesData[i].ModificationTime),
			os.ModePerm,
		)
		if err != nil {
			fmt.Println("Error creating date sub dir!")
			continue
		}

		var dst = filepath.Join(
			destinationDir,
			sortedFilesData[i].ModificationTime,
			filepath.Base(sortedFilesData[i].Data.FullPath),
		)

		source, err := os.Open(sortedFilesData[i].Data.FullPath)
		if err != nil {
			fmt.Println("Failure to create source reader!")
			continue
		}
		defer source.Close()

		destination, err := os.Create(dst)
		if err != nil {
			fmt.Println("Failure to create destination writer!")
		}

		nBytes, err := io.Copy(destination, source)
		if err != nil {
			return false
		}

		fmt.Printf(
			"Copied %d bites from %d for file %s \n",
			nBytes,
			sortedFilesData[i].Data.Size,
			filepath.Base(sortedFilesData[i].Data.FullPath),
		)
	}

	return true
}
