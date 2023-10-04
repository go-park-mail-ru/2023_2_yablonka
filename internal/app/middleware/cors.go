package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/rs/cors"
)

const configPath string = "internal/config/config.yml"

func Cors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedHeaders: []string{
			//"*",
			"Acccess-Control-Allow-Origin",
			"Content-Type",
		},
		AllowedOrigins: []string{
			//"*",
			"localhost:8080",
			"213.219.215.40:8080",
			"localhost:8081",
			"213.219.215.40:8081",
			"http://localhost:8081",
			"http://213.219.215.40:8081",
		},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		Debug:            true,
	}).Handler(next)
}

func CorsNew(h http.Handler) http.Handler {
	config, _ := NewConfig(configPath)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		w.Header().Add("Access-Control-Allow-Origin", strings.Join(config.Server.AllowedHosts, ", "))
		w.Header().Add("Access-Control-Allow-Headers", strings.Join(config.Server.AllowHeaders, ", "))
=======
		w.Header().Add("Access-Control-Allow-Origin", strings.Join(config.Server.AllowedHeaders, ", "))
		w.Header().Add("Access-Control-Allow-Headers", strings.Join(config.Server.AllowedHosts, ", "))
>>>>>>> e4fa8ead9a7add7e7e2679e377abd7628f94809f
		w.Header().Add("Access-Control-Allow-Methods", strings.Join(config.Server.AllowedMethods, ", "))
		w.Header().Add("Access-Control-Allow-Credentials", strconv.FormatBool(config.Server.AllowCredentials))
		h.ServeHTTP(w, r)
	})
}
