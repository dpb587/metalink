# metalink-repository

For enumerating blob receipts from an external repository.


## Installation

    $ go get github.com/dpb587/metalink/repository
    $ cd $GOPATH/src/github.com/dpb587/metalink/repository
    $ go build -o metalink-repository cli/blob_receipt_repository.go


## Usage

**List** all receipts...

    $ metalink-repository list git+ssh://git@github.com/dpb587/upstream-metalinks.git//wordpress.org/wordpress
    wordpress-4.7.3.tar.gz.json
    wordpress-4.7.2.tar.gz.json
    wordpress-4.7.1.tar.gz.json
    wordpress-4.7.tar.gz.json
    wordpress-4.6.4.tar.gz.json

Or further configure the listing with...

 * filter(s) `--filter=v:version:^4.7`
 * sorting `--sort=v:version:asc`
 * limit results `--limit=1`
 * raw receipt JSON `--raw`
