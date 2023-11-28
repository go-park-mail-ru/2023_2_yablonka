package middleware

import (
	"context"
	"errors"
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/service"
)

func AuthMiddleware(as auth.IAuthService, us user.IUserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			funcName := "AuthMiddleware"
			logger := rCtx.Value(dto.LoggerKey).(logger.ILogger)

			logger.Info("***** VERIFYING AUTH *****")

			cookie, err := r.Cookie("tabula_user")
			if err != nil {
				logger.Error("***** VERIFICATION FAIL *****")
				logger.Debug("Verifying user failed with error "+err.Error(), funcName, "middleware")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}
			logger.Debug("Cookie found", funcName, "middleware")

			token := dto.SessionToken{
				ID: cookie.Value,
			}

			userID, err := as.VerifyAuth(rCtx, token)
			if err != nil {
				logger.Error("***** VERIFICATION FAIL *****")
				logger.Debug("Verifying user failed with error "+err.Error(), funcName, "middleware")
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
				return
			}
			logger.Debug("Session verified", funcName, "middleware")

			userObj, err := us.GetWithID(rCtx, userID)
			if errors.Is(err, apperrors.ErrUserNotFound) {
				logger.Error("***** VERIFICATION FAIL *****")
				logger.Debug("Verifying user failed with error "+err.Error(), funcName, "middleware")
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			} else if err != nil {
				logger.Error("***** VERIFICATION FAIL *****")
				logger.Debug("Verifying user failed with error "+err.Error(), funcName, "middleware")
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
				return
			}

			logger.Info("***** VERIFICATION SUCCESS *****")

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.UserObjKey, userObj)))
		})
	}
}
