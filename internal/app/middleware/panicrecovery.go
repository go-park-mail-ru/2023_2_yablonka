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
				log.Println("************PANIC************")

				log.Println("recovered from", rcvr)

				w.WriteHeader(http.StatusInternalServerError)
				response := apperrors.ErrorJSON(apperrors.InternalServerErrorResponse)
				_, _ = w.Write(response)

				r.Body.Close()
				log.Println("************CONTINUING************")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
