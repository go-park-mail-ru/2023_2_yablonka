package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

const configPath string = "../../internal/config/config.yml"

func Cors(next http.Handler) http.Handler {
	config, _ := NewConfig(configPath)
	return cors.New(cors.Options{
		AllowedHeaders:   config.Server.AllowedHeaders,
		AllowedOrigins:   config.Server.AllowedHosts,
		AllowCredentials: true,
		AllowedMethods:   config.Server.AllowedMethods,
		Debug:            true,
	}).Handler(next)
}
