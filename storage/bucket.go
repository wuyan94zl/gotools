package storage

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

type base struct {
	maxFileSize int64
}

func (b *base) putParams(r *http.Request, fileField, toFilePath string, fileName ...string) (string, multipart.File, *multipart.FileHeader, error) {
	err := r.ParseMultipartForm(b.maxFileSize)
	if err != nil {
		return "", nil, nil, err
	}
	file, handle, err := r.FormFile(fileField)
	defer file.Close()
	if err != nil {
		return "", nil, nil, err
	}
	// 文件名处理
	name := handle.Filename
	if len(fileName) == 0 {
		name = genFileName(handle.Filename)
	}
	return name, file, handle, nil
}

func genFileName(name string) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Int63n(time.Now().UnixNano())
	fileName := Md5ByString(fmt.Sprintf("%s%d", name, num))
	return fmt.Sprintf("%s%s", fileName, filepath.Ext(name))
}

func Md5ByString(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
