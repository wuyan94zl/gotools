package apicmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
	"strings"
)

var baseTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"
)

type {{.Package}} struct {
	ctx *gin.Context
}

func New{{.Package}}(c *gin.Context) *{{.Package}} {
	return &{{.Package}}{ctx: c}
}

`

var logicTpl = `package {{.package}}

import (
	"{{.typePackageSrc}}"
)

func (logic *{{.Package}}) {{.name}}Logic({{if .isRequest}}req *{{.typePackage}}.{{.handler}}Request{{end}}) (*{{.typePackage}}.{{.handler}}Response, error) {
	// todo logic code
	return &{{.typePackage}}.{{.handler}}Response{}, nil
}

`

func genLogic(c *Command) error {
	wd := filepath.Join(c.wd, "app", c.dir, "logic")
	typePackage := fmt.Sprintf("%s/%s", c.projectPkg, "app/types")
	name := utils.UpperOne(c.name)

	utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "base.go",
		TemplateFile: baseTpl,
		Data: map[string]string{
			"package": filepath.Base(wd),
			"Package": utils.UpperOne(filepath.Base(wd)),
		},
	})

	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     strings.ToLower(name) + ".go",
		TemplateFile: logicTpl,
		Data: map[string]string{
			"package":        filepath.Base(wd),
			"Package":        utils.UpperOne(filepath.Base(wd)),
			"typePackageSrc": typePackage,
			"typePackage":    filepath.Base(typePackage),
			"name":           name,
			"handler":        c.handlerName,
			"isRequest":      c.isRequest,
		},
	})
}
