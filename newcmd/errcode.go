package newcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
)

var genErrCodeTpl = `package errcode

// 错误码格式为：数字code码 + 错误信息，以英文逗号隔开。
// 业务中使用 response.NewErrorCode(errcode.SystemError)

const (
	SystemError = "100000,系统错误"
)

`

func genErrCode(c *Command) error {
	wd := filepath.Join(c.wd, "common", "errcode")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "code.go",
		TemplateFile: genErrCodeTpl,
		Data:         map[string]string{},
	})
}
