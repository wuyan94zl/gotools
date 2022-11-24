package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
	"strings"
)

var noCacheModelTpl = `package {{.package}}

import (
	"context"
	"time"
	
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

{{.struct}}

func (m *{{.StructName}}) TableName() string {
	return "{{.tableName}}"
}

type (
	{{.StructName}}Model interface {
		Insert(ctx context.Context, info *{{.StructName}}) (*{{.StructName}}, error)
		First(ctx context.Context, id interface{}) (*{{.StructName}}, error)
		Update(ctx context.Context, info *{{.StructName}}) error
		Delete(ctx context.Context, info *{{.StructName}}) error
		{{if eq .hasSoftDelete "1"}}ForceDelete(ctx context.Context, info *{{.StructName}}) error{{end}}
	}
	Default{{.StructName}}Model struct {
		Conn *gorm.DB
		Cache *redis.Client
	}
)

func New{{.StructName}}Model(gormDb *gorm.DB, cache *redis.Client) *Default{{.StructName}}Model {
	model := &Default{{.StructName}}Model{}
	model.Conn = gormDb
	model.Cache = cache
	return model
}

func (m *Default{{.StructName}}Model) Insert(ctx context.Context, info *{{.StructName}}) (*{{.StructName}}, error) {
	err := m.Conn.WithContext(ctx).Create(info).Error
	return info, err
}

func (m *Default{{.StructName}}Model) First(ctx context.Context, id interface{}) (*{{.StructName}}, error) {
	info := new({{.StructName}})
	err := m.Conn.WithContext(ctx).Find(info, id).Error
	return info, err
}

func (m *Default{{.StructName}}Model) Update(ctx context.Context, info *{{.StructName}}) error {
	return m.Conn.WithContext(ctx).Save(info).Error
}

func (m *Default{{.StructName}}Model) Delete(ctx context.Context, info *{{.StructName}}) error {
	{{if eq .hasSoftDelete "1"}}return m.Conn.WithContext(ctx).Model(info).Update("{{.deletedFiled}}", time.Now().Unix()).Error{{else}}return m.Conn.WithContext(ctx).Delete(info).Error{{end}}
}

{{if eq .hasSoftDelete "1"}}func (m *Default{{.StructName}}Model) ForceDelete(ctx context.Context, info *{{.StructName}}) error {
	return m.Conn.WithContext(ctx).Delete(info).Error
}{{end}}
`

func setGormNoCacheModel(data *Command) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     data.packageName + "_gen.go",
		TemplateFile: noCacheModelTpl,
		Data: map[string]string{
			"package":       filepath.Base(data.wd),
			"projectPkg":    data.projectPkg,
			"struct":        data.structData,
			"StructName":    data.structName,
			"structName":    strings.ToLower(data.structName[:1]) + data.structName[1:],
			"tableName":     data.tableName,
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
