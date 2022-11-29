package logz

import (
	"fmt"
	"strings"
)

type GormWriter struct {
}

// Printf 实现打印日志接口
func (w *GormWriter) Printf(key string, data ...interface{}) {
	str := strings.Split(fmt.Sprintf(key, data...), "\n")
	Info(str[1])
}

func NewGormWriter() *GormWriter {
	return &GormWriter{}
}
