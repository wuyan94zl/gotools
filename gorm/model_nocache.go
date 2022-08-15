package gorm

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"strings"
)

var noCacheModelTpl = `package {{.package}}

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var NotFoundErr = errors.New("数据不存在")

{{.struct}}

type (
	{{.structName}}Model interface {
		Insert(data *{{.StructName}}) (*{{.StructName}}, error)
		First(id int64) (*{{.StructName}}, error)
		Update(data *{{.StructName}}) error
		Delete(id int64) error
	}
	default{{.StructName}}Model struct {
		Conn *gorm.DB
	}
)

func new{{.StructName}}Model(gormDb *gorm.DB) *default{{.StructName}}Model {
	model := &default{{.StructName}}Model{}
	model.Conn = gormDb
	return model
}

func (m *default{{.StructName}}Model) Insert(data *{{.StructName}}) (*{{.StructName}}, error) {
	err := m.Conn.Create(data).Error
	return m.empty(data, err)
}

func (m *default{{.StructName}}Model) First(id int64) (*{{.StructName}}, error) {
	info := new({{.StructName}})
	err := m.Conn.Find(info, id).Error
	return m.empty(info, err)
}

func (m *default{{.StructName}}Model) Update(data *{{.StructName}}) error {
	return m.Conn.Save(data).Error
}

func (m *default{{.StructName}}Model) Delete(id int64) error {
	return m.Conn.Where("id = ?", id).Delete(new({{.StructName}})).Error
}

func (m *default{{.StructName}}Model) empty(info *{{.StructName}}, err error) (*{{.StructName}}, error) {
	if err != nil {
		return nil, err
	}
	if info.ID != 0 {
		return info, nil
	}
	return nil, NotFoundErr
}
`

func setGormNoCacheModel(data *Command) error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     "model_gen.go",
		TemplateFile: noCacheModelTpl,
		Data: map[string]string{
			"package":    data.packageName,
			"struct":     data.structData,
			"StructName": data.structName,
			"structName": strings.ToLower(data.structName[:1]) + data.structName[1:],
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
