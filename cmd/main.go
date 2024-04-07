package main

import (
	"filesorter/internal/service/sorter"
	fileStorage "filesorter/internal/storage/file"
	"filesorter/internal/uploader/localFile"
	"fmt"
)

func main() {
	var storage = fileStorage.File{}
	var inputFiles, err = storage.GetFiles()
	if err != nil {
		fmt.Println("Error")
	}
	var fileSorter = sorter.Sorter{} //Why we can't use function SortFiles directly without struct?

	files := fileSorter.SortFiles(inputFiles)

	//for i := 0; i < len(files); i++ {
	//	fmt.Printf("%v\n", files[i])
	//}

	var uploader = localFile.LocalFile{}
	uploader.UploadFiles("./test", files)
}
