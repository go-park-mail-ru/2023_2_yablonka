package postgresql

import (
	"database/sql"
	"fmt"
	"server/internal/config"
)

func getPostgres(conf config.ServerConfig) (*sql.DB, error) {
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
