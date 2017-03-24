package cmd

import (
	"errors"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	blobreceipt "github.com/dpb587/blob-receipt"
)

type Meta4File struct {
	Meta4

	File string `long:"file" description:"The file name"`
}

func (f Meta4File) Get() (blobreceipt.File, error) {
	meta4, err := f.Meta4.Get()
	if err != nil {
		return blobreceipt.File{}, bosherr.WrapError(err, "Preparing storage")
	}

	fileName := f.File

	if fileName == "" && len(meta4.Files) == 1 {
		fileName = meta4.Files[0].Name
	}

	for _, file := range meta4.Files {
		if file.Name != fileName {
			continue
		}

		return file, nil
	}

	return blobreceipt.File{}, errors.New("File does not exist")
}

func (f Meta4File) Put(put blobreceipt.File) error {
	meta4, err := f.Meta4.Get()
	if err != nil {
		return bosherr.WrapError(err, "Preparing storage")
	}

	fileName := f.File

	if fileName == "" && len(meta4.Files) == 1 {
		fileName = meta4.Files[0].Name
	}

	for fileIdx, file := range meta4.Files {
		if file.Name != fileName {
			continue
		}

		meta4.Files[fileIdx] = put

		return f.Meta4.Put(meta4)
	}

	return errors.New("File does not exist")
}
