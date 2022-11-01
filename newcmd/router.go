package newcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
)

var genRouterTpl = `package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(app *gin.RouterGroup) {
}
`

func genRouter(c *Command) error {
	wd := filepath.Join(c.wd, "router")
	return utils.GenFile(utils.FileGenConfig{
		Dir:          wd,
		Filename:     "route.go",
		TemplateFile: genRouterTpl,
		Data:         map[string]string{},
	})
}
