package exists

import (
	"github.com/dpb587/metalink"
)

type Filter struct {
	Field string
}

func (f Filter) IsTrue(receipt metalink.BlobReceipt) (bool, error) {
	for _, metadata := range receipt.Metadata {
		if metadata.Key != f.Field {
			continue
		}

		return true, nil
	}

	return false, nil
}
