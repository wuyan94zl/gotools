# 一个简单的代码生成工具

[![GoDoc](https://godoc.org/github.com/hibiken/asynq?status.svg)](https://godoc.org/github.com/hibiken/asynq)
[![Go Report Card](https://goreportcard.com/badge/github.com/hibiken/asynq)](https://goreportcard.com/report/github.com/hibiken/asynq)
![Build Status](https://github.com/hibiken/asynq/workflows/build/badge.svg)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Gitter chat](https://badges.gitter.im/go-asynq/gitter.svg)](https://gitter.im/go-asynq/community)

## crontab 定时任务代码生成
> 一键生成定时任务功能包

### 命令
`gotools crontab --name mycron1`  
`gotools crontab --name mycron2`  
执行上面两个命令，就能生成以下文件目录代码

### 包结构
```sh
crontab         # 主文件夹             
    mycron1      # 定时任务功能代码包1
    mycron2      # 定时任务功能代码包2
    crontab.go  # 定时任务入口
```

### 使用
1、在main函数中使用
```go
func main(){
    go crontab.NewInstance().Start()
}
```
2、在go-zero中使用
```go
group := service.NewServiceGroup()
group.Add(crontab.NewInstance())
group.Start()
```

## queue 队列代码生成
> 一键生成队列代码功能包

### 命令
`gotools queue --name myqueue1`  
`gotools queue --name myqueue2`  
执行上面两个命令，就能生成以下文件目录代码

### 包结构
```sh
queue         # 主文件夹             
    myqueue1      # 定时任务功能代码包1
    myqueue2      # 定时任务功能代码包2
    queue.go  # 定时任务入口
```

### 使用
1、在main函数中使用
```go
func main(){
    go queue.NewInstance("127.0.0.1", 6379, "123456").Start()
}
```
2、在go-zero中使用
```go
group := service.NewServiceGroup()
group.Add(queue.NewInstance("127.0.0.1", 6379, "123456"))
group.Start()
```
> "127.0.0.1"：redis host地址, 6379：端口, "123456"：密码

## sql gorm model 生成
> 根据sql一键生成gorm model代码包

### 命令
`gotools gorm --src user.sql --dir user --cache true`

