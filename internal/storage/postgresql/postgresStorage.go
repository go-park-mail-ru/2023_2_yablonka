package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const nodeName string = "storage"

func GetDBConnection(conf config.DatabaseConfig) (*sql.DB, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?application_name=%s&search_path=%s&connect_timeout=%d",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
		conf.AppName,
		conf.Schema,
		conf.ConnectionTimeout,
	)

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDBConnectionOld(conf config.DatabaseConfig) (*pgxpool.Pool, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?application_name=%s&search_path=%s&connect_timeout=%d",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
		conf.AppName,
		conf.Schema,
		conf.ConnectionTimeout,
	)

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
