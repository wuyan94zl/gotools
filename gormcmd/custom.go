package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
)

var noCacheCustomModelTpl = `package {{.package}}

import (
	"gorm.io/gorm"

	"{{.nameSpace}}/{{.basePkg}}"
	"{{.nameSpace}}/{{.tablePkg}}"
)

type (
	I{{.StructName}} interface {
		base.IModel
	}
	custom{{.StructName}}Model struct {
		*base.Model
	}
)

func New{{.StructName}}Model(db *gorm.DB) I{{.StructName}} {
	return &custom{{.StructName}}Model{
		Model: base.NewModel(table.{{.StructName}}{}, db),
	}
}

`

func (c *Command) setGormCustomModel() error {
	wd := filepath.Join(c.wd, c.dir)
	basePkg := c.dir + "/base"
	tablePkg := c.dir + "/table"
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     c.tableName + ".go",
		TemplateFile: noCacheCustomModelTpl,
		Data: map[string]string{
			"package":    filepath.Base(c.dir),
			"nameSpace":  c.nameSpace,
			"StructName": c.structName,
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
