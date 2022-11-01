package queuecmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
)

var genTpl = `package {{.package}}

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
)

func Handle(ctx context.Context, t *asynq.Task) error {
	params := Params{}
	err := json.Unmarshal(t.Payload(), &params)
	if err != nil {
		return err
	}
	Do(ctx, params)
	return nil
}

const QueueKey = "key" // todo 自定义队列key

type Params struct {
	// todo 自定义队列参数结构体
}

func Do(ctx context.Context, params Params) {
	// todo 队列业务逻辑处理
}


`

func genJob(data *Command) error {
	return utils.GenFile(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     data.packageName + ".go",
		TemplateFile: genTpl,
		Data: map[string]string{
			"package": data.packageName,
		},
	})
}
