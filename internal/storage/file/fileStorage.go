package fileStorage

import (
	"filesorter/internal/storage"
	"os"
	"path/filepath"
)

type File struct {
	files storage.Storage
}

func (f *File) GetFiles() ([]*string, error) {
	entries, err := os.ReadDir("./")
	if err != nil {
		return nil, err
	}

	var unsortedFields []*string

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		var absPath, _ = filepath.Abs(e.Name())
		unsortedFields = append(unsortedFields, &absPath)
	}

	return unsortedFields, nil
}
