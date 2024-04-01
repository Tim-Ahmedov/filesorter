package uploader

import "filesorter/internal/service/sorter"

type Uploader interface {
	UploadFiles() ([]*sorter.SortedFilesData, error)
}
