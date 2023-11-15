package middleware

import (
	"fmt"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/sirupsen/logrus"
)

func CSRFMiddleware(cs service.ICSRFService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := r.Context().Value(dto.LoggerKey).(*logrus.Logger)
			funcName := "CSRFMiddleware"

			rCtx := r.Context()

			csrf := r.Header.Get("X-Csrf-Token")
			if csrf == "" {
				logger.Error("CSRF header not set on incoming request")
				middlewareDebugLog(logger, funcName, "CSRF verification failed")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}

			middlewareDebugLog(logger, funcName, fmt.Sprintf("Received CSRF token %s", csrf))

			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: csrf})
			if err != nil {
				logger.Error("Failed to verify CSRF")
				middlewareDebugLog(logger, funcName, "CSRF verification failed with error "+err.Error())
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}

			middlewareDebugLog(logger, funcName, "CSRF verification succeeded")

			next.ServeHTTP(w, r)
		})
	}
}
