# metalink

For receipt files that contain digests, metadata, and source information about blobs.


## Installation

    $ go get github.com/dpb587/metalink
    $ cd $GOPATH/src/github.com/dpb587/metalink
    $ go build -o metalink cli/blob_receipt.go


## Usage

**Create** a receipt from an existing blob...

    $ metalink create path/to/receipt.json path/to/blob.tgz

Or further configure the receipt with...

 * extra metadata using `--metadata=stringkey=stringvalue`
 * download origins using `--origin=https://s3.amazonaws.com/bucket/blob.tgz`
 * another name using `--name=otherblob.tgz`
 * another timestamp using `--time={ISO8601-format}`
 * subset of digests using `--digest=md5,sha1,sha256,sha512`

Or update existing receipts by leaving off the blob argument. Update the existing receipt file with `--overwrite`.

**Verify** a blob with the strongest digest of its receipt...

    $ metalink verify path/to/receipt.json path/to/blob.tgz

Or further configure the receipt verification with...

 * a specific digest `--digest=sha1`
 * all recognized digests `--digest=all`
 * quieter output `--quiet`

**Download** (and verify) a blob...

    $ metalink download path/to/receipt.json
    $ stat blob.tgz

Or further configure the blob downloading with...

 * a specific source URI using `--origin=https://...`

**Upload** a blob to another origin and add to receipt...

    $ metalink upload path/to/receipt.json s3://endpoint/bucket/new/origin/blob.tgz path/to/blob.tgz

Or further configure the blob uploading with...

 * ignoring the origin in the receipt using `--skip-receipt-update`


## Notes

 * review the [schema](schema.json) for details on the data format


## License

[MIT License](LICENSE)
