package gorm

import (
	"context"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/jinzhu/gorm"
)

type gormUpdateRepository struct{}

func NewGormUpdateRepository() contract.UpdateRepository {
	return &gormUpdateRepository{}
}

func (r *gormUpdateRepository) Update(c context.Context, table string, item interface{}) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)

	return tx.Table(table).Model(item).Set("gorm:save_associations", false).Save(item).Error
}

func (r *gormUpdateRepository) UpdateByCriteria(c context.Context, table string, item interface{}, crit *contract.RepoCriterias) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	return tx.Table(table).Model(item).Updates(item).Error
}

func (r *gormUpdateRepository) Delete(c context.Context, table string, item interface{}, relations ...string) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	for i := range relations {
		tx.Model(item).Association(relations[i]).Clear()
	}
	return tx.Table(table).Delete(item).Error
}

func (r *gormUpdateRepository) DeleteByCriteria(c context.Context, table string, item interface{}, crit *contract.RepoCriterias, relations ...string) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	for i := range relations {
		tx.Table(table).Association(relations[i]).Clear()
	}
	return tx.Table(table).Delete(item).Error
}

func (r *gormUpdateRepository) Attach(c context.Context, associated interface{}, relation string, entities interface{}) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)

	return tx.Model(associated).Association(relation).Append(entities).Error
}

func (r *gormUpdateRepository) Dettach(c context.Context, associated interface{}, relation string, entities interface{}) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)

	return tx.Model(associated).Association(relation).Delete(entities).Error
}

func (r *gormUpdateRepository) DettachAll(c context.Context, associated interface{}, relation string) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)

	return tx.Model(associated).Association(relation).Clear().Error
}
