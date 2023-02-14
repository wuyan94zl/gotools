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

func createTables(c *Command) error {
	filePath := filepath.Join(c.wd, "table", c.tableName+".go")
	_, err := os.Stat(filePath)
	custom := `// edit the custom code below, never delete this line of code

type ` + c.structName + `Model struct {
	` + c.structName + `
}`
	if err == nil {
		custom, _ = getCustomModel(filePath)
	}
	return createType(c, filePath, custom)
}

func createType(c *Command, filePath, customModel string) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          filepath.Dir(filePath),
		Filename:     filepath.Base(filePath),
		TemplateFile: typeTpl,
		Data: map[string]string{
			"package":      filepath.Base(filepath.Dir(filePath)),
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
