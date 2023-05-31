package storage

import "github.com/dpb587/metalink"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Storage
type Storage interface {
	String() string

	Exists() (bool, error)
	Get() (metalink.Metalink, error)
	Put(metalink.Metalink) error
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . StorageFactory
type StorageFactory interface {
	New(string) (Storage, error)
}
