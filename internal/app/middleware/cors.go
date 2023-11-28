package middleware

import (
	"net/http"
	"server/internal/config"

	logging "server/internal/logging"

	"github.com/rs/cors"
)

func GetCors(cc config.CORSConfig, logger logging.ILogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return cors.New(cors.Options{
			AllowedHeaders:   cc.AllowedHeaders,
			ExposedHeaders:   cc.ExposedHeaders,
			AllowedOrigins:   cc.AllowedHosts,
			AllowCredentials: cc.AllowCredentials,
			AllowedMethods:   cc.AllowedMethods,
			Debug:            logger.Level() <= 1,
			Logger:           logger,
		}).Handler(next)
	}
}
