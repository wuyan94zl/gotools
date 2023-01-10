package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var migrateTplCache = `package {{.package}}

import (
	"github.com/wuyan94zl/gotools/core/logz"
	"gorm.io/gorm"

	"{{.projectPkg}}/models/tables"
)

func AutoMigrate(gorm *gorm.DB) {
	tbs := setTables()
	err := gorm.AutoMigrate(tbs...)
	if err != nil {
		logz.InfoAny("database tables auto migrate err：", err)
	}
}

func setTables() []interface{} {
	var tbs []interface{}
	tbs = append(tbs, &tables.{{.StructName}}{})
	return tbs
}
`

func createMigrate(c *Command) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     "migrate.go",
		TemplateFile: migrateTplCache,
		Data: map[string]string{
			"package":    filepath.Base(c.wd),
			"StructName": c.structName,
			"projectPkg": c.projectPkg,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
	}
	return err
}
func appendMigrate(c *Command) error {
	filePath := filepath.Join(c.wd, "migrate.go")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileStr := string(file)
	code := fmt.Sprintf("tbs = append(tbs, &tables.%s{})", c.structName)
	i := strings.Index(fileStr, code)
	if i == -1 {
		point := strings.Index(fileStr, "return tbs")
		fileStr = fmt.Sprintf("%s\t%s\n%s", fileStr[0:point-1], code, fileStr[point-1:])
		return utils.WriteInfile(filePath, fileStr)
	}
	return nil
}

func setMigrate(c *Command) error {
	_, err := os.Stat(filepath.Join(c.wd, "migrate.go"))
	if err != nil {
		err = createMigrate(c)
	} else {
		err = appendMigrate(c)
	}
	return err
}
