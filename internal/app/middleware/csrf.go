package middleware

import (
	"fmt"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service/csrf"

	"github.com/sirupsen/logrus"
)

func CSRFMiddleware(cs csrf.ICSRFService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := r.Context().Value(dto.LoggerKey).(*logrus.Logger)
			funcName := "CSRFMiddleware"

			logger.Info("***** VERIFYING CSRF *****")

			rCtx := r.Context()

			csrf := r.Header.Get("X-Csrf-Token")
			if csrf == "" {
				middlewareDebugLog(logger, funcName, "CSRF verification failed")
				logger.Error("***** CSRF FAIL *****")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}

			middlewareDebugLog(logger, funcName, fmt.Sprintf("Received CSRF token %s", csrf))

			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: csrf})
			if err != nil {
				middlewareDebugLog(logger, funcName, "CSRF verification failed with error "+err.Error())
				logger.Error("***** CSRF FAIL *****")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}

			logger.Info("***** CSRF SUCCESS *****")

			next.ServeHTTP(w, r)
		})
	}
}
