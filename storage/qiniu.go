package storage

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/http"
	"path/filepath"
	"time"
)

type qnInstance struct {
	base
	cli           *storage.FormUploader
	uri           string
	token         string
	bucketName    string
	bucketManager *storage.BucketManager
}

func (q *qnInstance) Put(r *http.Request, fileField, toFilePath string, fileName ...string) (string, error) {
	name, file, handle, err := q.putParams(r, fileField, toFilePath, fileName...)
	if err != nil {
		return "", err
	}
	uri := fmt.Sprintf("%s/%s", toFilePath, name)
	ret := new(storage.PutRet)
	err = q.cli.Put(context.Background(), ret, q.token, uri, file, handle.Size, nil)
	return fmt.Sprintf("%s/%s", q.uri, ret.Key), err
}

func (q *qnInstance) PutFromFile(filePath, toFilePath string, fileName ...string) (string, error) {
	// 文件名处理
	_, name := filepath.Split(filePath)
	if len(fileName) == 0 {
		name = genFileName(filePath)
	}
	uri := fmt.Sprintf("%s/%s", toFilePath, name)
	ret := new(storage.PutRet)
	err := q.cli.PutFile(context.Background(), ret, q.token, uri, filePath, nil)
	return fmt.Sprintf("%s/%s", q.uri, ret.Key), err
}

func (q *qnInstance) Delete(path string) error {
	return q.bucketManager.Delete(q.bucketName, path)
}

var qnCli *qnInstance

func newQiniuInstance(conf *config) *qnInstance {
	if qnCli != nil {
		return qnCli
	}
	credentials := auth.New(conf.QN.AccessKey, conf.QN.SecretKey)
	policy := storage.PutPolicy{Scope: conf.QN.BucketName, Expires: uint64(time.Now().Unix()) + 60}
	token := policy.UploadToken(credentials)
	c := &storage.Config{UseHTTPS: false, UseCdnDomains: false}
	cli := storage.NewFormUploader(c)
	qnCli = &qnInstance{
		cli:           cli,
		uri:           conf.QN.SiteUrl,
		token:         token,
		bucketName:    conf.QN.BucketName,
		bucketManager: storage.NewBucketManager(credentials, c),
		base:          base{maxFileSize: conf.maxFileSize},
	}
	return qnCli
}
