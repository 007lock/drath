package gorm

import (
	"context"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/jinzhu/gorm"
)

type gormInsertRepository struct{}

func NewGormInsertRepository() contract.InsertRepository {
	return &gormInsertRepository{}
}

func (r *gormInsertRepository) Store(c context.Context, table string, item interface{}) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	err := tx.Table(table).Create(item).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *gormInsertRepository) StoreOrUpdate(c context.Context, table string, item interface{}) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	return tx.Table(table).Save(item).Error
}
