package blobreceipt

import (
	"encoding/json"
	"io"
	"sort"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/blob-receipt/crypto"
	"github.com/dpb587/blob-receipt/origin"
)

type BlobReceipt struct {
	Digest   BlobReceiptDigest     `json:"digest,omitempty"`
	Metadata []BlobReceiptMetadata `json:"metadata,omitempty"`
	Name     string                `json:"name"`
	Origin   []BlobReceiptOrigin   `json:"origin,omitempty"`
	Size     uint64                `json:"size,omitempty"`
	Time     time.Time             `json:"time,omitempty"`
}

type BlobReceiptMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (r *BlobReceipt) SetMetadata(key, value string) {
	for metadataIdx, metadata := range r.Metadata {
		if metadata.Key != key {
			continue
		}

		r.Metadata[metadataIdx].Value = value

		return
	}

	r.Metadata = append(r.Metadata, BlobReceiptMetadata{key, value})
}

func (r *BlobReceipt) SetOrigin(origin BlobReceiptOrigin) {
	for originIdx, existingOrigin := range r.Origin {
		if existingOrigin.URI() != origin.URI() {
			continue
		}

		r.Origin[originIdx] = origin

		return
	}

	r.Origin = append(r.Origin, origin)
}

func (r BlobReceipt) Write(w io.Writer) error {
	sort.Sort(blobReceiptMetadataByKey(r.Metadata))
	sort.Sort(blobReceiptOriginByURI(r.Origin))

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return bosherr.WrapError(err, "Marshaling JSON")
	}

	_, err = w.Write(data)
	if err != nil {
		return bosherr.WrapError(err, "Writing JSON")
	}

	w.Write([]byte("\n"))

	return nil
}

func (r *BlobReceipt) UpdateFromOrigin(blob origin.Origin, algorithms []string) error {
	blobName, err := blob.Name()
	if err != nil {
		return bosherr.WrapError(err, "Sourcing blob name")
	}

	r.Name = blobName

	blobSize, err := blob.Size()
	if err != nil {
		return bosherr.WrapError(err, "Sourcing blob size")
	}

	r.Size = blobSize

	blobTime, err := blob.Time()
	if err != nil {
		return bosherr.WrapError(err, "Sourcing blob time")
	}

	r.Time = blobTime

	if r.Digest == nil {
		r.Digest = BlobReceiptDigest{}
	}

	for _, algorithmName := range algorithms {
		algorithm, err := crypto.GetAlgorithm(algorithmName)
		if err != nil {
			return bosherr.WrapErrorf(err, "Loading digest algorithm")
		}

		digest, err := blob.Digest(algorithm)
		if err != nil {
			return bosherr.WrapErrorf(err, "Sourcing blob %s digest", algorithm.Name())
		}

		r.Digest[algorithm.Name()] = crypto.GetDigestHash(digest)
	}

	return nil
}
