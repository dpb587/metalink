// For metalink repositories stored on a S3-compatible server.
//
// **Supported schemes**
//
//   - `s3`
//
// **URI components**
//
// *Host* - for S3, this should be `s3.amazonaws.com`
// *Path* - the first component must be the bucket name, all following names are treated as a path prefix to find the
// repository
// *Username/Password* - the access and secret key (remember to encode the password, if necessary; e.g. `/` to `%2F`)
//
// **Example URIs**
//
// Simple example...
//
//	s3://s3.amazonaws.com/acmecorp-metalink-repository/prod
//
// Custom S3 server...
//
//	s3://minio.example.com:9000/prod-metalink-repository
//
// **Notes**
//
//   - this repository will download all `*.meta4` files within the repository URI whenever it is loaded
//   - insecure S3 endpoints are not supported
//   - environment variables `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` are used when present
package git
