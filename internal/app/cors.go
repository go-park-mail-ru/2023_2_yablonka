package app

import (
	"net/http"
	"strconv"
	"strings"
)

func EnableCors(h http.Handler, config Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", strings.Join(config.Server.AllowedHosts, ", "))
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.Server.AllowedMethods, ", "))
		w.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.Server.AllowCredentials))
		h.ServeHTTP(w, r)
	})
}
