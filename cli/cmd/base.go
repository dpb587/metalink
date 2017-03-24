package cmd

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	blobreceipt "github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/storage"
)

type Meta4 struct {
	Metalink string `long:"metalink" description:"The metalink.meta4 file"`

	StorageFactory storage.StorageFactory `no-flag:"true"`
}

func (f Meta4) Exists() (bool, error) {
	s, err := f.StorageFactory.New(f.Metalink)
	if err != nil {
		return false, bosherr.WrapError(err, "Preparing storage")
	}

	return s.Exists()
}

func (f Meta4) Get() (blobreceipt.Metalink, error) {
	s, err := f.StorageFactory.New(f.Metalink)
	if err != nil {
		return blobreceipt.Metalink{}, bosherr.WrapError(err, "Preparing storage")
	}

	return s.Get()
}

func (f Meta4) Put(put blobreceipt.Metalink) error {
	storage, err := f.StorageFactory.New(f.Metalink)
	if err != nil {
		return bosherr.WrapError(err, "Preparing storage")
	}

	return storage.Put(put)
}
