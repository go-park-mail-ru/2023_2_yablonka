package middleware

import (
	"context"
	"errors"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"
)

func AuthMiddleware(as service.IAuthService, us service.IUserAuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			cookie, err := r.Cookie("tabula_user")
			if err != nil {
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}
			token := cookie.Value
			userInfo, err := as.VerifyAuth(rCtx, token)
			if err != nil {
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
				return
			}

			userObj, err := us.GetUserByID(rCtx, userInfo.UserID)

			if errors.Is(err, apperrors.ErrUserNotFound) {
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			} else if err != nil {
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.UserObjKey, userObj)))
		})
	}
}
