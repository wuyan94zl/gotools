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

	{{if .isRequest}}"{{.typePackageSrc}}"{{end}}
	"{{.logicPackageSrc}}"
)

// {{.name}}Handler
// @Summary todo {{.name}}Handler
// @Description todo 接口说明
// @Tags {{.tag}}
{{.params}}{{.body}}// @Router /{{.route}} [{{.method}}]
func {{.name}}Handler(c *gin.Context) {
	{{if .isRequest}}req := new({{.typePackage}}.{{.handler}}Request)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(200, response.ValidateError(500, err, *req))
		return
	}
	resp, err := {{.logicPackage}}.New{{.LogicPackage}}(c).{{.name}}Logic(req){{else}}resp, err := {{.logicPackage}}.New{{.LogicPackage}}(c).{{.name}}Logic(){{end}}
	switch err.(type) {
	case nil:
		c.JSON(200, response.NewSuccess(resp))
	case response.Error:
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
	paramCode := ""
	route := c.routeUrl
	if c.params != "" {
		params := strings.Split(c.params, "/")
		for i := 0; i < len(params); i++ {
			v := params[i][1:]
			paramCode = fmt.Sprintf("%s %s %s %s \"%s\"\n", paramCode, "// @Param", v, "path string to", v)
			route = strings.Replace(route, ":"+v, "{"+v+"}", 1)
		}
	}
	body := ""
	if c.isRequest == "true" {
		body = fmt.Sprintf("// @Param request body %s.%sRequest to \"body params\"\n", filepath.Base(typePackage), c.handlerName)
	}

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
			"tag":             "api",
			"method":          strings.ToLower(c.method),
			"route":           route,
			"params":          paramCode,
			"body":            body,
		},
	})
}
