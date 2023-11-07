package postgresql

import (
	"context"
	"fmt"
	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBConnection(conf config.ServerConfig) (*pgxpool.Pool, error) {
	var (
		user           = "postgres"
		password       = "postgres"
		host           = "localhost"
		port           = "5432"
		dbname         = "Tabula"
		appName        = "Tabula"
		schema         = "public"
		connectTimeout = 5
	)

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?application_name=%s&search_path=%s&connect_timeout=%d",
		user, password, host, port, dbname, appName, schema, connectTimeout,
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
