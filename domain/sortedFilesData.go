package domain

import "time"

var SortedFieldsResult []*SortedFileData

type SortedFileData struct {
	modificationTime string
	data             fileData
}

type fileData struct {
	modificationTime time.Time
	fullPath         string
	extension        string
	size             int64
}

type SortedFilesData struct {
}

func (s *SortedFilesData) SetData(
	filePath string,
	fileTime time.Time,
	extension string,
	size int64,
) *SortedFileData {
	var data = &fileData{
		modificationTime: fileTime,
		fullPath:         filePath,
		extension:        extension,
		size:             size,
	}
	var file = &SortedFileData{
		modificationTime: fileTime.Format("2006-01-02"),
		data:             *data,
	}

	return file
}
