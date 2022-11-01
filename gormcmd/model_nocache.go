package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
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
		Insert(ctx context.Context, info *{{.StructName}}) (*{{.StructName}}, error)
		First(ctx context.Context, id interface{}) (*{{.StructName}}, error)
		Update(ctx context.Context, info *{{.StructName}}) error
		Delete(ctx context.Context, info *{{.StructName}}) error
		{{if eq .hasSoftDelete "1"}}ForceDelete(ctx context.Context, info *{{.StructName}}) error{{end}}
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

func (m *default{{.StructName}}Model) Insert(ctx context.Context, info *{{.StructName}}) (*{{.StructName}}, error) {
	err := m.Conn.WithContext(ctx).Create(info).Error
	return info, err
}

func (m *default{{.StructName}}Model) First(ctx context.Context, id interface{}) (*{{.StructName}}, error) {
	info := new({{.StructName}})
	err := m.Conn.WithContext(ctx).Find(info, id).Error
	return info, err
}

func (m *default{{.StructName}}Model) Update(ctx context.Context, info *{{.StructName}}) error {
	return m.Conn.WithContext(ctx).Save(info).Error
}

func (m *default{{.StructName}}Model) Delete(ctx context.Context, info *{{.StructName}}) error {
	{{if eq .hasSoftDelete "1"}}return m.Conn.WithContext(ctx).Model(info).Update("{{.deletedFiled}}", time.Now().Unix()).Error{{else}}return m.Conn.WithContext(ctx).Delete(info).Error{{end}}
}

{{if eq .hasSoftDelete "1"}}func (m *default{{.StructName}}Model) ForceDelete(ctx context.Context, info *{{.StructName}}) error {
	return m.Conn.WithContext(ctx).Delete(info).Error
}{{end}}
`

func setGormNoCacheModel(data *Command) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     "model_gen.go",
		TemplateFile: noCacheModelTpl,
		Data: map[string]string{
			"package":       data.packageName,
			"struct":        data.structData,
			"StructName":    data.structName,
			"structName":    strings.ToLower(data.structName[:1]) + data.structName[1:],
			"deletedFiled":  data.deletedFiled,
			"hasSoftDelete": data.hasSoftDelete,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
