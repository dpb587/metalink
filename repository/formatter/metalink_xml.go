package formatter

import (
	"encoding/xml"
	"fmt"

	"github.com/dpb587/metalink/repository"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type MetalinkXMLFormatter struct{}

func (f MetalinkXMLFormatter) DumpMetalink(metalink repository.RepositoryMetalink) error {
	fmt.Println(`<?xml version="1.0" encoding="utf-8"?>`)

	data, err := xml.MarshalIndent(metalink, "", "  ")
	if err != nil {
		return bosherr.WrapError(err, "Marshaling")
	}

	fmt.Println(string(data))

	return nil
}

func (f MetalinkXMLFormatter) DumpRepository(metalinks []repository.RepositoryMetalink) error {
	for _, meta4 := range metalinks {
		err := f.DumpMetalink(meta4)
		if err != nil {
			return err
		}
	}

	return nil
}
