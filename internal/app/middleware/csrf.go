package middleware

import (
	"fmt"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/google/uuid"
)

func CSRFMiddleware(cs service.ICSRFService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := r.Context().Value(dto.LoggerKey).(logger.ILogger)
			requestID := r.Context().Value(dto.RequestIDKey).(uuid.UUID)
			funcName := "CSRFMiddleware"

			logger.Info("*************** VERIFYING CSRF ***************")

			rCtx := r.Context()

			csrf := r.Header.Get("X-Csrf-Token")
			if csrf == "" {
				logger.DebugFmt("CSRF verification failed", requestID.String(), funcName, nodeName)
				logger.Error("*************** CSRF FAIL ***************")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}
			logger.DebugFmt(fmt.Sprintf("Received CSRF token %s", csrf), requestID.String(), funcName, nodeName)

			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: csrf})
			if err != nil {
				logger.DebugFmt("CSRF verification failed with error "+err.Error(), requestID.String(), funcName, nodeName)
				logger.Error("*************** CSRF FAIL ***************")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}

			logger.Info("*************** CSRF SUCCESS ***************")

			next.ServeHTTP(w, r)
		})
	}
}
