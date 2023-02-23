package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type IModel interface {
	First(ctx context.Context, first interface{}, id interface{}) error
	Create(ctx context.Context, info interface{}) error
	Update(ctx context.Context, info interface{}) error
	Delete(ctx context.Context, info interface{}) error
	FindByField(ctx context.Context, info any, filed string, val any) error

	DB() *gorm.DB
	ConditionWhere(query string, args ...interface{}) IModel
	ConditionJoins(query string, args ...interface{}) IModel
	ConditionWith(query string, args ...interface{}) IModel
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
func (m *Model) FindByField(ctx context.Context, info any, filed string, val any) error {
	return m.Conn.WithContext(ctx).Where(fmt.Sprintf("%s = ?", filed), val).First(info).Error
}

func (m *Model) DB() *gorm.DB {
	return m.BuildCondition
}
func (m *Model) ConditionWhere(query string, args ...interface{}) IModel {
	m.BuildCondition = m.BuildCondition.Where(query, args...)
	return m
}
func (m *Model) ConditionJoins(query string, args ...interface{}) IModel {
	m.BuildCondition = m.BuildCondition.Joins(query, args...)
	return m
}
func (m *Model) ConditionWith(query string, args ...interface{}) IModel {
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
