package source

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

func FilterFilesInMemory(files []repository.File, filter filter.Filter) ([]repository.File, error) {
	results := []repository.File{}

	for _, metalinkFile := range files {
		matched, err := filter.IsTrue(metalinkFile)
		if err != nil {
			return nil, bosherr.WrapError(err, "Matching metalink file")
		} else if !matched {
			continue
		}

		results = append(results, metalinkFile)
	}

	return results, nil
}
