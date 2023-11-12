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
			log.Println("************Verifying Auth************")

			rCtx := r.Context()

			log.Println("\tDEBUG cookie list")
			for _, c := range r.Cookies() {
				log.Println("\t", c)
			}

			cookie, err := r.Cookie("tabula_user")
			if err != nil {
				log.Println("ookie not found")
				log.Println("************Auth FAILED************")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}

			token := dto.SessionToken{
				ID: cookie.Value,
			}

			userID, err := as.VerifyAuth(rCtx, token)
			if err != nil {
				log.Println("Session not found")
				log.Println("************Auth FAILED************")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
				return
			}

			userObj, err := us.GetWithID(rCtx, userID)
			if errors.Is(err, apperrors.ErrUserNotFound) {
				log.Println("User not found")
				log.Println("************Auth FAILED************")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			} else if err != nil {
				log.Println("Other error")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.ErrorMap[err]))
				return
			}

			log.Println("************Auth Verified************")

			next.ServeHTTP(w, r.WithContext(context.WithValue(rCtx, dto.UserObjKey, userObj)))
		})
	}
}
