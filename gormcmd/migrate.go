package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"os"
	"path/filepath"
)

var migrateTplCache = `package {{.package}}

import (
	"github.com/wuyan94zl/gotools/core/logz"
	"gorm.io/gorm"
	"{{.projectPkg}}/{{.modelPkgSrc}}"
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
	tbs = append(tbs, &{{.modelPkg}}.{{.StructName}}{})
	return tbs
}
`

func createMigrate(c *Command) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          c.wd,
		Filename:     "migrate.go",
		TemplateFile: migrateTplCache,
		Data: map[string]string{
			"package":     filepath.Base(c.wd),
			"StructName":  c.structName,
			"projectPkg":  c.projectPkg,
			"modelPkgSrc": "models/table",
			"modelPkg":    "table",
		},
	})
	if err != nil {
		fmt.Println("err：", err)
	}
	return err
}
func appendMigrate(c *Command) error {
	filePath := filepath.Join(c.wd, "migrate.go")
	code := fmt.Sprintf("tbs = append(tbs, &%s.%s{})", "table", c.structName)
	fileCode, err := utils.AppendFileCode(filePath, code, code, "return tbs")
	if err != nil {
		return err
	}
	code = fmt.Sprintf("\"%s/models/%s\"", c.projectPkg, "table")
	fileCode, err = utils.AppendStrCode(fileCode, code, code, "\"gorm.io/gorm\"")
	if err != nil {
		return err
	}
	return utils.WriteInfile(filePath, fileCode)
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
