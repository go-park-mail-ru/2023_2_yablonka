package middleware

import (
	"net/http"
	"reflect"
	"runtime"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
)

func HandlerLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := r.Context().Value(dto.LoggerKey).(logger.ILogger)

		pc := reflect.ValueOf(next).Pointer()
		funcName := runtime.FuncForPC(pc).Name()

		logger.Info("--------------" + r.URL.Path + "--------------")
		next.ServeHTTP(w, r)
		logger.Info("--------------" + funcName + "--------------")
	})
}
