package paginator

import (
	"context"
	"reflect"
	"testing"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
)

var maxPerPage = uint64(10)

func TestNewGormOffsetPaginator(t *testing.T) {
	type args struct {
		maxPerPage uint64
	}
	tests := []struct {
		name string
		args args
		want contract.OffsetPaginator
	}{
		{
			name: "NewGormOffsetPaginator Success",
			args: args{maxPerPage},
			want: &gormOffsetPaginator{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormOffsetPaginator(tt.args.maxPerPage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGormOffsetPaginator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_SetPage(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	type args struct {
		page uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "gormOffsetPaginator_SetPage Success",
			args: args{maxPerPage},
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			p.SetPage(tt.args.page)
		})
	}
}

func Test_gormOffsetPaginator_Page(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "gormOffsetPaginator_Page Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.Page(); got != tt.want {
				t.Errorf("gormOffsetPaginator.Page() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_Results(t *testing.T) {
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	countRows := mock.NewRows([]string{"count"}).
		AddRow(1)
	mock.ExpectQuery("SELECT").WillReturnRows(countRows)

	rows := mock.NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
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
			name: "gormOffsetPaginator_Results Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
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
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if err := p.Results(tt.args.ctx, tt.args.table, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("gormOffsetPaginator.Results() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormOffsetPaginator_Nums(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "gormCursorPaginator_Nums Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.Nums(); got != tt.want {
				t.Errorf("gormOffsetPaginator.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_Limit(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "gormOffsetPaginator_Limit Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.Limit(); got != tt.want {
				t.Errorf("gormOffsetPaginator.Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_HasPages(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "gormOffsetPaginator_HasPages Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.HasPages(); got != tt.want {
				t.Errorf("gormOffsetPaginator.HasPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_HasNext(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "gormOffsetPaginator_HasNext Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.HasNext(); got != tt.want {
				t.Errorf("gormOffsetPaginator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_PrevPage(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "gormOffsetPaginator_PrevPage Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.PrevPage(); got != tt.want {
				t.Errorf("gormOffsetPaginator.PrevPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_NextPage(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "gormOffsetPaginator_NextPage Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.NextPage(); got != tt.want {
				t.Errorf("gormOffsetPaginator.NextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_HasPrev(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "gormOffsetPaginator_HasPrev Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.HasPrev(); got != tt.want {
				t.Errorf("gormOffsetPaginator.HasPrev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_PageNums(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "gormOffsetPaginator_PageNums Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			if got := p.PageNums(); got != tt.want {
				t.Errorf("gormOffsetPaginator.PageNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormOffsetPaginator_SetLimit(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
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
			name: "gormOffsetPaginator_SetLimit Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
				order:      "DESC",
			},
			args: args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			p.SetLimit(tt.args.limit)
		})
	}
}

func Test_gormOffsetPaginator_SetOrderBy(t *testing.T) {
	type fields struct {
		maxPerPage uint64
		page       uint64
		nums       uint64
		orderBy    string
		order      string
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
			name: "gormOffsetPaginator_SetOrderBy Success",
			fields: fields{
				maxPerPage: maxPerPage,
				page:       1,
				nums:       0,
				orderBy:    "created_at",
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
			p := &gormOffsetPaginator{
				maxPerPage: tt.fields.maxPerPage,
				page:       tt.fields.page,
				nums:       tt.fields.nums,
				orderBy:    tt.fields.orderBy,
				order:      tt.fields.order,
			}
			p.SetOrderBy(tt.args.key, tt.args.order)
		})
	}
}
