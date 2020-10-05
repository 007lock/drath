package contract

import "context"

/*
OffsetPaginator for offset paging
	Page() uint64
	SetLimit(limit uint64)
	SetPage(page uint64)
	Results(ctx context.Context, table string, data interface{}) error
	Nums() uint64
	HasPages() bool
	HasNext() bool
	PrevPage() uint64
	NextPage() uint64
	HasPrev() bool
	PageNums() uint64
	SetOrderBy(key string, order string)
*/
type OffsetPaginator interface {
	Page() uint64
	SetLimit(limit uint64)
	SetPage(page uint64)
	Results(ctx context.Context, table string, data interface{}) error
	Nums() uint64
	HasPages() bool
	HasNext() bool
	PrevPage() uint64
	NextPage() uint64
	HasPrev() bool
	PageNums() uint64
	SetOrderBy(key string, order string)
}

/*
CursorPaginator for cursor paging
	SetLimit(limit uint64)
	Results(ctx context.Context, table string, data interface{}) error
	Nums() uint64
	SetOrderBy(key string, order string)
	SetAfterCursor(cursor string)
	SetBeforeCursor(cursor string)
	GetAfterCursor() string
	GetBeforeCursor() string
*/
type CursorPaginator interface {
	SetLimit(limit uint64)
	Results(ctx context.Context, table string, data interface{}) error
	Nums() uint64
	SetOrderBy(key string, order string)
	SetAfterCursor(cursor string)
	SetBeforeCursor(cursor string)
	GetAfterCursor() string
	GetBeforeCursor() string
}
