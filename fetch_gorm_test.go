package main

import (
	"context"
	"testing"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	drathGorm "github.com/007lock/drath/gorm"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

type ModelBasis struct {
	ID uint64
}

func createGormSession() (*gorm.DB, sqlmock.Sqlmock, error) {
	// Init mock db; mock match extract query not using regex
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

func TestFetchCriterias(t *testing.T) {
	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT * FROM "table" WHERE (field = $1) AND (and_field = $2) AND (and_field = $3 OR and_or_field = $4) AND (field = $5 OR or_field = $6) ORDER BY random(),"table"."id" ASC LIMIT 1`).
		WithArgs(1, 2, 3, 4, 5, 6).
		WillReturnRows(mock.NewRows([]string{"id"}).
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
		args    args
		wantErr bool
	}{
		{
			name: "NewGormFetchRepository",
			args: args{
				c:     c,
				table: "table",
				item:  new(ModelBasis),
				crit: &contract.RepoCriterias{
					Conditions: []*contract.RepoCondition{
						{
							Field:     "field",
							Operation: "=",
							Value:     1,
						},
						{
							Field:     "and_field",
							Operation: "=",
							Value:     2,
							OrConditions: []*contract.RepoCondition{
								{
									Field:     "and_field",
									Operation: "=",
									Value:     3,
								},
								{
									Field:     "and_or_field",
									Operation: "=",
									Value:     4,
								},
							},
						},
					},
					OrConditions: []*contract.RepoCondition{
						{
							Field:     "field",
							Operation: "=",
							Value:     5,
						},
						{
							Field:     "or_field",
							Operation: "=",
							Value:     6,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := drathGorm.NewGormFetchRepository()
			if err := r.GetByRandom(tt.args.c, tt.args.table, tt.args.item, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.GetByRandom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
