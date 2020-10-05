package gorm

import (
	"context"
	"reflect"
	"testing"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestNewGormAggregationRepository(t *testing.T) {
	tests := []struct {
		name string
		want contract.AggregationRepository
	}{
		{
			name: "TestNewGormAggregationRepository",
			want: &gormAggRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormAggregationRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGormAggregationRepository() = %v, want %v", got, tt.want)
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

func createContext(db *gorm.DB) (context.Context, error) {
	// Echo
	ctx := context.Background()

	if db != nil {
		tx := db.Begin()
		ctx = context.WithValue(context.Background(), constants.ContextKeyTransaction, tx)
	}

	return ctx, nil
}

func Test_gormAggRepository_CountUnique(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}
	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"total"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		field string
		name  string
	}
	tests := []struct {
		name    string
		r       *gormAggRepository
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "gormAggRepository_CountUnique",
			r:    &gormAggRepository{},
			args: args{
				c:     c,
				table: "table",
				field: "field",
				name:  "name",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormAggRepository{}
			got, err := r.CountUnique(tt.args.c, tt.args.table, tt.args.field, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("gormAggRepository.CountUnique() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gormAggRepository.CountUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormAggRepository_CountByCriteria(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}
	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"total"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormAggRepository
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "gormAggRepository_CountByCriteria",
			r:    &gormAggRepository{},
			args: args{
				c:     c,
				table: "table",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormAggRepository{}
			got, err := r.CountByCriteria(tt.args.c, tt.args.table, tt.args.crit)
			if (err != nil) != tt.wantErr {
				t.Errorf("gormAggRepository.CountByCriteria() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gormAggRepository.CountByCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormAggRepository_SumByCriteria(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}
	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(mock.NewRows([]string{"n"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		field string
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormAggRepository
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "gormAggRepository_SumByCriteria",
			r:    &gormAggRepository{},
			args: args{
				c:     c,
				table: "table",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormAggRepository{}
			got, err := r.SumByCriteria(tt.args.c, tt.args.table, tt.args.field, tt.args.crit)
			if (err != nil) != tt.wantErr {
				t.Errorf("gormAggRepository.SumByCriteria() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gormAggRepository.SumByCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}
