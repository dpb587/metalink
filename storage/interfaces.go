package storage

import "github.com/dpb587/metalink"

//go:generate counterfeiter . Storage
type Storage interface {
	String() string

	Exists() (bool, error)
	Get() (metalink.Metalink, error)
	Put(metalink.Metalink) error
}

//go:generate counterfeiter . StorageFactory
type StorageFactory interface {
	New(string) (Storage, error)
}
