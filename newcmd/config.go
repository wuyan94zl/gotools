package newcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
)

func genConfig(c *Command) error {
	err := genConfigGo(c)
	if err != nil {
		return err
	}
	return genConfigYaml(c)
}

var genConfigGoTpl = `package config

import (
	"{{.packageSrc}}/container"
)

type Config struct {
	Name      string
	Host      string
	Port      int
	Container container.Config
}
`

func genConfigGo(c *Command) error {
	wd := filepath.Join(c.wd, "config")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "config.go",
		TemplateFile: genConfigGoTpl,
		Data: map[string]string{
			"packageSrc": c.packageSrc,
		},
	})
}

var genConfigYamlTpl = `Name: "example-api"
Host: "0.0.0.0"
Port: 8888

Container:
  # 数据库
  DB:
    DataSource: "root:123456@tcp(localhost:3306)/blogs?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
  # redis
  Redis:
    Host: "localhost:6379"
    Pass: "123456"
  Jwt:
    Export: 86400
    Secretary: "wuyan94zl"
`

func genConfigYaml(c *Command) error {
	return utils.GenFile(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     "config.yaml",
		TemplateFile: genConfigYamlTpl,
		Data:         map[string]string{},
	})
}
