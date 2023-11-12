package middleware

import (
	"net/http"
	"server/internal/config"

	"github.com/rs/cors"
)

func GetCors(conf config.CORSConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return cors.New(cors.Options{
			AllowedHeaders:   conf.AllowedHeaders,
			ExposedHeaders:   conf.ExposedHeaders,
			AllowedOrigins:   conf.AllowedHosts,
			AllowCredentials: conf.AllowCredentials,
			AllowedMethods:   conf.AllowedMethods,
			Debug:            conf.Debug,
		}).Handler(next)
	}
}
