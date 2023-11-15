package middleware

import (
	"net/http"
	"server/internal/config"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

func GetCors(cc config.CORSConfig, lc config.LoggingConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return cors.New(cors.Options{
			AllowedHeaders:   cc.AllowedHeaders,
			ExposedHeaders:   cc.ExposedHeaders,
			AllowedOrigins:   cc.AllowedHosts,
			AllowCredentials: cc.AllowCredentials,
			AllowedMethods:   cc.AllowedMethods,
			Debug:            lc.Logger.Level >= logrus.DebugLevel,
			Logger:           lc.Logger,
		}).Handler(next)
	}
}
