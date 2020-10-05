package gorm

import (
	"context"
	"reflect"
	"testing"

	"github.com/007lock/drath/contract"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewGormUpdateRepository(t *testing.T) {
	tests := []struct {
		name string
		want contract.UpdateRepository
	}{
		{
			name: "TestNewGormUpdateRepository",
			want: &gormUpdateRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormUpdateRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGormUpdateRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gormUpdateRepository_Update(t *testing.T) {
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
		r       *gormUpdateRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormUpdateRepository_Update",
			r:    &gormUpdateRepository{},
			args: args{
				c:    c,
				item: new(ModelBasis),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormUpdateRepository{}
			if err := r.Update(tt.args.c, tt.args.table, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("gormUpdateRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormUpdateRepository_UpdateByCriteria(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

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
		crit  *contract.RepoCriterias
	}
	tests := []struct {
		name    string
		r       *gormUpdateRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormUpdateRepository_UpdateByCriteria",
			r:    &gormUpdateRepository{},
			args: args{
				c:     c,
				item:  new(ModelBasis),
				table: "relationship",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormUpdateRepository{}
			if err := r.UpdateByCriteria(tt.args.c, tt.args.table, tt.args.item, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormUpdateRepository.UpdateByCriteria() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormUpdateRepository_Delete(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	// Mock DB
	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c         context.Context
		table     string
		item      interface{}
		relations []string
	}
	tests := []struct {
		name    string
		r       *gormUpdateRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormUpdateRepository_Delete",
			r:    &gormUpdateRepository{},
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormUpdateRepository{}
			if err := r.Delete(tt.args.c, tt.args.table, tt.args.item, tt.args.relations...); (err != nil) != tt.wantErr {
				t.Errorf("gormUpdateRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormUpdateRepository_DeleteByCriteria(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c         context.Context
		table     string
		item      interface{}
		crit      *contract.RepoCriterias
		relations []string
	}
	tests := []struct {
		name    string
		r       *gormUpdateRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormUpdateRepository_DeleteByCriteria",
			r:    &gormUpdateRepository{},
			args: args{
				c:     c,
				item:  new(ModelBasis),
				table: "relationship",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormUpdateRepository{}
			if err := r.DeleteByCriteria(tt.args.c, tt.args.table, tt.args.item, tt.args.crit, tt.args.relations...); (err != nil) != tt.wantErr {
				t.Errorf("gormUpdateRepository.DeleteByCriteria() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormUpdateRepository_Attach(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c          context.Context
		associated interface{}
		relation   string
		entities   interface{}
	}
	tests := []struct {
		name    string
		r       *gormUpdateRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormUpdateRepository_Attach",
			r:    &gormUpdateRepository{},
			args: args{
				c:          c,
				associated: new(ModelBasis),
				relation:   "relationship",
				entities:   new(ModelBasis),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormUpdateRepository{}
			if err := r.Attach(tt.args.c, tt.args.associated, tt.args.relation, tt.args.entities); (err != nil) != tt.wantErr {
				t.Errorf("gormUpdateRepository.Attach() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormUpdateRepository_Dettach(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c          context.Context
		associated interface{}
		relation   string
		entities   interface{}
	}
	tests := []struct {
		name    string
		r       *gormUpdateRepository
		args    args
		wantErr bool
	}{
		{
			name: "GORM Dettach Success",
			r:    &gormUpdateRepository{},
			args: args{
				c: c,
				associated: &ModelBasis{
					ID: 1,
				},
				relation: "relationship",
				entities: new(ModelBasis),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormUpdateRepository{}
			if err := r.Dettach(tt.args.c, tt.args.associated, tt.args.relation, tt.args.entities); (err != nil) != tt.wantErr {
				t.Errorf("gormUpdateRepository.Dettach() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gormUpdateRepository_DettachAll(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c          context.Context
		associated interface{}
		relation   string
	}
	tests := []struct {
		name    string
		r       *gormUpdateRepository
		args    args
		wantErr bool
	}{
		{
			name: "gormUpdateRepository_DettachAll",
			r:    &gormUpdateRepository{},
			args: args{
				c: c,
				associated: &ModelBasis{
					ID: 1,
				},
				relation: "relationship",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &gormUpdateRepository{}
			if err := r.DettachAll(tt.args.c, tt.args.associated, tt.args.relation); (err != nil) != tt.wantErr {
				t.Errorf("gormUpdateRepository.DettachAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
