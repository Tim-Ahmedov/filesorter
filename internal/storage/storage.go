package storage

type Storage interface {
	GetFiles() ([]*string, error)
}
