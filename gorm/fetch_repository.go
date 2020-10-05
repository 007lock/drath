package gorm

import (
	"context"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/jinzhu/gorm"
)

type gormFetchRepository struct{}

func NewGormFetchRepository() contract.FetchRepository {
	return &gormFetchRepository{}
}

func (r *gormFetchRepository) GetByRandom(c context.Context, table string, item interface{}, crit *contract.RepoCriterias) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	type random struct {
		ID uint64
	}
	tx = tx.Table(table).Order("random()")
	return tx.First(item).Error
}

func (r *gormFetchRepository) FetchByRandom(c context.Context, table string, item interface{}, crit *contract.RepoCriterias, limit uint64) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	type random struct {
		ID uint64
	}
	tx = tx.Table(table).Order("random()").Limit(limit)
	return tx.Find(item).Error
}

func (r *gormFetchRepository) FetchCursor(c context.Context, table string, item interface{}, p contract.CursorPaginator, crit *contract.RepoCriterias) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	ctx := context.WithValue(context.Background(), constants.ContextKeyTransaction, tx)
	err := p.Results(ctx, table, item)
	if err != nil {
		return err
	}

	return nil
}

func (r *gormFetchRepository) FetchPagination(c context.Context, table string, item interface{}, p contract.OffsetPaginator, crit *contract.RepoCriterias) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	ctx := context.WithValue(context.Background(), constants.ContextKeyTransaction, tx)
	err := p.Results(ctx, table, item)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return constants.DBError.ERROR_RECORD_NOT_FOUND
		}
		return err
	}
	return nil
}

func (r *gormFetchRepository) GetByID(c context.Context, table string, id string, item interface{}, crit *contract.RepoCriterias) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	err := tx.Table(table).First(item, "id = ?", id).Error
	if gorm.IsRecordNotFoundError(err) {
		return constants.DBError.ERROR_RECORD_NOT_FOUND
	}

	return err
}

func (r *gormFetchRepository) GetByCriteria(c context.Context, table string, item interface{}, crit *contract.RepoCriterias) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	err := tx.Table(table).First(item).Error
	if gorm.IsRecordNotFoundError(err) {
		return constants.DBError.ERROR_RECORD_NOT_FOUND
	}
	return err
}

func (r *gormFetchRepository) FetchByCriteria(c context.Context, table string, item interface{}, crit *contract.RepoCriterias) error {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}
	err := tx.Table(table).Find(item).Error
	if gorm.IsRecordNotFoundError(err) {
		return constants.DBError.ERROR_RECORD_NOT_FOUND
	}
	return err
}

func (r *gormFetchRepository) ApplyCriteria(c context.Context, crit *contract.RepoCriterias) context.Context {
	tx := c.Value(constants.ContextKeyTransaction).(*gorm.DB)
	if crit != nil {
		tx = criteriaApply(tx, crit)
	}

	return context.WithValue(context.Background(), constants.ContextKeyTransaction, tx)
}
