package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBConnection(conf config.ServerConfig) (*pgxpool.Pool, error) {
	var (
		user     = "postgres"
		password = "postgres"
		host     = "localhost"
		port     = "5432"
		dbname   = "Tabula"
	)

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname,
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

func NewPostgresStorageOld(conf config.ServerConfig) (*sql.DB, error) {
	var (
		user     = "postgres"
		dbname   = "Tabula"
		password = "postgres"
		host     = "127.0.0.1"
		port     = "8080"
		sslmode  = "disable"
	)

	dsn := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
		user, dbname, password, host, port, sslmode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return db, nil
}
