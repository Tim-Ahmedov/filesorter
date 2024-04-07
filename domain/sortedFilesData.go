package domain

import "time"

type SortedFileData struct {
	ModificationTime string
	Data             FileData
}

type FileData struct {
	ModificationTime time.Time
	FullPath         string
	Extension        string
	Size             int64
}

type SortedFilesData struct {
}
