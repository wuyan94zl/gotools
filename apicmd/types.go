package apicmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var typesTpl = `package {{.package}}

{{if .isRequest}}type {{.name}}Request struct{}{{end}}

type {{.name}}Response struct{}

`

func genTypes(c *Command) error {
	fileName := strings.ToLower(getName(c.dir)) + ".go"
	filePath := filepath.Join(c.wd, "app", "types", fileName)

	_, err := os.Stat(filePath)
	if err == nil {
		return appendType(c, filePath)
	} else {
		return createType(c, filePath)
	}

}

func appendType(c *Command, filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileStr := string(file)

	if c.isRequest != "" {
		requestStr := fmt.Sprintf("%sRequest", c.handlerName)
		i := strings.Index(fileStr, requestStr)
		if i == -1 {
			fileStr = fmt.Sprintf("%s\ntype %s struct {}\n", fileStr, requestStr)
		}
	}

	responseStr := fmt.Sprintf("%sResponse", c.handlerName)
	i := strings.Index(fileStr, responseStr)
	if i == -1 {
		fileStr = fmt.Sprintf("%s\ntype %s struct {}", fileStr, responseStr)
	}
	return utils.WriteInfile(filePath, fileStr)
}

func createType(c *Command, filePath string) error {
	return utils.GenFileCover(utils.FileGenConfig{
		Dir:          filepath.Dir(filePath),
		Filename:     filepath.Base(filePath),
		TemplateFile: typesTpl,
		Data: map[string]string{
			"package":   filepath.Base(filepath.Dir(filePath)),
			"name":      c.handlerName,
			"isRequest": c.isRequest,
		},
	})
}
