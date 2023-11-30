package middleware

import (
	"net/http"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcvr := recover(); rcvr != nil {
				logger := r.Context().Value(dto.LoggerKey).(logger.ILogger)
				logger.Error("*************** PANIC ***************")
				logger.Error("Recovered from panic " + rcvr.(string))

				apperrors.ReturnError(apperrors.InternalServerErrorResponse, w, r)

				logger.Error("*************** CONTINUING ***************")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
