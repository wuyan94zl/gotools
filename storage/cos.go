package storage

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"net/http"
	"net/url"
	"path/filepath"
)

type cosInstance struct {
	base
	cli *cos.Client
	uri string
}

func (c *cosInstance) Put(r *http.Request, fileField, toFilePath string, fileName ...string) (string, error) {
	name, file, _, err := c.putParams(r, fileField, toFilePath, fileName...)
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("%s/%s", toFilePath, name)
	_, err = c.cli.Object.Put(context.Background(), uri, file, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", c.uri, uri), nil
}

func (c *cosInstance) PutFromFile(filePath, toFilePath string, fileName ...string) (string, error) {
	// 文件名处理
	_, name := filepath.Split(filePath)
	if len(fileName) == 0 {
		name = genFileName(filePath)
	}
	uri := fmt.Sprintf("%s/%s", toFilePath, name)
	_, err := c.cli.Object.PutFromFile(context.Background(), uri, filePath, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", c.uri, uri), nil
}
func (c *cosInstance) Delete(path string) error {
	_, err := c.cli.Object.Delete(context.Background(), path)
	return err
}

var cosCli *cosInstance

func newCosInstance(conf *config) *cosInstance {
	if cosCli != nil {
		return cosCli
	}
	baseUri := fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com", conf.COS.BucketName, conf.COS.AppId, conf.COS.Region)
	parse, _ := url.Parse(baseUri)
	cli := cos.NewClient(&cos.BaseURL{BucketURL: parse}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.COS.SecretId,
			SecretKey: conf.COS.SecretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  false,
				RequestBody:    false,
				ResponseHeader: false,
				ResponseBody:   false,
			},
		},
	})
	cosCli = &cosInstance{cli: cli, uri: baseUri, base: base{maxFileSize: conf.maxFileSize}}
	return cosCli
}
