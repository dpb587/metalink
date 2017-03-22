package v

import (
	"github.com/Masterminds/semver"
	blobreceipt "github.com/dpb587/blob-receipt"
)

type Filter struct {
	Field      string
	Constraint semver.Constraints
}

func (f Filter) IsTrue(receipt blobreceipt.BlobReceipt) (bool, error) {
	for _, metadata := range receipt.Metadata {
		if metadata.Key != f.Field {
			continue
		}

		version, err := semver.NewVersion(metadata.Value)
		if err != nil {
			continue
		}

		return f.Constraint.Check(version), nil
	}

	return false, nil
}
