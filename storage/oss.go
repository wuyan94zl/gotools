package storage

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"path/filepath"
)

type ossInstance struct {
	base
	cli *oss.Bucket
	uri string
	//maxFileSize int64
}

func (c *ossInstance) Put(r *http.Request, fileField, toFilePath string, fileName ...string) (string, error) {
	name, file, _, err := c.putParams(r, fileField, toFilePath, fileName...)
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("%s/%s", toFilePath, name)
	err = c.cli.PutObject(uri, file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", c.uri, uri), nil
}

func (c *ossInstance) PutFromFile(filePath, toFilePath string, fileName ...string) (string, error) {
	// 文件名处理
	_, name := filepath.Split(filePath)
	if len(fileName) == 0 {
		name = genFileName(filePath)
	}
	uri := fmt.Sprintf("%s/%s", toFilePath, name)
	err := c.cli.PutObjectFromFile(uri, filePath)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s/%s", c.uri, uri), nil
}
func (c *ossInstance) Delete(path string) error {
	return c.cli.DeleteObject(path)
}

var ossCli *ossInstance

func newOssInstance(conf *config) *ossInstance {
	if ossCli != nil {
		return ossCli
	}
	cli, err := oss.New(conf.OSS.Endpoint, conf.OSS.AccessId, conf.OSS.AccessSecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	bucket, err := cli.Bucket(conf.OSS.BucketName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ossCli = &ossInstance{cli: bucket, uri: fmt.Sprintf("https://%s.%s", conf.OSS.BucketName, conf.OSS.Endpoint), base: base{maxFileSize: conf.maxFileSize}}
	return ossCli
}
