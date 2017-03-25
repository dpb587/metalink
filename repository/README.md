# metalink-repository

For enumerating metalinks from an external repository.


## Assumptions

 * metalinks are named with the convention of `v{semver}.meta4`
    * metalinks may contain multiple files, but file nodes must have the same `version`


## Usage

**List** all receipts...

    $ metalink-repository list git+ssh://git@github.com/dpb587/upstream-metalinks.git//repository/wordpress.org/wordpress
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
