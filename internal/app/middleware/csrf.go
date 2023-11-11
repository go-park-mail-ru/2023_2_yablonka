package middleware

import (
	"context"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/service"
)

func CSRFMiddleware(cs service.ICSRFService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rCtx := r.Context()

			err := cs.VerifyCSRF(rCtx, dto.CSRFToken{Value: r.Header.Get("X-CSRF-Token")})
			if err != nil {
				*r = *r.WithContext(context.WithValue(rCtx, dto.ErrorKey, apperrors.GenericUnauthorizedResponse))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
