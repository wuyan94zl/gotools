package storage

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type LocalInstance struct {
	base
	dir string
	url string
}

func (c *LocalInstance) Put(r *http.Request, fileField, toFilePath string, fileName ...string) (string, error) {
	name, file, _, err := c.putParams(r, fileField, toFilePath, fileName...)
	if err != nil {
		return "", err
	}
	tempFile, err := c.createFile(path.Join(c.dir, toFilePath), name)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	io.Copy(tempFile, file)

	return path.Join(c.url, c.dir, toFilePath, name), nil
}

func (c *LocalInstance) PutFromFile(filePath, toFilePath string, fileName ...string) (string, error) {
	return "该存储方式不支持", nil
}

func (c *LocalInstance) createFile(dir, filename string) (fp *os.File, err error) {
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0777)
	}
	filePath := filepath.Join(dir, filename)
	fp, err = os.Create(filePath)
	return
}
func (c *LocalInstance) Delete(path string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	return os.Remove(filepath.Join(wd, c.dir, path))
}

var localCli *LocalInstance

func newLocalInstance(conf *config) *LocalInstance {
	if localCli != nil {
		return localCli
	}
	localCli = &LocalInstance{dir: conf.Local.BaseDir, url: conf.Local.BaseUrl, base: base{maxFileSize: conf.maxFileSize}}
	return localCli
}
