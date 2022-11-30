# 基于 gin（http） gorm（数据库） redis（缓存） zap（日志） 等完成基础web-api项目基准框架

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
> --package 或 -p 表示项目包名

## 初始目录结构
```
common
    |-- errcode
        |-- code.go     # 全局错误码定义
config
    |-- config.go       # 配置信息结构体
container
    |-- conn
        |-- gorm.go     # gorm mysql连接实例
        |-- redis.go    # redis 连接实例
    |-- kernal.go
    |-- container.go    # 全局容器
router
    |-- route.go        # 路由
main.go
config.yaml             # 配置文件和config结构体对应
```

## api 生成
> 生成api

### 命令
`gotools api --dir v1/user --name info --params :id --method GET`
> --dir 表示api的文件目录 v1/user  
> --name 表示api接口名称 info  
> --params 表示路由参数 如：`:id`  
> 以上参数生成路由为：v1/user/info/:id，逻辑实现代码目录为：app/logic/v1/user/info.go  
> --method 表示api请求方式 GET POST DELETE PUT RESTFUL  


## crontab 定时任务代码生成
> 一键生成定时任务功能包

### 命令
`gotools crontab --name mycron`

### 目录结构
```
crontab
    |-- mycron
        |-- cronjob.go
    |-- crontab.go             
```

###  使用
1、在main函数中使用
```go
func main(){
	...
    go crontab.NewInstance().Start()
	...
}
```
### 实现
编辑 `crontab/mycron/cronjob.go` 代码即可


## queue 队列代码生成
> 一键生成队列代码功能包

###  命令
`gotools queue --name myqueue`

###  目录结构
```
queue
    |-- myqueue
        |-- myqueue.go
    |-- queue.go    
```

###  使用
在main函数中使用
```go
func main(){
    ...
	go queue.NewInstance("127.0.0.1:6379", "123456").Start()
    ...
}
```
### 实现消费
编辑 `queue/myqueue/myqueue.go` 代码

### 发布消息
```go
// myqueue.QueueKey 和 myqueue.Params{} 为上面设置的数据
queue.Add(myqueue.QueueKey, myqueue.Params{})
```

## sql to gorm model 生成
> 根据sql一键生成gorm model代码包

### 根目录下创建models目录
> 以下所有命令在models目录下执行

### 生成model
1、创建user.sql  
2、执行：`gotools gorm --src user.sql --dir user --cache true`

####  目录结构
```
models
    |-- user
        |-- custom.go
    |-- user_gen.go
    |-- user.sql
```

### 使用
注入
`container/container.go`
```go
type Container struct{
	...
    UserModel      user.UsersModel // 增加代码
	...
}

// NewContainer 增加UserModel 实例
UserModel:      user.NewUsersModel(gormConn, redisConn), // 增加代码
```
使用
```go
container.Instance().UserModel.Insert(ctx,&models.Users{})
container.Instance().UserModel.First(ctx,id)
container.Instance().UserModel.Update(ctx,&models.Users{})
container.Instance().UserModel.Delete(ctx,&models.Users{})
```
