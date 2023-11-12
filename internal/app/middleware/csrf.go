package middleware

import (
	"context"
	"log"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"
)

func CSRFMiddleware(cs service.ICSRFService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()
			log.Println("CSRF Middleware -- Verifying CSRF")
			csrf := r.Header.Get("X-Csrf-Token")
			if csrf == "" {
				log.Println("CSRF Middleware -- CSRF header not set on incoming request")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}
			log.Println("CSRF Middleware -- Received CSRF token", csrf)
			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: csrf})
			if err != nil {
				log.Println("CSRF Middleware -- CSRF invalid")
				log.Println(err)
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}
			log.Println("CSRF Middleware -- CSRF verified")
			next.ServeHTTP(w, r)
		})
	}
}
