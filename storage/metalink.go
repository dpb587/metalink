package storage

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/blob-receipt"
)

func ReadMetalink(r io.Reader) (blobreceipt.Metalink, error) {
	// sort.Sort(blobReceiptMetadataByKey(r.Metadata))
	// sort.Sort(blobReceiptOriginByURI(r.Origin))

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return blobreceipt.Metalink{}, bosherr.WrapError(err, "Reading XML")
	}

	meta4 := blobreceipt.Metalink{}

	err = xml.Unmarshal(data, &meta4)
	if err != nil {
		return blobreceipt.Metalink{}, bosherr.WrapError(err, "Unmarshaling XML")
	}

	return meta4, nil
}

func WriteMetalink(w io.Writer, r blobreceipt.Metalink) error {
	// sort.Sort(blobReceiptMetadataByKey(r.Metadata))
	// sort.Sort(blobReceiptOriginByURI(r.Origin))

	data, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return bosherr.WrapError(err, "Marshaling XML")
	}

	w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?>`))
	w.Write([]byte("\n"))

	_, err = w.Write(data)
	if err != nil {
		return bosherr.WrapError(err, "Writing XML")
	}

	w.Write([]byte("\n"))

	return nil
}
