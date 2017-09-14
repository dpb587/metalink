package cmd

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/storage"
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

func (f Meta4) Get() (metalink.Metalink, error) {
	s, err := f.StorageFactory.New(f.Metalink)
	if err != nil {
		return metalink.Metalink{}, bosherr.WrapError(err, "Preparing storage")
	}

	return s.Get()
}

func (f Meta4) Put(put metalink.Metalink) error {
	storage, err := f.StorageFactory.New(f.Metalink)
	if err != nil {
		return bosherr.WrapError(err, "Preparing storage")
	}

	metalink.Sort(&put)

	return storage.Put(put)
}
