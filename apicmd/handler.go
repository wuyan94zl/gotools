package apicmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
	"strings"
)

var handlerTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/gotools/core/response"
	{{if .isRequest}}"github.com/wuyan94zl/validator/v10"{{end}}

	{{if .isRequest}}"{{.typePackageSrc}}"{{end}}
	"{{.logicPackageSrc}}"
)

func {{.name}}Handler(c *gin.Context) {
	{{if .isRequest}}req := new({{.typePackage}}.{{.handler}}Request)
	c.ShouldBindJSON(req)
	validate := validator.New()
	err := validate.StructCtx(c.Copy(), req)
	if err != nil {
		c.JSON(200, response.NewError(500, err.Error()))
		return
	}
	resp, err := {{.logicPackage}}.New{{.LogicPackage}}(c).{{.name}}Logic(c, req){{else}}resp, err := {{.logicPackage}}.New{{.LogicPackage}}(c).{{.name}}Logic(c){{end}}
	switch err {
	case nil:
		c.JSON(200, response.NewSuccess(resp))
	case response.Error{}:
		c.JSON(200, err)
	default:
		c.JSON(500, err)
	}
}
`

func genHandler(c *Command) error {
	wd := filepath.Join(c.wd, "handler", c.dir)
	childDir := ""
	if c.dir != "" {
		childDir = "/" + c.dir
	}
	typePackage := fmt.Sprintf("%s/%s", c.projectPkg, "app/types")
	logicPackage := fmt.Sprintf("%s/%s%s", c.projectPkg, "app/logic", childDir)
	name := utils.UpperOne(c.name)

	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     strings.ToLower(c.name) + ".go",
		TemplateFile: handlerTpl,
		Data: map[string]string{
			"package":         filepath.Base(c.dir),
			"typePackageSrc":  typePackage,
			"logicPackageSrc": logicPackage,
			"typePackage":     filepath.Base(typePackage),
			"logicPackage":    filepath.Base(logicPackage),
			"LogicPackage":    utils.UpperOne(filepath.Base(logicPackage)),
			"name":            name,
			"handler":         c.handlerName,
			"isRequest":       c.isRequest,
		},
	})
}
