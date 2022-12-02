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
)

func AutoMigrate(gorm *gorm.DB) {
	tables := setTables()
	err := gorm.AutoMigrate(tables...)
	if err != nil {
		logz.InfoAny("database tables auto migrate err：", err)
	}
}

func setTables() []interface{} {
	var tables []interface{}
	tables = append(tables, &{{.StructName}}{})
	return tables
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
	code := fmt.Sprintf("tables = append(tables, &%s{})", c.structName)
	i := strings.Index(fileStr, code)
	if i == -1 {
		point := strings.Index(fileStr, "return tables")
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
