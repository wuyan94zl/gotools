package gorm

import (
	"fmt"
	"github.com/wuyan94zl/gotools/utils"
	"strings"
)

var modelTpl = `package {{.package}}

import (
	"database/sql"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/wuyan94zl/gotools/model"
	"gorm.io/gorm"
)

const (
	cacheKey = "modelCache:{{.StructName}}:"
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
		model.BashModel
	}
)

func new{{.StructName}}Model(gormDb *gorm.DB, cache *redis.Client) *default{{.StructName}}Model {
	m := &default{{.StructName}}Model{}
	m.Conn = gormDb
	m.Cache = cache
	return m
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

var modelTplCache = `package {{.package}}

import (
	"context"
	"database/sql"
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
		Insert(ctx context.Context, data *{{.StructName}}) (*{{.StructName}}, error)
		First(ctx context.Context, id int64) (*{{.StructName}}, error)
		Update(ctx context.Context, data *{{.StructName}}) error
		Delete(ctx context.Context, id int64) error
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

func (m *default{{.StructName}}Model) Insert(ctx context.Context, data *{{.StructName}}) (*{{.StructName}}, error) {
	err := m.Conn.Create(data).Error
	return data, err
}

func (m *default{{.StructName}}Model) First(ctx context.Context, id int64) (*{{.StructName}}, error) {
	key := fmt.Sprintf("%s%d", cacheKey, id)
	info := new({{.StructName}})
	err := m.CacheFirst(ctx, info, func() error {
		err := m.Conn.Find(info, id).Error
		_, err = m.empty(info, err)
		return err
	}, key)
	return info, err
}

func (m *default{{.StructName}}Model) Update(ctx context.Context, data *{{.StructName}}) error {
	key := fmt.Sprintf("%s%d", cacheKey, data.ID)
	return m.CacheUpdate(ctx, func() error {
		return m.Conn.Save(data).Error
	}, key)
}

func (m *default{{.StructName}}Model) Delete(ctx context.Context, id int64) error {
	key := fmt.Sprintf("%s%d", cacheKey, id)
	info, err := m.First(ctx, id)
	if err != nil {
		return err
	}
	return m.CacheDelete(ctx, func() error {
		return m.Conn.Delete(info, id).Error
	}, key)
}

func (m *default{{.StructName}}Model) empty(info *{{.StructName}}, err error) (*{{.StructName}}, error) {
	if err != nil {
		return nil, err
	}
	if info.ID != 0 {
		return info, nil
	}
	return nil, model.Notfound
}
`

func setGormModel(data *Command) error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          data.wd,
		Filename:     "model_gen.go",
		TemplateFile: modelTplCache,
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
