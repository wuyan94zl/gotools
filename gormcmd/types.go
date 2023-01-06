package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var typeTpl = `package {{.package}}

import (
	{{.structImport}}
)

{{.struct}}


{{.customModel}}
`

func createTypes(c *Command) error {
	filePath := filepath.Join(c.wd, "types", strings.ToLower(c.structName)+".go")
	_, err := os.Stat(filePath)
	custom := `// edit the custom code below, never delete this line of code

type ` + c.structName + `Model struct {
	` + c.structName + `
}`
	if err == nil {
		custom, _ = getCustomModel(filePath)
	}
	return createType(c, custom)
}

func createType(c *Command, customModel string) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          filepath.Join(c.wd, "types"),
		Filename:     strings.ToLower(c.structName) + ".go",
		TemplateFile: typeTpl,
		Data: map[string]string{
			"package":      "types",
			"struct":       c.structData,
			"structImport": c.structImport,
			"structName":   c.structName,
			"customModel":  customModel,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
	}
	return err
}
func getCustomModel(filePath string) (string, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	fileStr := string(file)
	code := "// edit the custom code below, never delete this line of code"
	i := strings.Index(fileStr, code)
	return fileStr[i-1:], nil
}
