package handlercmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func genRoute(c *Command) error {
	if c.dir != "" {
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
	register := fmt.Sprintf("register%sHandler", utils.UpperOne(VarStringDir))
	return appendCode(filePath, register, "(app)")
}

func registerRoute(c *Command) error {
	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "router", c.dir+".go")
	_, err := os.Stat(filePath)
	if err == nil {
		return appendRoute(c, filePath)
	} else {
		return createRoute(c)
	}
}

func appendRoute(c *Command, filePath string) error {
	route := fmt.Sprintf("app.%s(\"%s\", handler.%sHandler)", c.method, c.routeUrl, c.handlerName)
	return appendCode(filePath, route, "")
}

var routeTpl = `package {{.package}}

import (
	"github.com/gin-gonic/gin"
	"{{.handlerSrc}}"
)

func register{{.dir}}Handler(app *gin.RouterGroup) {
	app.{{.method}}("{{.routeUrl}}", {{.handler}}.{{.handlerName}}Handler)
}

`

func createRoute(c *Command) error {
	wd, _ := os.Getwd()
	wd = filepath.Join(wd, "router")

	packageStr, err := utils.GetPackage()
	if err != nil {
		return err
	}

	handlerSrc := fmt.Sprintf("%s/%s", packageStr, "app/handler")
	return utils.GenFileCover(utils.FileGenConfig{
		Dir:          wd,
		Filename:     c.dir + ".go",
		TemplateFile: routeTpl,
		Data: map[string]string{
			"package":     filepath.Base(wd),
			"handlerSrc":  handlerSrc,
			"handler":     filepath.Base(handlerSrc),
			"handlerName": c.handlerName,
			"method":      c.method,
			"routeUrl":    c.routeUrl,
			"dir":         utils.UpperOne(c.dir),
		},
	})
}

func appendCode(filePath, code string, ext string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileStr := string(file)
	i := strings.Index(fileStr, code)
	if i == -1 {
		point := strings.Index(fileStr, "}")
		fileStr := fmt.Sprintf("%s\n\t%s%s%s", fileStr[0:point-1], code, ext, fileStr[point-1:])
		return utils.WriteInfile(filePath, fileStr)
	}
	return nil
}
