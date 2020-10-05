package contract

import "context"

// Database interface
type Database interface {
	Begin() (interface{}, error)
	Get() (interface{}, error)
	MigrationUp() error
	MigrationDown() error
	Close() error
}

/*
FetchRepository include all function for fetch from object only:
	GetByRandom(c context.Context, table string, item interface{}, crit *RepoCriterias) error
	GetByID(c context.Context, table string, id string, item interface{}, crit *RepoCriterias) error
	GetByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error
	FetchByRandom(c context.Context, table string, item interface{}, crit *RepoCriterias, limit uint64) error
	FetchCursor(c context.Context, table string, item interface{}, p CursorPaginator, crit *RepoCriterias) error
	FetchPagination(c context.Context, table string, item interface{}, p OffsetPaginator, crit *RepoCriterias) error
	FetchByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error
*/
type FetchRepository interface {
	GetByRandom(c context.Context, table string, item interface{}, crit *RepoCriterias) error
	GetByID(c context.Context, table string, id string, item interface{}, crit *RepoCriterias) error
	GetByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error
	FetchByRandom(c context.Context, table string, item interface{}, crit *RepoCriterias, limit uint64) error
	FetchCursor(c context.Context, table string, item interface{}, p CursorPaginator, crit *RepoCriterias) error
	FetchPagination(c context.Context, table string, item interface{}, p OffsetPaginator, crit *RepoCriterias) error
	FetchByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error

	ApplyCriteria(c context.Context, crit *RepoCriterias) context.Context
}

/*
AggregationRepository include all function for aggregation object only:
	CountUnique(c context.Context, table string, field string, name string) (uint64, error)
	CountByCriteria(c context.Context, table string, crit *RepoCriterias) (uint64, error)
	SumByCriteria(c context.Context, table string, field string, crit *RepoCriterias) (uint64, error)
*/
type AggregationRepository interface {
	CountUnique(c context.Context, table string, field string, name string) (uint64, error)
	CountByCriteria(c context.Context, table string, crit *RepoCriterias) (uint64, error)
	SumByCriteria(c context.Context, table string, field string, crit *RepoCriterias) (uint64, error)
}

/*
UpdateRepository include all function for update object only:
	Update(c context.Context, table string, item interface{}) error
	UpdateByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error
	Delete(c context.Context, table string, item interface{}, relations ...string) error
	DeleteByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias, relations ...string) error
	Attach(c context.Context, associated interface{}, relation string, entities interface{}) error
	Dettach(c context.Context, associated interface{}, relation string, entities interface{}) error
	DettachAll(c context.Context, associated interface{}, relation string) error
*/
type UpdateRepository interface {
	Update(c context.Context, table string, item interface{}) error
	UpdateByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error
	Delete(c context.Context, table string, item interface{}, relations ...string) error
	DeleteByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias, relations ...string) error
	Attach(c context.Context, associated interface{}, relation string, entities interface{}) error
	Dettach(c context.Context, associated interface{}, relation string, entities interface{}) error
	DettachAll(c context.Context, associated interface{}, relation string) error
}

/*
InsertRepository include all function for insert object only:
	StoreOrUpdate(c context.Context, table string, item interface{}) error
	Store(c context.Context, table string, item interface{}) error
*/
type InsertRepository interface {
	StoreOrUpdate(c context.Context, table string, item interface{}) error
	Store(c context.Context, table string, item interface{}) error
}
