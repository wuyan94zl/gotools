package newcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
)

var genMainTpl = `package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"{{.packageSrc}}/config"
	"{{.packageSrc}}/container"
	"{{.packageSrc}}/router"
	"github.com/wuyan94zl/gotools/core/logz"
	"github.com/wuyan94zl/gotools/core/utils"
)

func main() {
	c := config.GlobalConfig
	utils.MustConfig("/config.yaml", c)
	logz.InitLog(c.Log)

	container.NewContainer(c)

	app := gin.New()
	app.Use(logz.GinLogger(), logz.GinRecovery(true))
	group := app.Group("")
	router.RegisterHandlers(group)
	app.Run(fmt.Sprintf("%s:%d", c.Host, c.Port))
}

`

func genMain(c *Command) error {
	return utils.GenFile(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     "main.go",
		TemplateFile: genMainTpl,
		Data: map[string]string{
			"packageSrc": c.packageSrc,
		},
	})
}
