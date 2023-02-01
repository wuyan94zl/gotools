package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
	"strings"
)

var customModelTpl = `package {{.package}}

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"{{.projectPkg}}/{{.modelPkg}}"
)

type (
	{{.StructName}}Model interface {
		{{.modelPkg}}.{{.StructName}}Model
	}
	custom{{.StructName}}Model struct {
		*{{.modelPkg}}.Default{{.StructName}}Model
	}
)

func New{{.StructName}}Model(gormDb *gorm.DB, cache *redis.Client) {{.StructName}}Model {
	return &custom{{.StructName}}Model{
		Default{{.StructName}}Model: {{.modelPkg}}.New{{.StructName}}Model(gormDb,cache),
	}
}

`

func setGormCustomModel(data *Command) error {
	wd := filepath.Join(data.wd, VarStringDir)
	modelPkg := filepath.Base(data.wd)
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "custom.go",
		TemplateFile: customModelTpl,
		Data: map[string]string{
			"package":    data.packageName,
			"projectPkg": data.projectPkg,
			"StructName": data.structName,
			"structName": strings.ToLower(data.structName[:1]) + data.structName[1:],
			"modelPkg":   modelPkg,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
