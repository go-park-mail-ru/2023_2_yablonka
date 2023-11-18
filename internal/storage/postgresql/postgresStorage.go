package postgresql

import (
	"context"
	"fmt"
	"server/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func GetDBConnection(conf config.DatabaseConfig) (*pgxpool.Pool, error) {
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

func storageDebugLog(logger *logrus.Logger, function string, message string) {
	logger.
		WithFields(logrus.Fields{
			"route_node": "storage",
			"function":   function,
		}).
		Debug(message)
}

// func storageWarnLog(logger *logrus.Logger, function string, message string) {
// 	logger.
// 		WithFields(logrus.Fields{
// 			"route_node": "storage",
// 			"function":   function,
// 		}).
// 		Warn(message)
// }
