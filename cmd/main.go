package main

import (
	"filesorter/internal/service/sorter"
	fileStorage "filesorter/internal/storage/file"
	"fmt"
)

func main() {
	var storage = fileStorage.File{}
	var inputFiles, err = storage.GetFiles()
	if err != nil {
		fmt.Println("Error")
	}

	files, err := sorter.SortFiles(inputFiles)
	if err != nil {
		return
	}

	fmt.Printf("%v", files)

}
