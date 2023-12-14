package middleware

import (
	"context"
	"errors"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/google/uuid"
)

func AuthMiddleware(as service.IAuthService, us service.IUserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			funcName := "AuthMiddleware"
			logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)
			requestID := rCtx.Value(dto.RequestIDKey).(uuid.UUID)

			logger.Info("*************** VERIFYING AUTH ***************")

			cookie, err := r.Cookie("tabula_user")
			if err != nil {
				logger.Error("*************** VERIFICATION FAIL ***************")
				logger.DebugFmt("Verifying user failed with error "+err.Error(), requestID.String(), funcName, nodeName)
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}
			logger.DebugFmt("Cookie found", requestID.String(), funcName, nodeName)

			token := dto.SessionToken{
				ID: cookie.Value,
			}

			userID, err := as.VerifyAuth(rCtx, token)
			if err != nil {
				logger.Error("*************** VERIFICATION FAIL ***************")
				logger.DebugFmt("Verifying user failed with error "+err.Error(), requestID.String(), funcName, nodeName)
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
				return
			}
			logger.DebugFmt("Session verified", requestID.String(), funcName, nodeName)

			userObj, err := us.GetWithID(rCtx, userID)
			if errors.Is(err, apperrors.ErrUserNotFound) {
				logger.Error("*************** VERIFICATION FAIL ***************")
				logger.DebugFmt("Verifying user failed with error "+err.Error(), requestID.String(), funcName, nodeName)
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			} else if err != nil {
				logger.Error("*************** VERIFICATION FAIL ***************")
				logger.DebugFmt("Verifying user failed with error "+err.Error(), requestID.String(), funcName, nodeName)
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
				return
			}

			logger.Info("*************** VERIFICATION SUCCESS ***************")

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.UserObjKey, userObj)))
		})
	}
}
