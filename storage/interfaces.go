package storage

import blobreceipt "github.com/dpb587/blob-receipt"

//go:generate counterfeiter . Storage
type Storage interface {
	String() string

	Exists() (bool, error)
	Get() (blobreceipt.BlobReceipt, error)
	Put(blobreceipt.BlobReceipt) error
}

//go:generate counterfeiter . StorageFactory
type StorageFactory interface {
	New(string) (Storage, error)
}
