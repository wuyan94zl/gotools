package newcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
)

var genMainTpl = `package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wuyan94zl/gotools/core/logz"
	"github.com/wuyan94zl/gotools/core/utils"
	
	"{{.packageSrc}}/config"
	"{{.packageSrc}}/container"
	"{{.packageSrc}}/router"
)

// @title 代码生成工具 gotools
// @version 1.0
// @description 代码生成工具

// @securityDefinitions.apikey JwtAuth
// @in header
// @name Authorization

// @host localhost:8888
// @BasePath /v1
func main() {
	c := config.GlobalConfig
	utils.MustConfig("/config.yaml", c)
	logz.InitLog(c.Log)

	container.NewContainer(c)

	app := gin.New()
	app.Use(logz.GinLogger(), logz.GinRecovery(true))

	// exec swag init && import _ "{{.packageSrc}}/docs"
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
