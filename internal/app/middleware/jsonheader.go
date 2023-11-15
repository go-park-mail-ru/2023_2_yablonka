package middleware

import (
	"net/http"
	"server/internal/pkg/dto"

	"github.com/sirupsen/logrus"
)

func JsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(dto.LoggerKey).(*logrus.Logger)
		funcName := "JsonHeader"

		w.Header().Set("Content-Type", "application/json")
		middlewareDebugLog(logger, funcName, "Content type header set")

		next.ServeHTTP(w, r)
	})
}
