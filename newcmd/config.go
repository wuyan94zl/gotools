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
	"github.com/wuyan94zl/gotools/core/jwt"
	"github.com/wuyan94zl/gotools/core/logz"
	"{{.packageSrc}}/container/conn"
)

var GlobalConfig = new(Config)

type Config struct {
	Name   string
	Host   string
	Port   int
	Jwt    jwt.Config
	DB     conn.GormConfig
	Redis  conn.RedisConfig
	Log    logz.Config
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

# 日志配置
Log:
  # Default：日志方式。 console:控制台日志打印，file：文件日志存储，kafka：写入kafka队列，需要配置elk相关组件
  Default: "console"
  Level: "info"    #  debug info error 等
  Encoder: "plain" # plain json

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
