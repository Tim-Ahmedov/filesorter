package uploader

import "filesorter/domain"

type Uploader interface {
	UploadFiles(string, []*domain.SortedFileData) bool
}
