package parsers

import (
	"context"
	"net/http"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func BoardCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "boardID")
		boardID, err := strconv.ParseUint(param, 10, 64)
		if err != nil {
			apperrors.ReturnError(apperrors.BadRequestResponse, w, r)
		} else {
			ctx := context.WithValue(r.Context(), dto.BoardIDKey, boardID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
