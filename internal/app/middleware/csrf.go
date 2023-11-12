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
			log.Println("Verifying CSRF")
			csrf := r.Header.Get("X-CSRF-Token")
			if csrf == "" {
				log.Println("CSRF header not set on incoming request")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}
			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: csrf})
			if err != nil {
				log.Println("CSRF invalid")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}
			log.Println("CSRF verified")
			next.ServeHTTP(w, r)
		})
	}
}
