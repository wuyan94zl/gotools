package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5ByString(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
