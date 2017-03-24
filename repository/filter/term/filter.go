package term

import (
	"github.com/dpb587/metalink"
)

type Filter struct {
	Field string
	Value string
}

func (f Filter) IsTrue(receipt metalink.BlobReceipt) (bool, error) {
	for _, metadata := range receipt.Metadata {
		if metadata.Key != f.Field {
			continue
		} else if metadata.Value != f.Value {
			continue
		}

		return true, nil
	}

	return false, nil
}
