package newcmd

import (
	"github.com/wuyan94zl/gotools/utils"
	"path/filepath"
)

func genContainer(c *Command) error {
	err := genGormConn(c)
	if err != nil {
		return err
	}
	err = genRedisConn(c)
	if err != nil {
		return err
	}
	err = genContainerKernel(c)
	if err != nil {
		return err
	}
	return genContainerMain(c)
}

var genGormConnTpl = `package conn

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormConfig struct {
	DataSource string
}

func GormConn(c GormConfig) *gorm.DB {
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
`

func genGormConn(c *Command) error {
	wd := filepath.Join(c.wd, "container", "conn")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "gorm.go",
		TemplateFile: genGormConnTpl,
		Data:         map[string]string{},
	})
}

var genRedisConnTpl = `package conn

import "github.com/go-redis/redis/v9"

var redisConfig RedisConfig

type RedisConfig struct {
	Host string
	Pass string
}

func RedisConn(c RedisConfig) *redis.Client {
	redisConfig = c
	return GetRedisConn()
}

func GetRedisConn() *redis.Client {
	redisConn := redis.NewClient(&redis.Options{Addr: redisConfig.Host, Password: redisConfig.Pass})
	return redisConn
}
`

func genRedisConn(c *Command) error {
	wd := filepath.Join(c.wd, "container", "conn")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "redis.go",
		TemplateFile: genRedisConnTpl,
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
	"github.com/wuyan94zl/gotools/jwt"
	"gorm.io/gorm"

	"{{.packageSrc}}/container/conn"
)

type Config struct {
	DB		conn.GormConfig
	Redis	conn.RedisConfig
	Jwt		jwt.Config
}

type Container struct {
	DB        *gorm.DB
	Redis     *redis.Client
	Jwt       jwt.Config
}

func NewContainer(c Config) {
	gormConn, redisConn := conn.GormConn(c.DB), conn.RedisConn(c.Redis)
	container = &Container{
		DB:        gormConn,
		Redis:     redisConn,
		Jwt:       c.Jwt,
	}
}

`

func genContainerMain(c *Command) error {
	wd := filepath.Join(c.wd, "container")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "container.go",
		TemplateFile: genContainerTpl,
		Data: map[string]string{
			"packageSrc": c.packageSrc,
		},
	})
}
