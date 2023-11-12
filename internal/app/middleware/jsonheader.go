package middleware

import (
	"log"
	"net/http"
)

func JsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("************Setting content type header************")
		w.Header().Set("Content-Type", "application/json")
		log.Println("************Content type header set************")

		next.ServeHTTP(w, r)
	})
}
