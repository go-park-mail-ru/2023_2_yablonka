package middleware

import (
	"net/http"
	"server/internal/config"

	"github.com/rs/cors"
)

func GetCors(conf config.BaseServerConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return cors.New(cors.Options{
			AllowedHeaders:   conf.Server.AllowedHeaders,
			AllowedOrigins:   conf.Server.AllowedHosts,
			AllowCredentials: true,
			AllowedMethods:   conf.Server.AllowedMethods,
			Debug:            true,
		}).Handler(next)
	}
}
