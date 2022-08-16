package queuecmd

import "github.com/wuyan94zl/gotools/utils"

var tpl = `package {{.package}}

import (
	"context"
)

const QueueKey = "key" // todo 自定义队列key

type Params struct {
	// todo 自定义队列参数结构体
}

func Do(ctx context.Context, params Params) {
	// todo 队列业务逻辑处理
}

`

func genQueue(data *Command) error {
	return utils.GenFile(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     data.packageName + ".go",
		TemplateFile: tpl,
		Data: map[string]string{
			"package": data.packageName,
		},
	})
}
