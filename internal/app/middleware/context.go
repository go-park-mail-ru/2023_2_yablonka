package middleware

import (
	"context"
	"net/http"
	"server/internal/config"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"strconv"
)

func SetContext(sc config.ServerConfig, logger logger.ILogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("***** SETTING UP LOGGER *****")

			funcName := "SetLogger"

			logger.Info(r.URL.Path)

			rCtx := context.WithValue(r.Context(), dto.LoggerKey, logger)
			logger.Debug("Added logger to context", funcName, "middleware")

			rCtx = context.WithValue(rCtx, dto.BaseURLKey,
				"http://"+sc.Host+":"+strconv.FormatUint(uint64(sc.BackendPort), 10)+"/",
			)
			logger.Debug("Added base URL to context", funcName, "middleware")

			logger.Info("***** LOGGER SET UP *****")

			next.ServeHTTP(w, r.WithContext(rCtx))
		})
	}
}
