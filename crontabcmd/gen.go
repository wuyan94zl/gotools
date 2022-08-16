package crontabcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"os"
	"path/filepath"
)

var (
	VarStringName string
	VarStringDir  string
)

var tpl = `package {{.package}}

import (
	"fmt"
	"time"

	"github.com/wuyan94zl/gotools/crontab"
)

func init() {
	crontab.AddJob(spec, newJob())
}

const spec = "0 * * * * *" // todo 设置定时时间 秒 分 时 日 月 周

func newJob() *Job {
	return &Job{}
}

type Job struct {}

func (j *Job) Run() {
	// todo 定时处理逻辑
	fmt.Println("Execution per minute", time.Now().Format("2006-01-02 15:4:05"))
}

`

type Command struct{}

func (c *Command) Run() error {
	wd, _ := os.Getwd()
	if VarStringDir != "." {
		wd = filepath.Join(wd, VarStringDir)
	}
	wd = filepath.Join(wd, VarStringName)
	_, packageName := filepath.Split(wd)

	err := utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "cronjob.go",
		TemplateFile: tpl,
		Data: map[string]string{
			"package": packageName,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
