package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"github.com/KurobaneShin/eulabs/types"
)

type DBInterface interface {
	CreateProduct(product *types.Product) error
	UpdateProduct(product *types.Product) error
	GetProductById(id string) (types.Product, error)
	DeleteProduct(id string) error
}

type DB struct {
	*bun.DB
}

func Create() DB {
	{
		var (
			host     = os.Getenv("DB_HOST")
			user     = os.Getenv("DB_USER")
			password = os.Getenv("DB_PASSWORD")
			name     = os.Getenv("DB_NAME")
		)

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
			user,
			password,
			host,
			name)

		sqldb, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		return DB{bun.NewDB(sqldb, mysqldialect.New())}
	}
}
