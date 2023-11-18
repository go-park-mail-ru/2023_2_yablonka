package middleware

import (
	"context"
	"net/http"
	"server/internal/config"
	"server/internal/pkg/dto"
	"strconv"
)

func SetLogger(lc config.LoggingConfig, sc config.ServerConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			logger := lc.Logger
			funcName := "SetLogger"

			logger.Info(r.URL.Path)

			rCtx = context.WithValue(rCtx, dto.LoggerKey, logger)
			middlewareDebugLog(logger, funcName, "Added logger to context")

			rCtx = context.WithValue(rCtx, dto.BaseURLKey,
				"http://"+sc.Host+":"+strconv.FormatUint(uint64(sc.BackendPort), 10)+"/",
			)
			middlewareDebugLog(logger, funcName, "Added base URL to context")

			next.ServeHTTP(w, r.WithContext(rCtx))
		})
	}
}
