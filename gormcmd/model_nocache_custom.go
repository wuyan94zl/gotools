package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
	"strings"
)

var noCacheCustomModelTpl = `package {{.package}}

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"{{.projectPkg}}/{{.modelPkg}}"
)

type (
	{{.StructName}}Model interface {
		{{.model}}.{{.StructName}}Model
	}
	custom{{.StructName}}Model struct {
		*{{.model}}.Default{{.StructName}}Model
	}
)

func New{{.StructName}}Model(gormDb *gorm.DB, cache *redis.Client) {{.StructName}}Model {
	return &custom{{.StructName}}Model{
		Default{{.StructName}}Model: {{.model}}.New{{.StructName}}Model(gormDb, cache),
	}
}

`

func setGormNoCacheCustomModel(data *Command) error {
	wd := filepath.Join(data.wd, VarStringDir)
	modelPkg := filepath.Base(data.wd) + "/_gen"
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "custom.go",
		TemplateFile: noCacheCustomModelTpl,
		Data: map[string]string{
			"package":    data.packageName,
			"projectPkg": data.projectPkg,
			"StructName": data.structName,
			"structName": strings.ToLower(data.structName[:1]) + data.structName[1:],
			"modelPkg":   modelPkg,
			"model":      "_gen",
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
