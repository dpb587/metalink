package source

import (
	"github.com/pkg/errors"
	"github.com/dpb587/metalink/repository"
	"github.com/dpb587/metalink/repository/filter"
)

func FilterInMemory(files []repository.RepositoryMetalink, filter filter.Filter) ([]repository.RepositoryMetalink, error) {
	results := []repository.RepositoryMetalink{}

	for _, meta4 := range files {
		matched, err := filter.IsTrue(meta4)
		if err != nil {
			return nil, errors.Wrap(err, "Matching metalink")
		} else if !matched {
			continue
		}

		results = append(results, meta4)
	}

	return results, nil
}
