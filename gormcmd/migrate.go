package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"os"
	"path/filepath"
)

var migrateTplCache = `package {{.package}}

import (
	"gorm.io/gorm"
)

func AutoMigrate(gorm *gorm.DB) {
	tbs := setTables()
	err := gorm.AutoMigrate(tbs...)
	if err != nil {
		panic(err)
	}
}

func setTables() []interface{} {
	var tbs []interface{}
	tbs = append(tbs, &{{.StructName}}{})
	return tbs
}
`

func (c *Command) createMigrate(filePath string) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          filePath,
		Filename:     "migrate.go",
		TemplateFile: migrateTplCache,
		Data: map[string]string{
			"package":    filepath.Base(filePath),
			"StructName": c.structName,
			"projectPkg": c.nameSpace,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
	}
	return err
}
func (c *Command) appendMigrate(filePath string) error {
	filePath = filepath.Join(filePath, "migrate.go")
	code := fmt.Sprintf("tbs = append(tbs, &%s{})", c.structName)
	fileCode, err := utils.AppendFileCode(filePath, code, code, "return tbs")
	if err != nil {
		return err
	}
	return utils.WriteInfile(filePath, fileCode)
}

func (c *Command) setMigrate() error {
	filePath := filepath.Join(c.wd, c.dir, "table")
	_, err := os.Stat(filepath.Join(filePath, "migrate.go"))
	if err != nil {
		err = c.createMigrate(filePath)
	} else {
		err = c.appendMigrate(filePath)
	}
	return err
}
