package model

import (
	"gorm.io/gorm"
)

type IModel interface {
	SetDB(db *gorm.DB) IModel
	First(dest interface{}) error
	List(list interface{}) error
	Paginate(list interface{}, page, pageSize int) (int64, error)
}

type Model struct {
	DB    *gorm.DB
	Table interface{}
}

func (m *Model) SetDB(db *gorm.DB) IModel {
	m.DB = db
	return m
}
func (m *Model) First(dest interface{}) error {
	return m.DB.First(dest).Error
}
func (m *Model) List(list interface{}) error {
	return m.DB.Find(list).Error
}
func (m *Model) Paginate(list interface{}, page, pageSize int) (int64, error) {
	var total int64
	err := m.DB.Count(&total).Error
	if total == 0 {
		return total, err
	}
	paginate := func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
	err = m.DB.Scopes(paginate).Find(list).Error
	return total, err
}
