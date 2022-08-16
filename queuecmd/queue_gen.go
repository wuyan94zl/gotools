package queuecmd

import (
	"github.com/wuyan94zl/gotools/utils"
)

var genTpl = `package {{.package}}

import (
	"context"
	"encoding/json"
	
	"github.com/hibiken/asynq"
	"github.com/wuyan94zl/gotools/queue"
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

func init() {
	queue.NewMux().HandleFunc(QueueKey, Handle)
}

`

func genQueueGen(data *Command) error {
	return utils.GenFile(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     data.packageName + "_gen.go",
		TemplateFile: genTpl,
		Data: map[string]string{
			"package": data.packageName,
		},
	})
}
