package middleware

import (
	"fmt"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/service"
)

func CSRFMiddleware(cs service.ICSRFService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := r.Context().Value(dto.LoggerKey).(logger.ILogger)
			funcName := "CSRFMiddleware"

			logger.Info("***** VERIFYING CSRF *****")

			rCtx := r.Context()

			csrf := r.Header.Get("X-Csrf-Token")
			if csrf == "" {
				logger.Debug("CSRF verification failed", funcName, "middleware")
				logger.Error("***** CSRF FAIL *****")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}
			logger.Debug(fmt.Sprintf("Received CSRF token %s", csrf), funcName, "middleware")

			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: csrf})
			if err != nil {
				logger.Debug("CSRF verification failed with error "+err.Error(), funcName, "middleware")
				logger.Error("***** CSRF FAIL *****")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}

			logger.Info("***** CSRF SUCCESS *****")

			next.ServeHTTP(w, r)
		})
	}
}
