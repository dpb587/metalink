package fs

import (
	"encoding/json"
	"fmt"
	"path"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/blob-receipt"
	"github.com/dpb587/blob-receipt/repository"
	"github.com/dpb587/blob-receipt/repository/filter"
	"github.com/dpb587/blob-receipt/repository/source"
)

type Source struct {
	uri  string
	fs   boshsys.FileSystem
	path string

	receipts []repository.BlobReceipt
}

var _ source.Source = &Source{}

func NewSource(uri string, fs boshsys.FileSystem, path string) *Source {
	return &Source{
		uri:  uri,
		fs:   fs,
		path: path,
	}
}

func (s *Source) Reload() error {
	files, err := s.fs.Glob(fmt.Sprintf("%s/*.json", s.path))
	if err != nil {
		return bosherr.WrapError(err, "Listing receipts")
	}

	s.receipts = []repository.BlobReceipt{}

	for _, file := range files {
		stat, err := s.fs.Stat(file)
		if err != nil {
			return bosherr.WrapError(err, "Stat receipt")
		}

		receiptBytes, err := s.fs.ReadFile(file)
		if err != nil {
			return bosherr.WrapError(err, "Reading receipt")
		}

		receipt := blobreceipt.BlobReceipt{}

		err = json.Unmarshal(receiptBytes, &receipt)
		if err != nil {
			return bosherr.WrapError(err, "Parsing receipt")
		}

		annotatedreceipt := repository.BlobReceipt{
			Repository: repository.BlobReceiptRepository{
				URI:     s.URI(),
				Path:    path.Base(file),
				Version: stat.ModTime().Format(time.RFC3339),
			},
			Receipt: receipt,
		}

		s.receipts = append(s.receipts, annotatedreceipt)
	}

	return nil
}

func (s Source) URI() string {
	return s.uri
}

func (s Source) FilterBlobReceipts(filter filter.Filter) ([]repository.BlobReceipt, error) {
	matches := []repository.BlobReceipt{}

	for _, receipt := range s.receipts {
		matched, err := filter.IsTrue(receipt.Receipt)
		if err != nil {
			return nil, bosherr.WrapError(err, "Matching receipt")
		} else if !matched {
			continue
		}

		matches = append(matches, receipt)
	}

	return matches, nil
}
