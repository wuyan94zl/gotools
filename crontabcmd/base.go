package crontabcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"io/ioutil"
)

var baseTpl = `package crontab

import (
	"github.com/robfig/cron/v3"
	{{.import}}
)

var Cron *cron.Cron

func init() {
	newCron()
	{{.init}}
}


func newCron() *cron.Cron {
	if Cron == nil {
		Cron = cron.New(cron.WithSeconds())
	}
	return Cron
}

type Instance struct {
}

func (q *Instance) Start() {
	q.run()
}

func (q *Instance) Stop() {
}

func (q *Instance) run() {
	Cron.Start()
}
func NewInstance() *Instance {
	return &Instance{}
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
			importStr = fmt.Sprintf("%s\n\"%s/crontab/%s\"", importStr, packageStr, v.Name())
			initStr = fmt.Sprintf("%s\nCron.AddJob(%s.Spec, %s.NewJob())", initStr, v.Name(), v.Name())
		}
	}
	return utils.GenFileCover(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     "crontab.go",
		TemplateFile: baseTpl,
		Data: map[string]string{
			"package": c.packageName,
			"import":  importStr,
			"init":    initStr,
		},
	})
}
