package model

import (
	"context"
	"gorm.io/gorm"
)

type IModel interface {
	First(ctx context.Context, first interface{}, id interface{}) error
	Create(ctx context.Context, info interface{}) error
	Update(ctx context.Context, info interface{}) error
	Delete(ctx context.Context, info interface{}) error

	Build(ctx context.Context) IModel
	Where(query string, args ...interface{}) IModel
	With(query string, args ...interface{}) IModel
	One(one interface{}) error
	List(list interface{}) error
	Paginate(list interface{}, page, pageSize int) (int64, error)
}

type Model struct {
	Conn           *gorm.DB
	BuildCondition *gorm.DB
	Table          interface{}
}

func (m *Model) First(ctx context.Context, first interface{}, id interface{}) error {
	return m.Conn.WithContext(ctx).First(first, id).Error
}
func (m *Model) Create(ctx context.Context, info interface{}) error {
	return m.Conn.WithContext(ctx).Create(info).Error
}
func (m *Model) Update(ctx context.Context, info interface{}) error {
	return m.Conn.WithContext(ctx).Save(info).Error
}
func (m *Model) Delete(ctx context.Context, info interface{}) error {
	return m.Conn.WithContext(ctx).Delete(info).Error
}

func (m *Model) Build(ctx context.Context) IModel {
	m.BuildCondition = m.Conn.WithContext(ctx).Model(m.Table)
	return m
}

func (m *Model) Where(query string, args ...interface{}) IModel {
	m.BuildCondition = m.BuildCondition.Where(query, args...)
	return m
}

func (m *Model) With(query string, args ...interface{}) IModel {
	m.BuildCondition = m.BuildCondition.Preload(query, args...)
	return m
}

func (m *Model) One(one interface{}) error {
	return m.BuildCondition.First(one).Error
}

func (m *Model) List(list interface{}) error {
	return m.BuildCondition.Find(list).Error
}

func (m *Model) Paginate(list interface{}, page, pageSize int) (int64, error) {
	var total int64
	err := m.BuildCondition.Count(&total).Error
	if total == 0 {
		return total, err
	}
	paginate := func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
	err = m.BuildCondition.Scopes(paginate).Find(list).Error
	return total, err
}
