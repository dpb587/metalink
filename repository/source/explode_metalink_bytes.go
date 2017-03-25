package source

import (
	"encoding/xml"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/repository"
)

func ExplodeMetalinkBytes(repo repository.Repository, metalinkBytes []byte) ([]repository.File, error) {
	results := []repository.File{}

	metalinkParsed := metalink.Metalink{}

	err := xml.Unmarshal(metalinkBytes, &metalinkParsed)
	if err != nil {
		return nil, bosherr.WrapError(err, "Parsing metalink")
	}

	metalinkAbbr := metalink.Metalink{
		Generator: metalinkParsed.Generator,
		Origin:    metalinkParsed.Origin,
		Published: metalinkParsed.Published,
		Updated:   metalinkParsed.Updated,
	}

	for _, metalinkFile := range metalinkParsed.Files {
		repositoryFile := repository.File{
			Repository: repo,
			Metalink:   metalinkAbbr,
			File:       metalinkFile,
		}

		results = append(results, repositoryFile)
	}

	return results, nil
}
