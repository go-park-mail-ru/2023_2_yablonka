package middleware

import (
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		ctx := r.Context()
		
		errorResponse, ok := ctx.Value(dto.ErrorKey).(apperrors.ErrorResponse)
		if ok {
			w.WriteHeader(errorResponse.Code)
			response := apperrors.ErrorJSON(errorResponse)
			_, _ = w.Write(response)
			r.Body.Close()
		}
	})
}
