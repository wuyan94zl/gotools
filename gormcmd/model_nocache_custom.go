package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
	"strings"
)

var noCacheCustomModelTpl = `package {{.package}}

import (
	"{{.projectPkg}}/container"
	"{{.projectPkg}}/{{.basePkg}}"
)

type (
	I{{.StructName}} interface {
		base.IBase
	}
	custom{{.StructName}}Model struct {
		*base.Base
	}
)

func New{{.StructName}}Model() I{{.StructName}} {
	return &custom{{.StructName}}Model{
		Base: base.NewBase({{.StructName}}Model{}, container.Instance().DB),
	}
}

`

func setGormNoCacheCustomModel(data *Command) error {
	wd := filepath.Join(data.wd, VarStringDir)
	basePkg := filepath.Base(data.wd) + "/base"
	tablePkg := filepath.Base(data.wd) + "/tables"
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "custom.go",
		TemplateFile: noCacheCustomModelTpl,
		Data: map[string]string{
			"package":    data.packageName,
			"projectPkg": data.projectPkg,
			"StructName": data.structName,
			"structName": strings.ToLower(data.structName[:1]) + data.structName[1:],
			"basePkg":    basePkg,
			"tablePkg":   tablePkg,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
