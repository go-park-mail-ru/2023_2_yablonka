package parsers

import (
	"context"
	"log"
	"net/http"
	"server/internal/pkg/dto"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userInfo := chi.URLParam(r, "userID")
		log.Println(userInfo)
		id, err := strconv.ParseUint(userInfo, 10, 64)
		var ctx context.Context
		if err != nil {
			ctx = context.WithValue(r.Context(), dto.UserIDKey, userInfo)
		} else {
			ctx = context.WithValue(r.Context(), dto.UserIDKey, id)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
