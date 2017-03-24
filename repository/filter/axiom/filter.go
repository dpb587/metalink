package axiom

import "github.com/dpb587/metalink"

type Filter struct{}

func (Filter) IsTrue(_ metalink.BlobReceipt) (bool, error) {
	return true, nil
}
