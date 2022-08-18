package handlercmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"path/filepath"
	"strings"
)

var logicTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"

	"{{.typePackageSrc}}"
)

func {{.name}}Logic(c *gin.Context, req *{{.typePackage}}.{{.name}}Request) (*{{.typePackage}}.{{.name}}Response, error) {
	// todo logic code
	return &{{.typePackage}}.{{.name}}Response{}, nil
}

`

func genLogic(c *Command) error {
	packageStr, err := utils.GetPackage()
	if err != nil {
		return err
	}
	wd := filepath.Join(c.wd, "logic", c.dir)
	typePackage := fmt.Sprintf("%s/%s", packageStr, "app/types")

	fmt.Println(c.name, "name")
	return utils.GenFileCover(utils.FileGenConfig{
		Dir:          wd,
		Filename:     c.name + ".go",
		TemplateFile: logicTpl,
		Data: map[string]string{
			"package":        filepath.Base(wd),
			"typePackageSrc": typePackage,
			"typePackage":    filepath.Base(typePackage),
			"name":           strings.ToUpper(c.name[0:1]) + c.name[1:],
		},
	})
}
