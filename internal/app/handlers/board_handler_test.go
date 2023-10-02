package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	authservice "server/internal/service/auth"
	boardservice "server/internal/service/board"
	"server/internal/storage/in_memory"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_GetUserBoards(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		authorized  bool
		user        *entities.User
		ownedBoards int
		guestBoards int
	}{
		{
			name:       "Valid response (owned and guest boards)",
			authorized: true,
			user: &entities.User{
				ID:    1,
				Email: "test@example.com",
			},
			ownedBoards: 2,
			guestBoards: 1,
		},
		{
			name:       "Valid response (only owned boards)",
			authorized: true,
			user: &entities.User{
				ID:    2,
				Email: "example@email.com",
			},
			ownedBoards: 1,
			guestBoards: 0,
		},
		{
			name:       "Valid response (only guest boards)",
			authorized: true,
			user: &entities.User{
				ID:    3,
				Email: "newchallenger@example.com",
			},
			ownedBoards: 0,
			guestBoards: 2,
		},
		{
			name:       "Valid response (no boards)",
			authorized: true,
			user: &entities.User{
				ID:    4,
				Email: "ghostinthem@chi.ne",
			},
			ownedBoards: 0,
			guestBoards: 0,
		},
		{
			name:       "Auth error",
			authorized: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			serverConfig := &entities.ServerConfig{
				SessionDuration: time.Duration(14 * 24 * time.Hour),
				SessionIDLength: 32,
				JWTSecret:       "TESTJWTSECRET123",
			}
			boardStorage := in_memory.NewBoardStorage()
			boardService := boardservice.NewBoardService(boardStorage)
			authStorage := in_memory.NewAuthStorage()
			authService := authservice.NewAuthSessionService(serverConfig, authStorage)
			boardHandler := NewBoardHandler(authService, boardService)

			body := bytes.NewReader([]byte(""))

			r := httptest.NewRequest("GET", "/api/v1/auth/login/", body)
			w := httptest.NewRecorder()

			if test.authorized {
				ctx := context.Background()
				token, expiresAt, err := authService.AuthUser(ctx, test.user)
				require.NoError(t, err)
				cookie := &http.Cookie{
					Name:     "tabula_user",
					Value:    token,
					HttpOnly: true,
					SameSite: http.SameSiteDefaultMode,
					Expires:  expiresAt,
				}
				r.AddCookie(cookie)
			}

			w.Header().Set("Content-Type", "application/json")

			boardHandler.GetUserBoards(w, r)
			status := w.Result().StatusCode
			responseBody := w.Body.Bytes()

			if !test.authorized {
				require.Equal(t, http.StatusUnauthorized, status)
			} else {
				require.Equal(t, http.StatusOK, status)
				var jsonBody map[string]dto.UserTotalBoardInfo
				err := json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err)
				userBoards := jsonBody["body"]
				require.Equal(t, test.ownedBoards, len(userBoards.OwnedBoards))
				require.Equal(t, test.guestBoards, len(userBoards.GuestBoards))
			}
		})
	}
}
