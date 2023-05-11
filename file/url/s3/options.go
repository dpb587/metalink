package s3

type Options struct {
	AccessKey string `json:"access_key" yaml:"access_key"`
	SecretKey string `json:"secret_key" yaml:"secret_key"`
	RoleARN   string `json:"role_arn"   yaml:"role_arn"`
}
