package blobreceipt

type BlobReceiptOrigin map[string]interface{}

func (bro BlobReceiptOrigin) URI() string {
	uri, found := bro["uri"]
	if !found {
		panic("expected uri field in origin")
	}

	stringUri, ok := uri.(string)
	if !ok {
		panic("expected string in origin uri")
	}

	return stringUri
}
