package handlercmd

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
	"github.com/wuyan94zl/validator/v10"

	"{{.typePackageSrc}}"
	"{{.logicPackageSrc}}"
)

func {{.handler}}Handler(c *gin.Context) {
	req := new({{.typePackage}}.{{.handler}}Request)
	c.ShouldBindJSON(req)
	validate := validator.New()
	err := validate.StructCtx(c.Copy(), req)
	if err != nil {
		c.JSON(200, response.Error(500, err))
		return
	}
	resp, err := {{.logicPackage}}.{{.name}}Logic(c, req)
	if err != nil {
		c.JSON(200, response.Error(500, err))
		return
	}
	c.JSON(200, response.Success(resp))
}
`

func genHandler(c *Command) error {
	packageStr, err := utils.GetPackage()
	if err != nil {
		return err
	}
	wd := filepath.Join(c.wd, "handler")
	childDir := ""
	if c.dir != "" {
		childDir = "/" + c.dir
	}
	typePackage := fmt.Sprintf("%s/%s", packageStr, "app/types")
	logicPackage := fmt.Sprintf("%s/%s%s", packageStr, "app/logic", childDir)
	name := c.handlerName
	if c.dir != "" {
		name = c.handlerName[len(c.dir):]
	}

	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     strings.ToLower(c.handlerName) + ".go",
		TemplateFile: handlerTpl,
		Data: map[string]string{
			"package":         filepath.Base(wd),
			"typePackageSrc":  typePackage,
			"logicPackageSrc": logicPackage,
			"typePackage":     filepath.Base(typePackage),
			"logicPackage":    filepath.Base(logicPackage),
			"name":            name,
			"handler":         c.handlerName,
		},
	})
}
