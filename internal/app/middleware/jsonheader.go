package middleware

import (
	"net/http"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
)

func JsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(dto.LoggerKey).(logger.ILogger)
		funcName := "JsonHeader"

		w.Header().Set("Content-Type", "application/json")
		logger.DebugFmt("Content type header set", funcName, "middleware")

		next.ServeHTTP(w, r)
	})
}
