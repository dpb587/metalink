package repository

import blobreceipt "github.com/dpb587/blob-receipt"

type BlobReceipt struct {
	Repository BlobReceiptRepository   `json:"repository"`
	Receipt    blobreceipt.BlobReceipt `json:"receipt"`
}

type BlobReceiptRepository struct {
	URI     string `json:"uri"`
	Version string `json:"version"`
	Path    string `json:"path"`
}
