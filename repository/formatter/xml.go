package formatter

import (
	"encoding/xml"
	"fmt"

	"github.com/dpb587/metalink/repository"

	"github.com/pkg/errors"
)

type XMLFormatter struct{}

func (f XMLFormatter) DumpMetalink(metalink repository.RepositoryMetalink) error {
	fmt.Println(`<?xml version="1.0" encoding="utf-8"?>`)

	data, err := xml.MarshalIndent(metalink, "", "  ")
	if err != nil {
		return errors.Wrap(err, "Marshaling")
	}

	fmt.Println(string(data))

	return nil
}

func (f XMLFormatter) DumpRepository(metalinks []repository.RepositoryMetalink) error {
	fmt.Println(`<?xml version="1.0" encoding="utf-8"?>`)

	repo := repository.Repository{
		Metalinks: []repository.RepositoryMetalink{},
	}

	for _, meta4 := range metalinks {
		repo.Metalinks = append(repo.Metalinks, meta4)
	}

	data, err := xml.MarshalIndent(repo, "", "  ")
	if err != nil {
		return errors.Wrap(err, "Marshaling")
	}

	fmt.Println(string(data))

	return nil
}
