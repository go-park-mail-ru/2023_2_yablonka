package middleware

import (
	"net/http"
	"server/internal/apperrors"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		ctx := r.Context()
		errorResponse, ok := ctx.Value("errorResponse").(apperrors.ErrorResponse)
		if ok {
			w.WriteHeader(errorResponse.Code)
			response := apperrors.ErrorJSON(errorResponse)
			w.Write(response)
			r.Body.Close()
		}
	})
}
