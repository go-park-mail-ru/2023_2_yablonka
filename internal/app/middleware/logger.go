package middleware

import (
	"context"
	"net/http"
	"server/internal/config"
	"server/internal/pkg/dto"
)

func SetLogger(config config.LoggingConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			logger := config.Logger
			funcName := "SetLogger"

			logger.Info(r.URL.Path)
			middlewareDebugLog(logger, funcName, "Added logger to context")

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.LoggerKey, logger)))
		})
	}
}
