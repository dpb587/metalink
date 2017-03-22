package v

import (
	"github.com/Masterminds/semver"
	"github.com/dpb587/blob-receipt/repository"
	"github.com/dpb587/blob-receipt/repository/sorter"
)

type Sorter struct {
	Field string
}

var _ sorter.Sorter = Sorter{}

func (s Sorter) Less(a, b repository.BlobReceipt) bool {
	var av, bv *semver.Version
	var af, bf bool
	var err error

	for _, metadata := range a.Receipt.Metadata {
		if metadata.Key != s.Field {
			continue
		}

		av, err = semver.NewVersion(metadata.Value)
		if err != nil {
			continue
		}

		af = true
	}

	if !af {
		return false
	}

	for _, metadata := range b.Receipt.Metadata {
		if metadata.Key != s.Field {
			continue
		}

		bv, err = semver.NewVersion(metadata.Value)
		if err != nil {
			continue
		}

		bf = true
	}

	if !bf {
		return true
	}

	return !av.LessThan(bv)
}
