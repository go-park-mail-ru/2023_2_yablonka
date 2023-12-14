package middleware

import (
	"net/http"
	logger "server/internal/logging"
	"server/internal/pkg/dto"

	"github.com/google/uuid"
)

func JsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(dto.LoggerKey).(logger.ILogger)
		requestID := r.Context().Value(dto.RequestIDKey).(uuid.UUID)
		funcName := "JsonHeader"

		w.Header().Set("Content-Type", "application/json")
		logger.DebugFmt("Content type header set", requestID.String(), funcName, nodeName)

		next.ServeHTTP(w, r)
	})
}
