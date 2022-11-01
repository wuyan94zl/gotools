package storage

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
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
	if err != nil {
		return "", nil, nil, err
	}
	defer file.Close()
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
	fileName := utils.Md5ByString(fmt.Sprintf("%s%d", name, num))
	return fmt.Sprintf("%s%s", fileName, filepath.Ext(name))
}
