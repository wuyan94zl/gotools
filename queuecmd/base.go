package queuecmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"io/ioutil"
	"path/filepath"
)

var tpl = `package {{.package}}

import (
	"github.com/hibiken/asynq"
	{{.import}}
)

func init() {
	newMux()
	{{.init}}
}

var queue *Instance
var mux *asynq.ServeMux

func newMux() *asynq.ServeMux {
	if mux == nil {
		mux = asynq.NewServeMux()
	}
	return mux
}

func NewInstance(host string, pwd string) *Instance {
	queue = &Instance{
		RedisHost: host,
		RedisPwd:  pwd,
	}
	return queue
}

type Instance struct {
	RedisHost string
	RedisPwd  string
}

func (q *Instance) Start() {
	q.run()
}

func (q *Instance) Stop() {
}

func (q *Instance) run() {
	asy := asynq.NewServer(
		asynq.RedisClientOpt{Addr: q.RedisHost, Password: q.RedisPwd},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	asy.Run(mux)
}
`

func serve(c *Command) error {
	dir, _ := ioutil.ReadDir(c.wd)
	importStr := ""
	initStr := ""
	for _, v := range dir {
		if v.IsDir() == true && v.Name() != "serve" {
			importStr = fmt.Sprintf("%s\n\"%s/queue/%s\"", importStr, c.projectPkg, v.Name())
			initStr = fmt.Sprintf("%s\nmux.HandleFunc(%s.QueueKey, %s.Handle)", initStr, v.Name(), v.Name())
		}
	}
	return utils.GenFileCover(utils.FileGenConfig{
		Dir:          filepath.Join(c.wd, "serve"),
		Filename:     "serve.go",
		TemplateFile: tpl,
		Data: map[string]string{
			"package": "serve",
			"import":  importStr,
			"init":    initStr,
		},
	})
}

var tplQueueClient = `package queue

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	{{.import}}
)

func Add(queueKey string, params interface{}, option ...asynq.Option) (*asynq.TaskInfo, error) {
	task, err := addTask(queueKey, params)
	if err != nil {
		return nil, err
	}
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: config.GlobalConfig.Redis.Host, Password: config.GlobalConfig.Redis.Pass})
	defer client.Close()
	return client.Enqueue(task, option...)
}

func addTask(queueKey string, params interface{}) (*asynq.Task, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(queueKey, payload), nil
}
`

func client(c *Command) error {
	importStr := fmt.Sprintf("\"%s/config\"", c.projectPkg)
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     c.packageName + ".go",
		TemplateFile: tplQueueClient,
		Data: map[string]string{
			"package": c.packageName,
			"import":  importStr,
		},
	})
	return err
}

func genBase(c *Command) error {
	err := client(c)
	if err != nil {
		return err
	}
	return serve(c)
}
