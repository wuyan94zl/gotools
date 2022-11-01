package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"strings"
)

var customModelTpl = `package {{.package}}

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type (
	{{.StructName}}Model interface {
		{{.structName}}Model
	}
	custom{{.StructName}}Model struct {
		*default{{.StructName}}Model
	}
)

func New{{.StructName}}Model(gormDb *gorm.DB, cache *redis.Client) {{.StructName}}Model {
	return &custom{{.StructName}}Model{
		default{{.StructName}}Model: new{{.StructName}}Model(gormDb,cache),
	}
}

`

func setGormCustomModel(data *Command) error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     "model.go",
		TemplateFile: customModelTpl,
		Data: map[string]string{
			"package":    data.packageName,
			"StructName": data.structName,
			"structName": strings.ToLower(data.structName[:1]) + data.structName[1:],
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
