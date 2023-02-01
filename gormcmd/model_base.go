package gormcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
)

var modelBaseTpl = `package base

import (
	"github.com/wuyan94zl/gotools/core/model"
	"gorm.io/gorm"
)

type IBase interface {
	model.IModel
}

type Base struct {
	model.Model
}

func NewBase(table interface{}, conn *gorm.DB) *Base {
	return &Base{model.Model{Conn: conn, Table: table}}
}`

func setGormBaseModel(data *Command) error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          filepath.Join(data.wd, "base"),
		Filename:     "base.go",
		TemplateFile: modelBaseTpl,
		Data:         map[string]string{},
	})
	if err != nil {
		return err
	}
	err = createTypes(data)
	return err
}