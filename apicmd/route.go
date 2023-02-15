package apicmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func genRoute(c *Command) error {
	if c.dirCamelCase != "" {
		err := appendRouteRegister(c)
		if err != nil {
			return err
		}
	} else {
		c.dir = "route"
	}
	err := registerRoute(c)
	if err != nil {
		return err
	}
	return nil
}

func appendRouteRegister(c *Command) error {
	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "router", "route.go")

	registerPackage := fmt.Sprintf("\"%s/app/%s\"", c.projectPkg, c.dir)
	appendCode(c, filePath, registerPackage, "", ")")
	register := fmt.Sprintf("%s.Register%sHandler", filepath.Base(c.dir), c.dirCamelCase)
	return appendCode(c, filePath, register, "(app)", "}")
}

func registerRoute(c *Command) error {
	filePath := filepath.Join(c.wd, "app", c.dir, "route.go")
	_, err := os.Stat(filePath)
	if err == nil {
		return appendRoute(c, filePath)
	} else {
		return createRoute(c, filePath)
	}
}

func appendRoute(c *Command, filePath string) error {
	route := fmt.Sprintf("app.%s(\"%s\", %s.%sHandler)", c.method, c.routeUrl, "handler", c.nameCamelCase)
	return appendCode(c, filePath, route, "", "}")
}

var routeTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"
	"{{.handlerPkgSrc}}"
)

func Register{{.dirCamelCase}}Handler(app gin.IRoutes) {
	app.{{.method}}("{{.routeUrl}}", {{.handler}}.{{.nameCamelCase}}Handler)
}

`

func createRoute(c *Command, filePath string) error {
	wd := filepath.Dir(filePath)
	handlerPkgSrc := fmt.Sprintf("%s/app/%s/handler", c.projectPkg, c.dir)
	return utils.GenFileCover(utils.FileGenConfig{
		Dir:          wd,
		Filename:     filepath.Base(filePath),
		TemplateFile: routeTpl,
		Data: map[string]string{
			"package":       filepath.Base(wd),
			"handlerPkgSrc": handlerPkgSrc,
			"handler":       "handler",
			"nameCamelCase": c.nameCamelCase,
			"method":        c.method,
			"routeUrl":      c.routeUrl,
			"dirCamelCase":  c.dirCamelCase,
		},
	})
}

func appendCode(c *Command, filePath, code, ext, find string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileStr := string(file)
	i := strings.Index(fileStr, code)
	if i == -1 {
		point := strings.Index(fileStr, find)
		fileStr := fmt.Sprintf("%s\n\t%s%s%s", fileStr[0:point-1], code, ext, fileStr[point-1:])
		return utils.WriteInfile(filePath, fileStr)
	}
	return nil
}
