package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
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
	})
	return c.Handler(next)
}
