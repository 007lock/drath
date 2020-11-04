package main

import (
	"context"
	"testing"
	"time"

	"github.com/007lock/drath/constants"
	"github.com/007lock/drath/contract"
	drathGorm "github.com/007lock/drath/gorm"
	"github.com/007lock/drath/paginator"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

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

func TestFetchRandomCriterias(t *testing.T) {
	type ModelBasis struct {
		ID uint64
	}

	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	// fetch_random_repository_condition_test
	mock.ExpectQuery(`SELECT * FROM "table" WHERE (field = $1) AND (and_field = $2) AND (and_field = $3 OR and_or_field = $4) AND (field = $5 OR or_field = $6) ORDER BY random(),"table"."id" ASC LIMIT 1`).
		WithArgs(1, 2, 3, 4, 5, 6).
		WillReturnRows(mock.NewRows([]string{"id"}).
			AddRow(1))
	// fetch_random_repository_join_test
	mock.ExpectQuery(`SELECT "table".* FROM "table" join relation on relation.table_id=table.id join relation2 on relation2.table_id=table.id WHERE (field = $1) AND (relation.field = $2) AND (relation2.field = $3) ORDER BY random(),"table"."id" ASC LIMIT 1`).
		WithArgs(1, 2, 3).
		WillReturnRows(mock.NewRows([]string{"id"}).
			AddRow(1))
	// fetch_random_repository_group_by_test
	mock.ExpectQuery(`SELECT "table".* FROM "table" join relation on relation.table_id=table.id join relation2 on relation2.table_id=table.id WHERE (field = $1) AND (relation.field = $2) AND (relation2.field = $3) GROUP BY relation.id ORDER BY random(),"table"."id" ASC LIMIT 1`).
		WithArgs(1, 2, 3).
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
			name: "fetch_random_repository_condition_test",
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
		{
			name: "fetch_random_repository_join_test",
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
							Field:     "relation.field",
							Operation: "=",
							Value:     2,
						},
						{
							Field:     "relation2.field",
							Operation: "=",
							Value:     3,
						},
					},
					Joins: []*contract.RepoJoin{
						&contract.RepoJoin{
							Join: "join relation on relation.table_id=table.id",
						},
						&contract.RepoJoin{
							Join: "join relation2 on relation2.table_id=table.id",
						},
					},
				},
			},
		},
		{
			name: "fetch_random_repository_group_by_test",
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
							Field:     "relation.field",
							Operation: "=",
							Value:     2,
						},
						{
							Field:     "relation2.field",
							Operation: "=",
							Value:     3,
						},
					},
					Joins: []*contract.RepoJoin{
						&contract.RepoJoin{
							Join: "join relation on relation.table_id=table.id",
						},
						&contract.RepoJoin{
							Join: "join relation2 on relation2.table_id=table.id",
						},
					},
					GroupBy: []string{"relation.id"},
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

