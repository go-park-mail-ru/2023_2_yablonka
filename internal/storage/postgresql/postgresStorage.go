package postgresql

import (
	"context"
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
