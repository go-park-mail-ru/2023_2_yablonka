package middleware

import (
	"net/http"
	"server/internal/pkg/dto"

	"github.com/sirupsen/logrus"
)

func JsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(dto.LoggerKey).(*logrus.Logger)

		logger.
			WithFields(logrus.Fields{
				"route_node": "middleware",
				"function":   "JsonHeader",
			}).
			Debug("Setting content type header")

		w.Header().Set("Content-Type", "application/json")

		logger.
			WithFields(logrus.Fields{
				"route_node": "middleware",
				"function":   "JsonHeader",
			}).
			Debug("Content type header set")

		next.ServeHTTP(w, r)
	})
}
