# 一个简单的代码生成工具

[![GoDoc](https://godoc.org/github.com/hibiken/asynq?status.svg)](https://godoc.org/github.com/hibiken/asynq)
[![Go Report Card](https://goreportcard.com/badge/github.com/hibiken/asynq)](https://goreportcard.com/report/github.com/hibiken/asynq)
![Build Status](https://github.com/hibiken/asynq/workflows/build/badge.svg)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Gitter chat](https://badges.gitter.im/go-asynq/gitter.svg)](https://gitter.im/go-asynq/community)


## 安装
`go install github.com/wuyan94zl/gotools`

[完整使用示例](https://learnku.com/articles/70767)

## 项目初始化
`gotools init --package wuyan94zl/project`

### 目录结构
```
config
    |-- config.go
container
    |-- conn.go
    |-- kernal.go
    |-- container.go
router
    |-- route.go
main.go
config.yaml
```

## crontab 定时任务代码生成
> 一键生成定时任务功能包

#### 命令
`gotools crontab --name mycron`
执行上面两个命令，就能生成以下文件目录代码

#### 目录结构
```
crontab
    |-- mycron
        |-- cronjob.go
    |-- crontab.go             
```

####  使用
1、在main函数中使用
```go
func main(){
    go crontab.NewInstance().Start() 
    // gin、beego等 启动http server
}
```
2、在go-zero中使用
```go
group := service.NewServiceGroup() // go-zero service group
group.Add(crontab.NewInstance()) // 添加 crontab service 到 group
group.Start() // go-zero 启动 service group
```
####  生成初始代码
```go
package test

import (...)

const Spec = "0 * * * * *" // todo 设置定时时间 秒 分 时 日 月 周

func NewJob() *Job {
	return &Job{}
}

type Job struct{}

func (j *Job) Run() {
	// todo 定时处理逻辑
	fmt.Println("Execution per minute", time.Now().Format("2006-01-02 15:4:05"))
}
```

## queue 队列代码生成
> 一键生成队列代码功能包

####  命令
`gotools queue --name myqueue`   
执行上面两个命令，就能生成以下文件目录代码

####  目录结构
```
queue
    |-- myqueue
        |-- myqueue.go
    |-- queue.go    
```

####  使用
1、在main函数中使用
```go
func main(){
    go queue.NewInstance("127.0.0.1:6379", "123456").Start()
    // gin、beego等 启动http server
}
```
2、在go-zero中使用
```go
group := service.NewServiceGroup() // go-zero service group
group.Add(queue.NewInstance("127.0.0.1:6379", "123456")) // 添加 crontab service 到 group
group.Start() // go-zero 启动 service group
```
####  生成初始代码
```go
package test

import (...)

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
	// todo 队列消息消费业务逻辑处理
}
```
####  代码中发布消息
```go
// test.QueueKey 和 test.Params{} 为上面设置的数据
// queue 队列入口包
queue.Add(test.QueueKey, test.Params{})
```

## sql gorm model 生成
> 根据sql一键生成gorm model代码包

####  命令
`gotools gorm --src user.sql --dir user --cache true`

####  目录结构
```
models
    |-- user
        |-- model.go
        |-- model_gen.go
    |-- user.sql
```
### 使用
```go
// cache model
uModel := user.NewChatUsersModel(gormDB,redisCli)

// no cache model
// uModel := user.NewChatUsersModel(gormDB)

uModel.Insert(ctx,user.Users{})
uModel.First(ctx,id)
uModel.Update(ctx,user.Users{})
uModel.Delete(ctx,id)

```
