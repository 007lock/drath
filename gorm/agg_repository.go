package gorm

import (
	"context"
	"fmt"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/jinzhu/gorm"
)

type gormAggRepository struct{}

func NewGormAggregationRepository() contract.AggregationRepository {
	return &gormAggRepository{}
}

func (r *gormAggRepository) CountUnique(c context.Context, table string, field string, name string) (uint64, error) {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	var total uint64
	err := tx.Table(table).Where(fmt.Sprintf("%s LIKE ?", field), name).Count(&total).Error
	return total, err
}

func (r *gormAggRepository) CountByCriteria(c context.Context, table string, crit *contract.RepoCriterias) (uint64, error) {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	var total uint64
	err := tx.Table(table).Count(&total).Error
	return total, err
}

func (r *gormAggRepository) SumByCriteria(c context.Context, table string, field string, crit *contract.RepoCriterias) (uint64, error) {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	type NResult struct {
		N uint64 //or int ,or some else
	}
	var n NResult
	err := tx.Table(table).Select(fmt.Sprintf("sum(%s) as n", field)).Scan(&n).Error
	return n.N, err
}
