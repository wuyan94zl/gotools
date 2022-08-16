package queue

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

var queue *Instance
var Mux *asynq.ServeMux

func NewMux() *asynq.ServeMux {
	if Mux != nil {
		Mux = asynq.NewServeMux()
	}
	return Mux
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
	asy.Run(Mux)
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
