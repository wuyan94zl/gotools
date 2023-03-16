package apicmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"path/filepath"
	"strings"
)

var handlerTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"

	"{{.logicPackageSrc}}"
	{{if .isRequest}}"{{.typePackageSrc}}"{{end}}
	"{{.projectPkg}}/common/response"
)

// {{.nameCamelCase}}Handler
// @Summary todo {{.nameCamelCase}}Handler
// @Description todo 接口说明
// @Tags {{.tag}}
// @Security JwtAuth
{{.params}}{{.body}}// @Router /{{.route}} [{{.method}}]
{{.response}}
func {{.nameCamelCase}}Handler(c *gin.Context) {
	{{if .isRequest}}req := new({{.typePackage}}.{{.dirCamelCase}}{{.nameCamelCase}}Request)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(200, response.ValidateError(400, err, *req))
		return
	}
	{{end}}c.JSON(response.Result({{if .isRequest}}{{.logicPackage}}.New{{.LogicPackage}}(c).{{.nameCamelCase}}Logic(req){{else}}{{.logicPackage}}.New{{.LogicPackage}}(c).{{.nameCamelCase}}Logic(){{end}}))
}
`

func genHandler(c *Command) error {
	wd := filepath.Join(c.wd, "app", c.dir, "handler")
	childDir := ""
	if c.dir != "" {
		childDir = "/" + c.dir
	}
	typePackage := fmt.Sprintf("%s/%s/types", c.projectPkg, "app")
	logicPackage := fmt.Sprintf("%s/%s%s/logic", c.projectPkg, "app", childDir)

	paramCode := ""
	route := c.routeBase
	if c.routeUrl != "" {
		route = fmt.Sprintf("%s/%s", route, c.routeUrl)
	}
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
		body = fmt.Sprintf("// @Param request body %s.%s%sRequest to \"body params\"\n", "types", c.dirCamelCase, c.nameCamelCase)
	}
	response := fmt.Sprintf("// @Success 200 {object} %s.%s%sResponse", "types", c.dirCamelCase, c.nameCamelCase)

	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     strings.ToLower(c.nameCamelCase) + ".go",
		TemplateFile: handlerTpl,
		Data: map[string]string{
			"package":         filepath.Base(wd),
			"projectPkg":      c.projectPkg,
			"typePackageSrc":  typePackage,
			"logicPackageSrc": logicPackage,
			"typePackage":     filepath.Base(typePackage),
			"logicPackage":    filepath.Base(logicPackage),
			"LogicPackage":    utils.UpperFirst(filepath.Base(logicPackage)),
			"nameCamelCase":   c.nameCamelCase,
			"dirCamelCase":    c.dirCamelCase,
			"isRequest":       c.isRequest,
			"tag":             c.dirCamelCase,
			"method":          strings.ToLower(c.method),
			"route":           route,
			"params":          paramCode,
			"body":            body,
			"response":        response,
		},
	})
}
