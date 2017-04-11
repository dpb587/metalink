package formatter

import (
	"fmt"

	"github.com/dpb587/metalink/repository"
)

type VersionFormatter struct{}

func (f VersionFormatter) DumpMetalink(metalink repository.RepositoryMetalink) error {
	fmt.Println(metalink.Files[0].Version)

	return nil
}

func (f VersionFormatter) DumpRepository(metalinks []repository.RepositoryMetalink) error {
	for _, meta4 := range metalinks {
		fmt.Println(meta4.Metalink.Files[0].Version)
	}

	return nil
}
