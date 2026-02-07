package db

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
)

var (
	Postgres       = "postgres"
	Mysql          = "mysql"
	Sqlite         = "sqlite"
	PostgresDriver = "pgx"
	SqliteDriver   = "sqlite3"
	MysqlDriver    = "mysql"
)

// Connects to the database
func Connect(ctx context.Context, dbName string) (*sqlx.DB, error) {
	var dsn string

	if config.CustomDataSourceName == "" {
		dsn = config.DataSourceName
	} else {
		dsn = config.CustomDataSourceName
	}

	switch dbName {
	case Postgres, PostgresDriver:
		db, err := sqlx.Connect(PostgresDriver, dsn)
		if err != nil {
			return nil, err
		}
		return db, nil
	case Mysql:
		db, err := sqlx.Connect(MysqlDriver, dsn)
		if err != nil {
			return nil, err
		}
		return db, nil
	case Sqlite, SqliteDriver:
		db, err := sqlx.Connect(SqliteDriver, dsn)
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		err := fmt.Errorf("ERROR:Unsupported database")
		return nil, err
	}
}
