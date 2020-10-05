package gorm

import (
	"context"
	"reflect"
	"testing"

	"github.com/007lock/drath/contract"
)

func TestNewGormInsertRepository(t *testing.T) {
	tests := []struct {
		name string
		want contract.InsertRepository
	}{
		{
			name: "TestNewGormInsertRepository",
			want: &gormInsertRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormInsertRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGormInsertRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormInsertRepository_Store(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(mock.NewRows([]string{"id"}).
		AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c     context.Context
		table string
		item  interface{}
	}
	tests := []struct {
		name    string
		r       *gormInsertRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormInsertRepository_Store",
			r:    &gormInsertRepository{},
			args: args{
				c:    c,
				item: new(ModelBasis),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormInsertRepository{}
			if err := r.Store(tt.args.c, tt.args.table, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("gormInsertRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormInsertRepository_StoreOrUpdate(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(mock.NewRows([]string{"id"}).
		AddRow(1))
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
	}
	tests := []struct {
		name    string
		r       *gormInsertRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormInsertRepository_StoreOrUpdate",
			r:    &gormInsertRepository{},
			args: args{
				c:    c,
				item: new(ModelBasis),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormInsertRepository{}
			if err := r.StoreOrUpdate(tt.args.c, tt.args.table, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("gormInsertRepository.StoreOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
