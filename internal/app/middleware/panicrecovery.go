package middleware

import (
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"

	"github.com/sirupsen/logrus"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcvr := recover(); rcvr != nil {
				logger := r.Context().Value(dto.LoggerKey).(*logrus.Logger)
				logger.Error("***** PANIC *****")
				logger.Error("Recovered from panic ", rcvr)

				apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)

				logger.Error("***** CONTINUING *****")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
