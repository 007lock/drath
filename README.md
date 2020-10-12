# go drath repository

The Repository pattern: a painless way to simplify your Go service logic

## Initial connection

Currently we are support gorm only for easy integration with this repository pattern, In your custom repository you could use whatever driver as you want to integration this pattern.

Step 1: Init gorm return Database interface

```go
import (
	drathContract "github.com/007lock/drath/contract"
	"github.com/007lock/krarks/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jinzhu/gorm"

	// Postgres migrate driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// File migrate driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type gormSession struct {
	cfg     *config.Config
	session *gorm.DB
	m       *migrate.Migrate
}

func InitGorm() (drathContract.Database, error) {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres password=postgres dbname=postgres")

	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.DB().SetMaxOpenConns(5)
	db.DB().SetMaxIdleConns(0)

    // Optional if you don't like using migration please ignore codes from here
	m, err := migrate.New(
		"file://./migrations",
		"postgres://username:password@host:port/dbname")
	if err != nil {
		return nil, err
	}

	return &gormSession{
		cfg:     conf,
		session: db,
		m:       m,
	}, nil
}

func (db *gormSession) Get() (interface{}, error) {
	return db.session, nil
}

func (db *gormSession) Begin() (interface{}, error) {
	return db.session.Begin(), nil
}

func (db *gormSession) Close() error {
	return db.session.Close()
}

func (db *gormSession) MigrationUp() error {
	return db.m.Up()
}

func (db *gormSession) MigrationDown() error {
	return db.m.Down()
}
```

Step 2: Using DB interface at your code

```go
import (
    ...
	drathGorm "github.com/007lock/drath/gorm"
	"github.com/jinzhu/gorm"
    drathContants "github.com/007lock/drath/constants"
    drathContract "github.com/007lock/drath/contract"
)
// DB session
dbSession, err := db.InitGorm(cfg)
if err != nil {
    fmt.Fatal(err)
}
    
// Transaction using gorm
txi, err := dbSession.Begin().(*gorm.DB)
if err != nil {
    fmt.Fatal(err)
}

defer func() {
    if err != nil {
        txi.Rollback()
        return
    }

    txi.Commit()
}() 

// None transaction
// txi, err := dbSession.get()
// if err != nil {
//     fmt.Fatal(err)
// }

c := context.WithValue(context.Background(), constants.ContextKeyTransaction, txi)

var FetchRepository drathContract.FetchRepository = drathGorm.NewGormFetchRepository()
type Product struct {
    ID             int64             `json:"id"`
    Title          string            `json:"title"`
    Status         uint8             `json:"status"`
}

randomProduct := new(Product)

// Filter follow criteria
crit = &drathContract.RepoCriterias{
    Conditions: []*drathContract.RepoCondition{
        {
            Field:     "status",
            Operation: "=",
            Value:     "1", // search active product
        },
    },
}
err = s.FetchRepository.GetByRandom(c, "products", randomProduct, crit)
if err != nil {
    if err == drathContants.DBError.ERROR_RECORD_NOT_FOUND {
        fmt.Println("Product not found")
        return 
    }
    fmt.Fatal(err)
}

fmt.Println(randomProduct)
```