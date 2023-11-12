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
			log.Println("************Verifying CSRF************")

			rCtx := r.Context()

			csrf := r.Header.Get("X-CSRF-Token")
			if csrf == "" {
				log.Println("CSRF header not set on incoming request")
				log.Println("************CSRF INVALID************")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}

			log.Println("Received CSRF token", csrf)

			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: csrf})
			if err != nil {
				log.Println(err)
				log.Println("************CSRF INVALID************")
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}

			log.Println("************CSRF INVALID************")

			next.ServeHTTP(w, r)
		})
	}
}
