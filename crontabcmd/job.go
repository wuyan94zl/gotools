package crontabcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
)

var jobTpl = `package {{.package}}

import (
	"fmt"
	"time"
)

const Spec = "* * * * * *" // todo 设置定时时间 秒 分 时 日 月 周

func NewJob() *Job {
	return &Job{}
}

type Job struct {}

func (j *Job) Run() {
	// todo 定时处理逻辑
	fmt.Println("crontab exec：", time.Now().Format("2006-01-02 15:4:05"))
}

`

func genJob(c *Command) error {
	return utils.GenFile(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     "cronjob.go",
		TemplateFile: jobTpl,
		Data: map[string]string{
			"package": c.packageName,
		},
	})
}
