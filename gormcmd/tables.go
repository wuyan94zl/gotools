package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var typeTpl = `package {{.package}}

import (
	{{.structImport}}
)

{{.structData}}


{{.customModel}}
`

func (c *Command) createTables() error {
	filePath := filepath.Join(c.wd, c.dir, "table", c.tableName+".go")
	_, err := os.Stat(filePath)
	custom := `// edit the custom code below, never delete this line of code

type ` + c.structName + `Model struct {
	` + c.structName + `
}`
	if err == nil {
		custom, _ = c.getCustomModel(filePath)
	}
	return c.createType(filePath, custom)
}

func (c *Command) createType(filePath, customModel string) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          filepath.Dir(filePath),
		Filename:     filepath.Base(filePath),
		TemplateFile: typeTpl,
		Data: map[string]string{
			"package":      filepath.Base(filepath.Dir(filePath)),
			"structData":   c.structData,
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
func (c *Command) getCustomModel(filePath string) (string, error) {
	f, err := os.Open(filePath)
	file, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	fileStr := string(file)
	code := "// edit the custom code below, never delete this line of code"
	i := strings.Index(fileStr, code)
	return fileStr[i-1:], nil
}
