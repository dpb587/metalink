package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/dpb587/metalink/repository"

	"github.com/pkg/errors"
)

type JSONFormatter struct{}

func (f JSONFormatter) DumpMetalink(metalink repository.RepositoryMetalink) error {
	data, err := json.MarshalIndent(metalink, "", "  ")
	if err != nil {
		return errors.Wrap(err, "Marshaling")
	}

	fmt.Println(string(data))

	return nil
}

func (f JSONFormatter) DumpRepository(metalinks []repository.RepositoryMetalink) error {
	repo := repository.Repository{
		Metalinks: []repository.RepositoryMetalink{},
	}

	for _, meta4 := range metalinks {
		repo.Metalinks = append(repo.Metalinks, meta4)
	}

	data, err := json.MarshalIndent(repo, "", "  ")
	if err != nil {
		return errors.Wrap(err, "Marshaling")
	}

	fmt.Println(string(data))

	return nil
}
