package verify

import (
	"errors"
	"io"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/dpb587/metalink"
	"github.com/dpb587/metalink/crypto"
	"github.com/dpb587/metalink/origin"
	"github.com/dpb587/metalink/signature"
)

type VerifyResult struct {
	Type     string
	Expected string
	Actual   string
	Error    error
}

func (vr VerifyResult) HasError() bool {
	return vr.Error != nil
}

type MultipleVerifyResult []VerifyResult

func (vr MultipleVerifyResult) HasError() bool {
	for _, result := range vr {
		if result.HasError() {
			return true
		}
	}

	return false
}

type Verifier struct{}

func (v Verifier) VerifyHash(expected metalink.File, actual origin.Origin, hashAlgorithm string) VerifyResult {
	var expectedHash *metalink.Hash
	var result = VerifyResult{
		Type:     hashAlgorithm,
		Expected: "UNKNOWN",
		Actual:   "UNKNOWN",
	}

	for _, seekHash := range expected.Hashes {
		if seekHash.Type != hashAlgorithm {
			continue
		}

		expectedHash = &seekHash

		break
	}

	if expectedHash == nil {
		result.Error = errors.New("Hash not found")

		return result
	}

	result.Expected = expectedHash.Hash

	reader, err := actual.Reader()
	if err != nil {
		result.Error = bosherr.WrapError(err, "Reading file")

		return result
	}

	algorithm, err := crypto.GetAlgorithm(expectedHash.Type)
	if err != nil {
		result.Error = bosherr.WrapError(err, "Getting algorithm")

		return result
	}

	actualHash, err := algorithm.CreateDigest(reader)
	if err != nil {
		result.Error = bosherr.WrapError(err, "Creating digest")

		return result
	}

	result.Expected = crypto.GetDigestHash(actualHash)

	if expectedHash.Hash == crypto.GetDigestHash(actualHash) {
		return result
	}

	result.Error = errors.New("Hash mismatch")

	return result
}

func (v Verifier) VerifyHashes(expected metalink.File, actual origin.Origin) (MultipleVerifyResult, error) {
	var result = MultipleVerifyResult{}

	if len(expected.Hashes) == 0 {
		return result, errors.New("No hashes to verify")
	}

	for _, verifyHashType := range expected.Hashes {
		result = append(result, v.VerifyHash(expected, actual, verifyHashType.Type))
	}

	return result, nil
}

func (v Verifier) VerifySignature(expected metalink.File, actual origin.Origin, trustStore io.Reader) (VerifyResult, error) {
	if expected.Signature == nil {
		return VerifyResult{}, errors.New("No signature to verify")
	}

	var result = VerifyResult{
		Type:     "sign",
		Expected: "",
		Actual:   "UNKNOWN",
	}

	var verifier signature.Signature

	if expected.Signature.MediaType == "application/pgp-signature" {
		verifier = signature.PGPSignature{}
	} else {
		result.Error = errors.New("Unknown signature type")

		return result, nil
	}

	reader, err := actual.Reader()
	if err != nil {
		result.Error = bosherr.WrapError(err, "Reading file")

		return result, nil
	}

	if trustStore == nil {
		trustStore, err = verifier.DefaultTrustStore()
		if err != nil {
			result.Error = bosherr.WrapError(err, "Loading default trust store")

			return result, nil
		}
	}

	confirmation, err := verifier.Verify(trustStore, reader, strings.NewReader(expected.Signature.Signature))
	if err != nil {
		result.Error = err
	} else {
		result.Actual = confirmation
	}

	return result, nil
}
