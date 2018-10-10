package s3

type Options struct {
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	Private   bool   `json:"private" yaml:"private"`
}
