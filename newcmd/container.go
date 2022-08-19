package newcmd

import (
	"github.com/wuyan94zl/gotools/utils"
	"path/filepath"
)

func genContainer(c *Command) error {
	err := genContainerConn(c)
	if err != nil {
		return err
	}
	err = genContainerKernel(c)
	if err != nil {
		return err
	}
	return genContainerM(c)
}

var genContainerConnTpl = `package container

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormConfig struct {
	DataSource string
}

type RedisConfig struct {
	Host string
	Pass string
}

func dbConn(c GormConfig) *gorm.DB {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      c.DataSource,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("mysql 链接错误")
	}
	return gormDB
}

func redisConn(c RedisConfig) *redis.Client {
	redisConn := redis.NewClient(&redis.Options{Addr: c.Host, Password: c.Pass})
	return redisConn
}

`

func genContainerConn(c *Command) error {
	wd := filepath.Join(c.wd, "container")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "conn.go",
		TemplateFile: genContainerConnTpl,
		Data:         map[string]string{},
	})
}

var genContainerKernelTpl = `package container

var container *Container

func Instance() *Container {
	return container
}
`

func genContainerKernel(c *Command) error {
	wd := filepath.Join(c.wd, "container")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "kernel.go",
		TemplateFile: genContainerKernelTpl,
		Data:         map[string]string{},
	})
}

var genContainerTpl = `package container

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	DB        GormConfig
	Redis     RedisConfig
}

type Container struct {
	DB        *gorm.DB
	Redis     *redis.Client
}

func NewContainer(c Config) {
	dbConn, redisConn := dbConn(c.DB), redisConn(c.Redis)
	container = &Container{
		DB:        dbConn,
		Redis:     redisConn,
	}
}

`

func genContainerM(c *Command) error {
	wd := filepath.Join(c.wd, "container")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "kernel.go",
		TemplateFile: genContainerTpl,
		Data:         map[string]string{},
	})
}
