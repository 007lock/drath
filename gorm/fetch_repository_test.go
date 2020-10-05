package gorm

import (
	"context"
	"reflect"
	"testing"

	"github.com/007lock/drath/contract"
	"github.com/jinzhu/gorm"
)

func TestNewGormFetchRepository(t *testing.T) {
	tests := []struct {
		name string
		want contract.FetchRepository
	}{
		{
			name: "TestNewGormFetchRepository",
			want: &gormFetchRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormFetchRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGormFetchRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

type ModelBasis struct {
	ID uint64
}

func Test_gormFetchRepository_GetByRandom(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}
	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"id"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		item  interface{}
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormFetchRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormFetchRepository_GetByRandom",
			r:    &gormFetchRepository{},
			args: args{
				c:     c,
				item:  new(ModelBasis),
				crit:  nil,
				table: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			if err := r.GetByRandom(tt.args.c, tt.args.table, tt.args.item, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.GetByRandom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormFetchRepository_FetchByRandom(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"id"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		item  interface{}
		crit  *contract.RepoCriterias
		limit uint64
	}
	tests := []struct {
		name    string
		r       *gormFetchRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormFetchRepository_FetchByRandom",
			r:    &gormFetchRepository{},
			args: args{
				c:     c,
				item:  new(ModelBasis),
				limit: 1,
				table: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			if err := r.FetchByRandom(tt.args.c, tt.args.table, tt.args.item, tt.args.crit, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.FetchByRandom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormFetchRepository_FetchCursor(t *testing.T) {
	// Init mock db
	db, _, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		item  interface{}
		p     contract.CursorPaginator
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormFetchRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormFetchRepository_FetchCursor",
			r:    &gormFetchRepository{},
			args: args{
				c:    c,
				item: nil,
				p: &contract.CursorPaginatorMock{
					ResultsFunc: func(ctx context.Context, table string, data interface{}) error {
						return nil
					},
				},
				table: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			if err := r.FetchCursor(tt.args.c, tt.args.table, tt.args.item, tt.args.p, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.FetchCursor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormFetchRepository_FetchPagination(t *testing.T) {
	// Init mock db
	db, _, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		item  interface{}
		p     contract.OffsetPaginator
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormFetchRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormFetchRepository_FetchPagination_error",
			r:    &gormFetchRepository{},
			args: args{
				c:    c,
				item: nil,
				p: &contract.OffsetPaginatorMock{
					ResultsFunc: func(ctx context.Context, table string, data interface{}) error {
						return gorm.ErrRecordNotFound
					},
				},
				table: "1",
			},
			wantErr: true,
		},
		{
			name: "gormFetchRepository_FetchPagination_success",
			r:    &gormFetchRepository{},
			args: args{
				c:    c,
				item: nil,
				p: &contract.OffsetPaginatorMock{
					ResultsFunc: func(ctx context.Context, table string, data interface{}) error {
						return nil
					},
				},
				table: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			if err := r.FetchPagination(tt.args.c, tt.args.table, tt.args.item, tt.args.p, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.FetchPagination() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormFetchRepository_GetByID(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"id"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		id    string
		item  interface{}
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormFetchRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormFetchRepository_GetByID",
			r:    &gormFetchRepository{},
			args: args{
				c:     c,
				id:    "1",
				item:  new(ModelBasis),
				table: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			if err := r.GetByID(tt.args.c, tt.args.table, tt.args.id, tt.args.item, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormFetchRepository_GetByCriteria(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"id"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		item  interface{}
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormFetchRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormFetchRepository_GetByCriteria",
			r:    &gormFetchRepository{},
			args: args{
				c:     c,
				item:  new(ModelBasis),
				table: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			if err := r.GetByCriteria(tt.args.c, tt.args.table, tt.args.item, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.GetByCriteria() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormFetchRepository_FetchByCriteria(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"id"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		item  interface{}
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormFetchRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormFetchRepository_FetchByCriteria",
			r:    &gormFetchRepository{},
			args: args{
				c:     c,
				item:  new(ModelBasis),
				table: "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			if err := r.FetchByCriteria(tt.args.c, tt.args.table, tt.args.item, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.FetchByCriteria() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormFetchRepository_ApplyCriteria(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c    context.Context
		crit *contract.RepoCriterias
	}
	tests := []struct {
		name string
		r    *gormFetchRepository
		args args
	}{
		{
			name: "gormFetchRepository_ApplyCriteria",
			r:    &gormFetchRepository{},
			args: args{
				c: c,
				crit: &contract.RepoCriterias{
					Conditions: []*contract.RepoCondition{
						&contract.RepoCondition{
							Field:     "field",
							Operation: "=",
							Value:     "value",
						},
					},
					OrConditions: []*contract.RepoCondition{
						&contract.RepoCondition{
							Field:     "field",
							Operation: "=",
							Value:     "value",
						},
						&contract.RepoCondition{
							Field:     "field",
							Operation: "=",
							Value:     "value",
						},
					},
					Joins: []*contract.RepoJoin{
						&contract.RepoJoin{
							Join: "join",
						},
					},
					GroupBy: []string{"GroupBy"},
					Preloads: []*contract.Preload{
						&contract.Preload{
							Relation: "Preload",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormFetchRepository{}
			r.ApplyCriteria(tt.args.c, tt.args.crit)
		})
	}
}
