package storage

type config struct {
	Store       string
	Local       LocalConfig
	OSS         OSSConfig
	COS         COSConfig
	QN          QNConfig
	maxFileSize int64
}

type LocalConfig struct {
	BaseDir string
	BaseUrl string
}

type OSSConfig struct {
	BucketName   string
	Endpoint     string
	AccessId     string
	AccessSecret string
}

type COSConfig struct {
	BucketName string
	AppId      string
	Region     string
	SecretId   string
	SecretKey  string
}

type QNConfig struct {
	AccessKey  string
	SecretKey  string
	SiteUrl    string
	BucketName string
}
