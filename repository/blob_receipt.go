package repository

import "github.com/dpb587/metalink"

type BlobReceipt struct {
	Repository BlobReceiptRepository   `json:"repository"`
	Receipt    metalink.BlobReceipt `json:"receipt"`
}

type BlobReceiptRepository struct {
	URI     string `json:"uri"`
	Version string `json:"version"`
	Path    string `json:"path"`
}
