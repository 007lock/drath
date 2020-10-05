package paginator

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

type ModelBasis struct {
	ID uint64
}

func TestNewGormCursorPaginator(t *testing.T) {
	maxPerPage := uint64(10)
	type args struct {
		maxPerPage uint64
	}
	tests := []struct {
		name string
		args args
		want contract.CursorPaginator
	}{
		{
			name: "NewGormRequest Success",
			args: args{maxPerPage},
			want: &gormCursorPaginator{
				maxPerPage: maxPerPage,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormCursorPaginator(tt.args.maxPerPage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGormCursorPaginator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createGormSession() (*gorm.DB, sqlmock.Sqlmock, error) {
	// Init mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	DB, err := gorm.Open("postgres", db)
	if err != nil {
		return nil, mock, err
	}
	return DB, mock, nil
}

func Test_gormCursorPaginator_Results(t *testing.T) {
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	rows := mock.NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		ctx   context.Context
		table string
		data  interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "gormCursorPaginator_Results Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			args: args{
				table: "table",
				ctx:   context.WithValue(context.Background(), constants.ContextKeyTransaction, db),
				data:  &[]ModelBasis{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if err := p.Results(tt.args.ctx, tt.args.table, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("gormCursorPaginator.Results() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormCursorPaginator_SetLimit(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		limit uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gormCursorPaginator_SetLimit Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			args: args{
				limit: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			p.SetLimit(tt.args.limit)
		})
	}
}

func Test_gormCursorPaginator_SetOrderBy(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		key   string
		order string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gormCursorPaginator_SetOrderBy Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			args: args{
				key:   "order",
				order: "desc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			p.SetOrderBy(tt.args.key, tt.args.order)
		})
	}
}

func Test_gormCursorPaginator_SetAfterCursor(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		cursor string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gormCursorPaginator_SetAfterCursor Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			args: args{
				cursor: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			p.SetAfterCursor(tt.args.cursor)
		})
	}
}

func Test_gormCursorPaginator_SetBeforeCursor(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		cursor string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gormCursorPaginator_SetBeforeCursor Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			args: args{
				cursor: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			p.SetBeforeCursor(tt.args.cursor)
		})
	}
}

func Test_gormCursorPaginator_Nums(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "gormCursorPaginator_Nums Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.Nums(); got != tt.want {
				t.Errorf("gormCursorPaginator.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormCursorPaginator_GetAfterCursor(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "gormCursorPaginator_GetAfterCursor Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.GetAfterCursor(); got != tt.want {
				t.Errorf("gormCursorPaginator.GetAfterCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormCursorPaginator_GetBeforeCursor(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "gormCursorPaginator_GetBeforeCursor Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.GetBeforeCursor(); got != tt.want {
				t.Errorf("gormCursorPaginator.GetBeforeCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormCursorPaginator_hasAfterCursor(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "gormCursorPaginator_hasAfterCursor Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.hasAfterCursor(); got != tt.want {
				t.Errorf("gormCursorPaginator.hasAfterCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormCursorPaginator_hasBeforeCursor(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "gormCursorPaginator_hasBeforeCursor Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.hasBeforeCursor(); got != tt.want {
				t.Errorf("gormCursorPaginator.hasBeforeCursor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormCursorPaginator_getOperator(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "gormCursorPaginator_getOperator equal",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			want: "=",
		},
		{
			name: "gormCursorPaginator_getOperator order asc & has after cursor",
			fields: fields{
				maxPerPage:  1,
				orderBy:     "CreatedAt",
				order:       "ASC",
				afterCursor: "afterCursor",
			},
			want: ">",
		},
		{
			name: "gormCursorPaginator_getOperator order desc & has after cursor",
			fields: fields{
				maxPerPage:  1,
				orderBy:     "CreatedAt",
				order:       "DESC",
				afterCursor: "afterCursor",
			},
			want: "<",
		},
		{
			name: "gormCursorPaginator_getOperator order asc & has before cursor",
			fields: fields{
				maxPerPage:   1,
				orderBy:      "CreatedAt",
				order:        "ASC",
				beforeCursor: "beforeCursor",
			},
			want: "<",
		},
		{
			name: "gormCursorPaginator_getOperator order desc & has before cursor",
			fields: fields{
				maxPerPage:   1,
				orderBy:      "CreatedAt",
				order:        "DESC",
				beforeCursor: "beforeCursor",
			},
			want: ">",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.getOperator(); got != tt.want {
				t.Errorf("gormCursorPaginator.getOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormCursorPaginator_decode(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		cursor string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "gormCursorPaginator_decode Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			args: args{
				cursor: "cursor",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.decode(tt.args.cursor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gormCursorPaginator.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormCursorPaginator_postProcess(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		out interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gormCursorPaginator_decode Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "CreatedAt",
				order:      "DESC",
			},
			args: args{
				out: &[]ModelBasis{
					ModelBasis{ID: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			p.postProcess(tt.args.out)
		})
	}
}

func Test_gormCursorPaginator_encode(t *testing.T) {
	type fields struct {
		maxPerPage       uint64
		afterCursor      string
		beforeCursor     string
		nextAfterCursor  string
		nextBeforeCursor string
		orderBy          string
		order            string
	}
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "gormCursorPaginator_encode Success",
			fields: fields{
				maxPerPage: 1,
				orderBy:    "ID",
				order:      "DESC",
			},
			args: args{
				v: reflect.ValueOf(&ModelBasis{ID: 1}),
			},
			want: "MT9TVFJJTkc=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormCursorPaginator{
				maxPerPage:       tt.fields.maxPerPage,
				afterCursor:      tt.fields.afterCursor,
				beforeCursor:     tt.fields.beforeCursor,
				nextAfterCursor:  tt.fields.nextAfterCursor,
				nextBeforeCursor: tt.fields.nextBeforeCursor,
				orderBy:          tt.fields.orderBy,
				order:            tt.fields.order,
			}
			if got := p.encode(tt.args.v); got != tt.want {
				t.Errorf("gormCursorPaginator.encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	now := time.Now()
	type args struct {
		field interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "gormCursorPaginator_convert time success",
			args: args{
				field: now,
			},
			wantResult: fmt.Sprintf("%s?%s", now.UTC().Format(time.RFC3339Nano), fieldTime),
		},
		{
			name: "gormCursorPaginator_convert string success",
			args: args{
				field: "field",
			},
			wantResult: fmt.Sprintf("field?%s", fieldString),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := convert(tt.args.field); gotResult != tt.wantResult {
				t.Errorf("convert() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_deconvert(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name       string
		args       args
		wantResult interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := deconvert(tt.args.field); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("deconvert() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
