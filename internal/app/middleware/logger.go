package middleware

import (
	"context"
	"net/http"
	"server/internal/config"
	"server/internal/pkg/dto"

	"github.com/sirupsen/logrus"
)

func SetLogger(config config.LoggingConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			logger := config.Logger

			logger.
				WithFields(logrus.Fields{
					"route_node": "middleware",
					"function":   "SetLogger",
				}).
				Debug("Added logger to context")
			logger.Info(r.URL.Path)

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.LoggerKey, logger)))
		})
	}
}
