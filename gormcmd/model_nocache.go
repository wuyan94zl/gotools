package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"strings"
)

var noCacheModelTpl = `package {{.package}}

import (
	"context"
	"time"

	"gorm.io/gorm"
)

{{.struct}}

type (
	{{.structName}}Model interface {
		Insert(ctx context.Context, data *{{.StructName}}) (*{{.StructName}}, error)
		First(ctx context.Context, id int64) (*{{.StructName}}, error)
		Update(ctx context.Context, data *{{.StructName}}) error
		Delete(ctx context.Context, id int64) error
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

func (m *default{{.StructName}}Model) Insert(ctx context.Context, data *{{.StructName}}) (*{{.StructName}}, error) {
	err := m.Conn.WithContext(ctx).Create(data).Error
	return data, err
}

func (m *default{{.StructName}}Model) First(ctx context.Context, id int64) (*{{.StructName}}, error) {
	info := new({{.StructName}})
	err := m.Conn.WithContext(ctx).Find(info, id).Error
	return info, err
}

func (m *default{{.StructName}}Model) Update(ctx context.Context, data *{{.StructName}}) error {
	return m.Conn.WithContext(ctx).Save(data).Error
}

func (m *default{{.StructName}}Model) Delete(ctx context.Context, id int64) error {
	return m.Conn.WithContext(ctx).Where("id = ?", id).Delete(new({{.StructName}})).Error
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
