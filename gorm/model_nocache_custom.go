package gorm

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"strings"
)

var noCacheCustomModelTpl = `package {{.package}}

import "gorm.io/gorm"

type (
	{{.StructName}}Model interface {
		{{.structName}}Model
	}
	custom{{.StructName}}Model struct {
		*default{{.StructName}}Model
	}
)

func New{{.StructName}}Model(gormDb *gorm.DB) {{.StructName}}Model {
	return &custom{{.StructName}}Model{
		default{{.StructName}}Model: new{{.StructName}}Model(gormDb),
	}
}

`

func setGormNoCacheCustomModel(data *Command) error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     "model.go",
		TemplateFile: noCacheCustomModelTpl,
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
