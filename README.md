# blob-receipt

For receipt files that contain digests, metadata, and source information about blobs.


## Installation

    $ go get github.com/dpb587/blob-receipt


## Usage

**Create** a receipt from an existing blob...

    $ blob-receipt create path/to/receipt.json path/to/blob.tgz

Or further configure the receipt with...

 * extra metadata using `--metadata=stringkey=stringvalue`
 * download locations using `--origin=https://s3.amazonaws.com/bucket/blob.tgz`
 * another name using `--name=otherblob.tgz`
 * another timestamp using `--time={ISO8601-format}`
 * subset of digests using `--digest=md5,sha1,sha256,sha512`

Or update existing receipts by leaving off the blob argument. Update the existing receipt file with `--overwrite`.

**Verify** a blob...

    $ blob-receipt verify path/to/receipt.json path/to/blob.tgz

Or further configure the receipt verification with...

 * a specific digest `--digest=sha1`

**Download** (and verify) a blob...

    $ blob-receipt download path/to/receipt.json
    $ stat blob.tgz

Or further configure the blob downloading with...

 * a specific source URI using `--source=https://...`


## License

[MIT License](LICENSE)
