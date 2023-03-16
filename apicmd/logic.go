package apicmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"path/filepath"
	"strings"
)

var logicTpl = `package {{.package}}

import (
	"{{.typePackageSrc}}"
)

func (logic *{{.Package}}) {{.nameCamelCase}}Logic({{if .isRequest}}req *{{.typePackage}}.{{.dirCamelCase}}{{.nameCamelCase}}Request{{end}}) (*{{.typePackage}}.{{.dirCamelCase}}{{.nameCamelCase}}Response, error) {
	// todo logic code
	return &{{.typePackage}}.{{.dirCamelCase}}{{.nameCamelCase}}Response{}, nil
}

`

func genLogic(c *Command) error {
	err := genBaseLogic(c)
	if err != nil {
		return err
	}

	wd := filepath.Join(c.wd, "app", c.dir, "logic")
	typePackage := fmt.Sprintf("%s/%s/types", c.projectPkg, "app")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     strings.ToLower(c.nameCamelCase) + ".go",
		TemplateFile: logicTpl,
		Data: map[string]string{
			"package":        filepath.Base(wd),
			"Package":        utils.UpperFirst(filepath.Base(wd)),
			"typePackageSrc": typePackage,
			"typePackage":    filepath.Base(typePackage),
			"nameCamelCase":  c.nameCamelCase,
			"dirCamelCase":   c.dirCamelCase,
			"isRequest":      c.isRequest,
		},
	})
}

var baseTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"{{.projectPkg}}/container"
)

type {{.Package}} struct {
	ctx *gin.Context
	db  *gorm.DB
}

func New{{.Package}}(c *gin.Context) *{{.Package}} {
	return &{{.Package}}{ctx: c, db: container.Instance().DB.WithContext(c)}
}

`

func genBaseLogic(c *Command) error {
	wd := filepath.Join(c.wd, "app", c.dir, "logic")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "base.go",
		TemplateFile: baseTpl,
		Data: map[string]string{
			"package":    filepath.Base(wd),
			"Package":    utils.UpperFirst(filepath.Base(wd)),
			"projectPkg": c.projectPkg,
		},
	})
}
