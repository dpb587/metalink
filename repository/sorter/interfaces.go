package sorter

import (
	"github.com/dpb587/blob-receipt/repository"
)

type Sorter interface {
	Less(repository.BlobReceipt, repository.BlobReceipt) bool
}
