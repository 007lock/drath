package paginator

import (
	"context"
	"encoding/base64"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
)

type fieldType string

const (
	fieldString fieldType = "STRING"
	fieldTime   fieldType = "TIME"
)

// Paginator structure
type gormCursorPaginator struct {
	maxPerPage       uint64
	afterCursor      string
	beforeCursor     string
	nextAfterCursor  string
	nextBeforeCursor string
	orderNameSpace   *string
	orderBy          string
	order            string
}

// New paginator constructor
func NewGormCursorPaginator(maxPerPage uint64) contract.CursorPaginator {
	if maxPerPage <= 0 {
		maxPerPage = 10
	}

	return &gormCursorPaginator{
		maxPerPage: maxPerPage,
		orderBy:    "CreatedAt",
		order:      "DESC",
	}
}

// Results stores the current page results into data argument which must be a pointer to a slice.
func (p *gormCursorPaginator) Results(ctx context.Context, table string, data interface{}) error {
	var cursor interface{}
	if p.hasAfterCursor() {
		cursor = p.decode(p.afterCursor)
	} else if p.hasBeforeCursor() {
		cursor = p.decode(p.beforeCursor)
	}
	tx := ctx.Value(constants.ContextKeyTransaction).(*gorm.DB)
	orderCon := fmt.Sprintf("%s %s", strcase.ToSnake(p.orderBy), p.order)
	if p.orderNameSpace != nil {
		orderCon = fmt.Sprintf("%s.%s", *p.orderNameSpace, orderCon)
	}
	stmt := tx.Table(table).Order(orderCon)
	if strings.HasPrefix(p.order, "FIELD") {
		stmt = tx.Table(table).Order(gorm.Expr(strings.ReplaceAll(p.order, "FIELD", "")))
	}
	if cursor != nil {
		orderCon := fmt.Sprintf("%s %s ?", strcase.ToSnake(p.orderBy), p.getOperator())
		if p.orderNameSpace != nil {
			orderCon = fmt.Sprintf("%s.%s", *p.orderNameSpace, orderCon)
		}
		stmt = stmt.Where(orderCon, cursor)
	}
	nextPage := p.maxPerPage + 1
	err := stmt.Limit(nextPage).Find(data).Error
	if err != nil {
		return err
	}
	if reflect.ValueOf(data).Elem().Type().Kind() == reflect.Slice && reflect.ValueOf(data).Elem().Len() > 0 {
		p.postProcess(data)
	}
	return nil
}

func (p *gormCursorPaginator) SetLimit(limit uint64) {
	p.maxPerPage = limit
}

func (p *gormCursorPaginator) SetOrderBy(key string, order string) {
	fields := strings.Split(key, ".")
	if len(fields) > 1 {
		p.orderBy = fields[1]
		p.orderNameSpace = &fields[0]
	} else {
		p.orderBy = fields[0]
	}
	p.order = order
}

func (p *gormCursorPaginator) SetAfterCursor(cursor string) {
	p.afterCursor = cursor
}

func (p *gormCursorPaginator) SetBeforeCursor(cursor string) {
	p.beforeCursor = cursor
}

func (p *gormCursorPaginator) Nums() uint64 {
	return p.maxPerPage
}

func (p *gormCursorPaginator) GetAfterCursor() string {
	return p.nextAfterCursor
}

func (p *gormCursorPaginator) GetBeforeCursor() string {
	return p.nextBeforeCursor
}

func (p *gormCursorPaginator) hasAfterCursor() bool {
	return p.afterCursor != ""
}

func (p *gormCursorPaginator) hasBeforeCursor() bool {
	return p.beforeCursor != ""
}

func (p *gormCursorPaginator) getOperator() string {
	if p.hasAfterCursor() {
		if p.order == "ASC" {
			return ">"
		}
		return "<"
	}
	if p.hasBeforeCursor() {
		if p.order == "ASC" {
			return "<"
		}
		return ">"
	}
	return "="
}

func (p *gormCursorPaginator) decode(cursor string) interface{} {
	bytes, err := base64.StdEncoding.DecodeString(cursor)

	if err != nil {
		return nil
	}
	field := string(bytes)

	return deconvert(field)
}

func (p *gormCursorPaginator) postProcess(out interface{}) {
	elems := reflect.ValueOf(out).Elem()
	// has more data after/before given cursor
	hasMore := uint64(elems.Len()) > p.maxPerPage

	if hasMore {
		elems.Set(elems.Slice(0, elems.Len()-1))
	}
	// reverse out in before cursor scenario
	if !p.hasAfterCursor() && p.hasBeforeCursor() {
		elems.Set(reverse(elems))
	}
	if p.hasBeforeCursor() || hasMore {
		p.nextAfterCursor = p.encode(elems.Index(elems.Len() - 1))
	}
	if p.hasAfterCursor() || (hasMore && p.hasBeforeCursor()) {
		p.nextBeforeCursor = p.encode(elems.Index(0))
	}
	return
}

func (p *gormCursorPaginator) encode(v reflect.Value) string {
	val := v
	if v.Kind() == reflect.Ptr {
		val = v.Elem()
	}
	field := convert(val.FieldByName(p.orderBy).Interface())
	return base64.StdEncoding.EncodeToString([]byte(field))
}

func convert(field interface{}) (result string) {
	switch field.(type) {
	case time.Time:
		result = fmt.Sprintf("%s?%s", field.(time.Time).UTC().Format(time.RFC3339Nano), fieldTime)
	default:
		result = fmt.Sprintf("%v?%s", field, fieldString)
	}
	return
}

func deconvert(field string) (result interface{}) {
	fieldTypeSepIndex := strings.LastIndex(field, "?")
	fieldType := fieldType(field[fieldTypeSepIndex+1:])
	field = field[:fieldTypeSepIndex]

	switch fieldType {
	case fieldTime:
		t, err := time.Parse(time.RFC3339Nano, field)

		if err != nil {
			t = time.Now().UTC()
		}
		result = t
	default:
		result = field
	}
	return
}

func reverse(v reflect.Value) reflect.Value {
	result := reflect.MakeSlice(v.Type(), 0, v.Cap())

	for i := v.Len() - 1; i >= 0; i-- {
		result = reflect.Append(result, v.Index(i))
	}
	return result
}
