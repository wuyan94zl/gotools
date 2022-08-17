package storage

import "net/http"

type Client interface {
	Put(r *http.Request, fileField, toFilePath string, fileName ...string) (string, error)
	PutFromFile(filePath, toFilePath string, fileName ...string) (string, error)
	Delete(path string) error
}

var cli Client

func NewInstance(c *config) {
	switch c.Store {
	case "cos":
		cli = newCosInstance(c)
	case "oss":
		cli = newOssInstance(c)
	case "qn":
		cli = newQiniuInstance(c)
	case "local":
		cli = newLocalInstance(c)
	}
}

func Parse(store string, maxFileSize int64, conf interface{}) {
	c := &config{
		Store:       store,
		maxFileSize: maxFileSize,
	}
	switch store {
	case "local":
		c.Local = conf.(LocalConfig)
	case "oss":
		c.OSS = conf.(OSSConfig)
	case "cos":
		c.COS = conf.(COSConfig)
	case "qn":
		c.QN = conf.(QNConfig)
	}
	NewInstance(c)
}

func Put(r *http.Request, fileField, toFilePath string, fileName ...string) (string, error) {
	return cli.Put(r, fileField, toFilePath, fileName...)
}
func PutFromFile(filePath, toFilePath string, fileName ...string) (string, error) {
	return cli.PutFromFile(filePath, toFilePath, fileName...)
}
func Delete(path string) error {
	return cli.Delete(path)
}
