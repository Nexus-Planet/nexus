package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	Postgres       = "postgres"
	Mysql          = "mysql"
	Sqlite         = "sqlite"
	PostgresDriver = "pgx"
	SqliteDriver   = "sqlite3"
	MysqlDriver    = "mysql"
)

// Connects to the database
func Connect(dbName string, dsn string) (*sqlx.DB, error) {

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

// Connects to the database with context
func ConnectContext(ctx context.Context, dbName string, dsn string) (*sqlx.DB, error) {

	switch dbName {
	case Postgres, PostgresDriver:
		db, err := sqlx.ConnectContext(ctx, PostgresDriver, dsn)
		if err != nil {
			return nil, err
		}
		return db, nil
	case Mysql:
		db, err := sqlx.ConnectContext(ctx, MysqlDriver, dsn)
		if err != nil {
			return nil, err
		}
		return db, nil
	case Sqlite, SqliteDriver:
		db, err := sqlx.ConnectContext(ctx, SqliteDriver, dsn)
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		err := fmt.Errorf("ERROR:Unsupported database")
		return nil, err
	}
}

func ToNullString(s *string) *sql.NullString {
	if s == nil {
		return &sql.NullString{
			Valid: false,
		}
	}
	return &sql.NullString{
		String: *s,
		Valid:  true,
	}
}
