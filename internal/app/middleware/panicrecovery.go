package middleware

import (
	"log"
	"net/http"
	"server/internal/apperrors"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcvr := recover(); rcvr != nil {
				log.Println("recovered from", rcvr)
				w.WriteHeader(http.StatusInternalServerError)
				response := apperrors.ErrorJSON(apperrors.InternalServerErrorResponse)
				w.Write(response)
				r.Body.Close()
			}
		}()

		next.ServeHTTP(w, r)
	})
}
