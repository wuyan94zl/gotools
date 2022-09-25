package gormcmd

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"strings"
)

var modelTplCache = `package {{.package}}

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/wuyan94zl/gotools/model"
	"gorm.io/gorm"
)

const (
	cacheKey = "modelCache:{{.StructName}}:"
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
		model.BashModel
	}
)

func new{{.StructName}}Model(gormDb *gorm.DB, cache *redis.Client) *default{{.StructName}}Model {
	m := &default{{.StructName}}Model{}
	m.Conn = gormDb
	m.Cache = cache
	return m
}

func (m *default{{.StructName}}Model) Insert(ctx context.Context, info *{{.StructName}}) (*{{.StructName}}, error) {
	err := m.Conn.WithContext(ctx).Create(info).Error
	return info, err
}

func (m *default{{.StructName}}Model) First(ctx context.Context, id interface{}) (*{{.StructName}}, error) {
	key := fmt.Sprintf("%s%v", cacheKey, id)
	info := new({{.StructName}})
	err := m.CacheFirst(ctx, info, func() error {
		{{if eq .hasSoftDelete "1"}}err := m.Conn.WithContext(ctx).Where("{{.deletedFiled}} = 0").First(info, id).Error{{else}}err := m.Conn.WithContext(ctx).First(info, id).Error{{end}}
		return err
	}, key)
	return info, err
}

func (m *default{{.StructName}}Model) Update(ctx context.Context, info *{{.StructName}}) error {
	key := fmt.Sprintf("%s%d", cacheKey, info.ID)
	return m.CacheUpdate(ctx, func() error {
		return m.Conn.WithContext(ctx).Save(info).Error
	}, key)
}

func (m *default{{.StructName}}Model) Delete(ctx context.Context, info *{{.StructName}}) error {
	key := fmt.Sprintf("%s%v", cacheKey, info.ID)
	return m.CacheDelete(ctx, func() error {
		{{if eq .hasSoftDelete "1"}}return m.Conn.WithContext(ctx).Model(info).Update("{{.deletedFiled}}", time.Now().Unix()).Error{{else}}return m.Conn.WithContext(ctx).Delete(info).Error{{end}}
	}, key)
}

{{if eq .hasSoftDelete "1"}}func (m *default{{.StructName}}Model) ForceDelete(ctx context.Context, info *{{.StructName}}) error {
	key := fmt.Sprintf("%s%v", cacheKey, info.ID)
	return m.CacheDelete(ctx, func() error {
		return m.Conn.WithContext(ctx).Delete(info).Error
	}, key)
}{{end}}
`

func setGormModel(data *Command) error {
	err := utils.GenFileCover(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     "model_gen.go",
		TemplateFile: modelTplCache,
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
