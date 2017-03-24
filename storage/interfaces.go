package storage

import blobreceipt "github.com/dpb587/blob-receipt"

//go:generate counterfeiter . Storage
type Storage interface {
	String() string

	Exists() (bool, error)
	Get() (blobreceipt.Metalink, error)
	Put(blobreceipt.Metalink) error
}

//go:generate counterfeiter . StorageFactory
type StorageFactory interface {
	New(string) (Storage, error)
}
