package gormcmd

import (
	"github.com/wuyan94zl/gotools/core/utils"
	"path/filepath"
	"strings"
)

var noCacheModelTpl = `package {{.package}}

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
	"{{.projectPkg}}/models/types"
)

type (
	{{.StructName}}Model interface {
		Insert(ctx context.Context, info *types.{{.modelName}}) (*types.{{.modelName}}, error)
		First(ctx context.Context, id interface{}) (*types.{{.modelName}}, error)
		Update(ctx context.Context, info *types.{{.modelName}}) error
		Delete(ctx context.Context, info *types.{{.modelName}}) error

		Build(ctx context.Context) *Default{{.StructName}}Model
		Where(query string, args ...interface{}) *Default{{.StructName}}Model
		With(query string, args ...interface{}) *Default{{.StructName}}Model
		One() (*types.{{.modelName}}, error)
		List() ([]types.{{.modelName}}, error)
		Paginate(page, pageSize int) ([]types.{{.modelName}}, error)
	}
	Default{{.StructName}}Model struct {
		Conn           *gorm.DB
		Cache          *redis.Client
		BuildCondition *gorm.DB
	}
)

func New{{.StructName}}Model(gormDb *gorm.DB, cache *redis.Client) *Default{{.StructName}}Model {
	model := &Default{{.StructName}}Model{}
	model.Conn = gormDb
	model.Cache = cache
	return model
}

func (m *Default{{.StructName}}Model) Insert(ctx context.Context, info *types.{{.modelName}}) (*types.{{.modelName}}, error) {
	err := m.Conn.WithContext(ctx).Create(info).Error
	return info, err
}

func (m *Default{{.StructName}}Model) First(ctx context.Context, id interface{}) (*types.{{.modelName}}, error) {
	info := new(types.{{.modelName}})
	err := m.Conn.WithContext(ctx).Find(info, id).Error
	return info, err
}

func (m *Default{{.StructName}}Model) Update(ctx context.Context, info *types.{{.modelName}}) error {
	return m.Conn.WithContext(ctx).Save(info).Error
}

func (m *Default{{.StructName}}Model) Delete(ctx context.Context, info *types.{{.modelName}}) error {
	return m.Conn.WithContext(ctx).Delete(info).Error
}

func (m *Default{{.StructName}}Model) Build(ctx context.Context) *Default{{.StructName}}Model {
	m.BuildCondition = m.Conn.WithContext(ctx)
	return m
}

func (m *Default{{.StructName}}Model) Where(query string, args ...interface{}) *Default{{.StructName}}Model {
	m.BuildCondition = m.BuildCondition.Where(query, args...)
	return m
}

func (m *Default{{.StructName}}Model) With(query string, args ...interface{}) *Default{{.StructName}}Model {
	m.BuildCondition = m.BuildCondition.Preload(query, args...)
	return m
}

func (m *Default{{.StructName}}Model) One() (*types.{{.modelName}}, error) {
	info := new(types.{{.modelName}})
	err := m.BuildCondition.First(info).Error
	return info, err
}

func (m *Default{{.StructName}}Model) List() ([]types.{{.modelName}}, error) {
	var list []types.{{.modelName}}
	err := m.BuildCondition.Find(&list).Error
	return list, err
}

func (m *Default{{.StructName}}Model) Paginate(page, pageSize int) ([]types.{{.modelName}}, error) {
	paginate := func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
	var list []types.{{.modelName}}
	err := m.BuildCondition.Scopes(paginate).Find(&list).Error
	return list, err
}

`

func setGormNoCacheModel(data *Command) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          filepath.Join(data.wd, "_gen"),
		Filename:     data.packageName + ".go",
		TemplateFile: noCacheModelTpl,
		Data: map[string]string{
			"package":       "_gen",
			"projectPkg":    data.projectPkg,
			"struct":        data.structData,
			"StructName":    data.structName,
			"modelName":     data.structName + "Model",
			"structName":    strings.ToLower(data.structName[:1]) + data.structName[1:],
			"structImport":  data.structImport,
			"tableName":     data.tableName,
			"deletedFiled":  data.deletedFiled,
			"hasSoftDelete": data.hasSoftDelete,
		},
	})
	if err != nil {
		return err
	}
	err = createTypes(data)
	return err
}
