package middleware

import (
	"context"
	"net/http"
	"server/internal/config"
	logger "server/internal/logging"
	"server/internal/pkg/dto"

	"github.com/google/uuid"
)

const nodeName = "middleware"

func SetContext(sc config.ServerConfig, logger logger.ILogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("*************** SETTING UP CONTEXT ***************")

			funcName := "SetLogger"

			logger.Info(r.URL.Path)
			requestID := uuid.New()

			rCtx := context.WithValue(r.Context(), dto.LoggerKey, logger)
			logger.DebugFmt("Added logger to context", requestID.String(), funcName, nodeName)

			rCtx = context.WithValue(rCtx, dto.RequestIDKey,
				uuid.New(),
			)
			logger.DebugFmt("Added request ID to context", requestID.String(), funcName, nodeName)

			logger.Info("*************** CONTEXT SET UP ***************")

			next.ServeHTTP(w, r.WithContext(rCtx))
		})
	}
}