func TestFetchCursorPaginationCriterias(t *testing.T) {
	type ModelBasis struct {
		ID        uint64
		CreatedAt time.Time
	}

	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	// fetch_repository_pagination_cursor_test
	mock.ExpectQuery(`SELECT "table".* FROM "table" join relation on relation.table_id=table.id join relation2 on relation2.table_id=table.id WHERE (field = $1) AND (and_field = $2) AND (and_field = $3 OR and_or_field = $4) AND (relation.field = $5) AND (relation2.field = $6) AND (field = $7 OR or_field = $8) GROUP BY relation.id ORDER BY created_at DESC LIMIT 6`).
		WithArgs(1, 2, 3, 4, 5, 6, 7, 8).
		WillReturnRows(mock.NewRows([]string{"id"}).
			AddRow(1))
	//fetch_repository_pagination_cursor_next_page_test
	cutOffTime, err := time.Parse(
		"2006-01-02 15:04:05 -0700 MST",
		"2020-10-28 08:37:12 +0000 UTC")
	if err != nil {
		t.Fatal(err)
	}
	mock.ExpectQuery(`SELECT "table".* FROM "table" join relation on relation.table_id=table.id join relation2 on relation2.table_id=table.id WHERE (field = $1) AND (and_field = $2) AND (and_field = $3 OR and_or_field = $4) AND (relation.field = $5) AND (relation2.field = $6) AND (field = $7 OR or_field = $8) AND (created_at < $9) GROUP BY relation.id ORDER BY created_at DESC LIMIT 6`).
		WithArgs(1, 2, 3, 4, 5, 6, 7, 8, cutOffTime).
		WillReturnRows(mock.NewRows([]string{"id", "created_at"}).
			AddRow(1, cutOffTime))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c          context.Context
		table      string
		items      interface{}
		nextCursor string
		orderBy    string
		paginator  contract.CursorPaginator
		crit       *contract.RepoCriterias
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "fetch_repository_pagination_cursor_first_page_test",
			args: args{
				c:          c,
				table:      "table",
				items:      &[]*ModelBasis{},
				nextCursor: "",
				orderBy:    "CreatedAt",
				paginator:  paginator.NewGormCursorPaginator(5),
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
						{
							Field:     "relation.field",
							Operation: "=",
							Value:     5,
						},
						{
							Field:     "relation2.field",
							Operation: "=",
							Value:     6,
						},
					},
					OrConditions: []*contract.RepoCondition{
						{
							Field:     "field",
							Operation: "=",
							Value:     7,
						},
						{
							Field:     "or_field",
							Operation: "=",
							Value:     8,
						},
					},
					Joins: []*contract.RepoJoin{
						&contract.RepoJoin{
							Join: "join relation on relation.table_id=table.id",
						},
						&contract.RepoJoin{
							Join: "join relation2 on relation2.table_id=table.id",
						},
					},
					GroupBy: []string{"relation.id"},
				},
			},
		},
		{
			name: "fetch_repository_pagination_cursor_next_page_test",
			args: args{
				c:          c,
				table:      "table",
				items:      &[]*ModelBasis{},
				orderBy:    "CreatedAt",
				nextCursor: "MjAyMC0xMC0yOFQwODozNzoxMlo/VElNRQ==", //2020-10-28 08:37:12 +0000 UTC
				paginator:  paginator.NewGormCursorPaginator(5),
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
						{
							Field:     "relation.field",
							Operation: "=",
							Value:     5,
						},
						{
							Field:     "relation2.field",
							Operation: "=",
							Value:     6,
						},
					},
					OrConditions: []*contract.RepoCondition{
						{
							Field:     "field",
							Operation: "=",
							Value:     7,
						},
						{
							Field:     "or_field",
							Operation: "=",
							Value:     8,
						},
					},
					Joins: []*contract.RepoJoin{
						&contract.RepoJoin{
							Join: "join relation on relation.table_id=table.id",
						},
						&contract.RepoJoin{
							Join: "join relation2 on relation2.table_id=table.id",
						},
					},
					GroupBy: []string{"relation.id"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := drathGorm.NewGormFetchRepository()
			tt.args.paginator.SetAfterCursor(tt.args.nextCursor)
			tt.args.paginator.SetOrderBy(tt.args.orderBy, "DESC")
			if err := r.FetchCursor(tt.args.c, tt.args.table, tt.args.items, tt.args.paginator, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.FetchCursor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFetchOffetPaginationCriterias(t *testing.T) {
	type ModelBasis struct {
		ID        uint64
		CreatedAt time.Time
	}

	// Init mock db
	db, mock, err := createGormSession()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectBegin()
	// fetch_repository_pagination_offset_test
	mock.ExpectQuery(`SELECT count(*) FROM ( SELECT count(*) as name FROM "table" join relation on relation.table_id=table.id join relation2 on relation2.table_id=table.id WHERE (field = $1) AND (and_field = $2) AND (and_field = $3 OR and_or_field = $4) AND (relation.field = $5) AND (relation2.field = $6) AND (field = $7 OR or_field = $8) GROUP BY relation.id ) AS count_table`)
	mock.ExpectQuery(`SELECT "table".* FROM "table" join relation on relation.table_id=table.id join relation2 on relation2.table_id=table.id WHERE (field = $1) AND (and_field = $2) AND (and_field = $3 OR and_or_field = $4) AND (relation.field = $5) AND (relation2.field = $6) AND (field = $7 OR or_field = $8) GROUP BY relation.id ORDER BY created_at DESC LIMIT 5 OFFSET 0`).
		WithArgs(1, 2, 3, 4, 5, 6, 7, 8).
		WillReturnRows(mock.NewRows([]string{"id"}).
			AddRow(1))

	c, err := createContext(db)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		c         context.Context
		table     string
		items     interface{}
		page      uint64
		paginator contract.OffsetPaginator
		crit      *contract.RepoCriterias
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "fetch_repository_pagination_offset_test",
			args: args{
				c:         c,
				table:     "table",
				items:     &[]*ModelBasis{},
				page:      1,
				paginator: paginator.NewGormOffsetPaginator(5),
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
						{
							Field:     "relation.field",
							Operation: "=",
							Value:     5,
						},
						{
							Field:     "relation2.field",
							Operation: "=",
							Value:     6,
						},
					},
					OrConditions: []*contract.RepoCondition{
						{
							Field:     "field",
							Operation: "=",
							Value:     7,
						},
						{
							Field:     "or_field",
							Operation: "=",
							Value:     8,
						},
					},
					Joins: []*contract.RepoJoin{
						&contract.RepoJoin{
							Join: "join relation on relation.table_id=table.id",
						},
						&contract.RepoJoin{
							Join: "join relation2 on relation2.table_id=table.id",
						},
					},
					GroupBy: []string{"relation.id"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := drathGorm.NewGormFetchRepository()
			tt.args.paginator.SetPage(tt.args.page)
			if err := r.FetchPagination(tt.args.c, tt.args.table, tt.args.items, tt.args.paginator, tt.args.crit); (err != nil) != tt.wantErr {
				t.Errorf("gormFetchRepository.FetchCursor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
