package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"
)

func AuthMiddleware(as service.IAuthService, us service.IUserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			cookie, err := r.Cookie("tabula_user")
			if err != nil {
				log.Println("Middleware -- Cookie not found")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}

			token := dto.SessionToken{
				ID: cookie.Value,
			}

			userID, err := as.VerifyAuth(rCtx, token)
			if err != nil {
				log.Println("Middleware -- Session not found")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
				return
			}

			userObj, err := us.GetWithID(rCtx, userID)

			if errors.Is(err, apperrors.ErrUserNotFound) {
				log.Println("User not found")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			} else if err != nil {
				log.Println("Other error")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.UserObjKey, userObj)))
		})
	}
}
