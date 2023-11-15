package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"

	"github.com/sirupsen/logrus"
)

func AuthMiddleware(as service.IAuthService, us service.IUserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			funcName := "AuthMiddleware"
			logger := rCtx.Value(dto.LoggerKey).(*logrus.Logger)

			logger.Info("Verifying user's authentication")

			cookie, err := r.Cookie("tabula_user")
			if err != nil {
				logger.Error("Verification failed")
				middlewareDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			}
			middlewareDebugLog(logger, funcName, "Cookie found")

			token := dto.SessionToken{
				ID: cookie.Value,
			}

			userID, err := as.VerifyAuth(rCtx, token)
			if err != nil {
				logger.Error("Verification failed")
				middlewareDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
				return
			}
			middlewareDebugLog(logger, funcName, "Session verified")

			middlewareDebugLog(logger, funcName, fmt.Sprintf("Getting user info for user ID %d", userID.Value))
			userObj, err := us.GetWithID(rCtx, userID)
			if errors.Is(err, apperrors.ErrUserNotFound) {
				logger.Error("Verification failed")
				middlewareDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.GenericUnauthorizedResponse, w, r)
				return
			} else if err != nil {
				logger.Error("Verification failed")
				middlewareDebugLog(logger, funcName, "Verifying user failed with error "+err.Error())
				w.Header().Set("X-Csrf-Token", "")
				apperrors.ReturnError(apperrors.ErrorMap[err], w, r)
				return
			}

			logger.Info("Finished verifying user")

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.UserObjKey, userObj)))
		})
	}
}
