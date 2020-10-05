package paginator

import (
	"fmt"
	"math"
	"strings"

	"context"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/jinzhu/gorm"
)

// Paginator structure
type gormOffsetPaginator struct {
	maxPerPage     uint64
	page           uint64
	nums           uint64
	orderNameSpace *string
	orderBy        string
	order          string
}

// New paginator constructor
func NewGormOffsetPaginator(maxPerPage uint64) contract.OffsetPaginator {
	if maxPerPage <= 0 {
		maxPerPage = 10
	}

	return &gormOffsetPaginator{
		maxPerPage: maxPerPage,
		page:       1,
		nums:       0,
		orderBy:    "created_at",
		order:      "DESC",
	}
}

// SetPage set current page
func (p *gormOffsetPaginator) SetPage(page uint64) {
	if page <= 0 {
		page = 1
	}

	p.page = page
}

// Page returns current page
func (p *gormOffsetPaginator) Page() uint64 {
	pn := p.PageNums()
	if p.page > pn {
		return pn
	}

	return p.page
}

// Results stores the current page results into data argument which must be a pointer to a slice.
func (p *gormOffsetPaginator) Results(ctx context.Context, table string, data interface{}) error {
	tx := ctx.Value(constants.ContextKeyTransaction).(*gorm.DB)
	// Count offset first
	tx.Table(table).Model(data).Count(&p.nums)

	page := p.Page()
	offset := (page - 1) * p.maxPerPage
	// Return data
	if strings.HasPrefix(p.order, "FIELD") {
		return tx.Table(table).Order(gorm.Expr(strings.ReplaceAll(p.order, "FIELD", ""))).Limit(p.maxPerPage).Offset(offset).Find(data).Error
	}
	orderCon := fmt.Sprintf("%s %s", p.orderBy, p.order)
	if p.orderNameSpace != nil {
		orderCon = fmt.Sprintf("%s.%s", *p.orderNameSpace, orderCon)
	}
	return tx.Table(table).Order(fmt.Sprintf("%s %s", p.orderBy, p.order)).Limit(p.maxPerPage).Offset(offset).Find(data).Error
}

// Nums returns the total number of records
func (p *gormOffsetPaginator) Nums() uint64 {
	return p.nums
}

// Limit returns the total number of records
func (p *gormOffsetPaginator) Limit() uint64 {
	return p.maxPerPage
}

// HasPages returns true if there is more than one page
func (p *gormOffsetPaginator) HasPages() bool {
	return p.Nums() > p.maxPerPage
}

// HasNext returns true if current page is not the last page
func (p *gormOffsetPaginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

// PrevPage returns previous page number or ErrNoPrevPage if current page is first page
func (p *gormOffsetPaginator) PrevPage() uint64 {
	if !p.HasPrev() {
		return 0
	}

	return p.Page() - 1
}

// NextPage returns next page number or ErrNoNextPage if current page is last page
func (p *gormOffsetPaginator) NextPage() uint64 {
	if !p.HasNext() {
		return 0
	}

	return p.Page() + 1
}

// HasPrev returns true if current page is not the first page
func (p *gormOffsetPaginator) HasPrev() bool {
	return p.Page() > 1
}

// PageNums returns the total number of pages
func (p *gormOffsetPaginator) PageNums() uint64 {
	n := uint64(math.Ceil(float64(p.Nums()) / float64(p.maxPerPage)))
	if n == 0 {
		n = 1
	}

	return n
}

func (p *gormOffsetPaginator) SetLimit(limit uint64) {
	p.maxPerPage = limit
}

func (p *gormOffsetPaginator) SetOrderBy(key string, order string) {
	fields := strings.Split(key, ".")
	if len(fields) > 1 {
		p.orderBy = fields[1]
		p.orderNameSpace = &fields[0]
	} else {
		p.orderBy = fields[0]
	}
	p.order = order
}
