package blobreceipt

type blobReceiptMetadataByKey []BlobReceiptMetadata

func (s blobReceiptMetadataByKey) Len() int {
	return len(s)
}
func (s blobReceiptMetadataByKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s blobReceiptMetadataByKey) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

type blobReceiptOriginByURI []BlobReceiptOrigin

func (s blobReceiptOriginByURI) Len() int {
	return len(s)
}
func (s blobReceiptOriginByURI) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s blobReceiptOriginByURI) Less(i, j int) bool {
	return s[i].URI() < s[j].URI()
}
