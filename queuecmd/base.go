package queuecmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"io/ioutil"
)

var tpl = `package {{.package}}

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	//"github.com/wuyan94zl/gotools/queue/myqueue"
	{{.import}}
)

func init() {
	newMux()
	//mux.HandleFunc(myqueue.QueueKey, myqueue.Handle)
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

func NewInstance(addr string, port int, pwd string) *Instance {
	queue = &Instance{
		RedisAddr: addr,
		RedisPort: port,
		RedisPwd:  pwd,
	}
	return queue
}

type Instance struct {
	RedisAddr string
	RedisPort int
	RedisPwd  string
}

func (q *Instance) Start() {
	q.run()
}

func (q *Instance) Stop() {
}

func (q *Instance) run() {
	asy := asynq.NewServer(
		asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", q.RedisAddr, q.RedisPort), Password: q.RedisPwd},
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

func Add(queueKey string, params interface{}, option ...asynq.Option) {
	task, err := addTask(queueKey, params)
	if err != nil {
		return
	}

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", queue.RedisAddr, queue.RedisPort), Password: queue.RedisPwd})
	defer client.Close()
	client.Enqueue(task, option...)
}

func addTask(queueKey string, params interface{}) (*asynq.Task, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(queueKey, payload), nil
}

`

func genBase(c *Command) error {
	packageStr, err := utils.GetPackage()
	if err != nil {
		return err
	}
	dir, _ := ioutil.ReadDir(c.wd)
	importStr := ""
	initStr := ""
	for _, v := range dir {
		if v.IsDir() == true {
			importStr = fmt.Sprintf("%s\n\"%s/queue/%s\"", importStr, packageStr, v.Name())
			initStr = fmt.Sprintf("%s\nmux.HandleFunc(%s.QueueKey, %s.Handle)", initStr, v.Name(), v.Name())
		}
	}
	return utils.GenFile(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     c.packageName + ".go",
		TemplateFile: tpl,
		Data: map[string]string{
			"package": c.packageName,
			"import":  importStr,
			"init":    initStr,
		},
	})
}
