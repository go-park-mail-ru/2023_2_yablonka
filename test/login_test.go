package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "server/internal/app/handlers"
	datatypes "server/internal/pkg/datatypes"

	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         uint64
		email          string
		passwordHash   string
		successful     bool
		expectedStatus int
	}{
		{
			name:           "Existing entry",
			userID:         1,
			email:          "test@email.com",
			passwordHash:   "123456",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User not found",
			email:          "notfound@email.com",
			passwordHash:   "123456",
			successful:     false,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Wrong password",
			email:          "test@email.com",
			passwordHash:   "654321",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			testApi := handlers.TestApi()

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password_hash":"%s"}`, test.email, test.passwordHash),
			))

			r := httptest.NewRequest("POST", "/api/v1/login/", body)
			w := httptest.NewRecorder()

			testApi.HandleLoginUser(w, r)

			require.EqualValuesf(t, test.expectedStatus, w.Code,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))

			user := datatypes.User{
				ID:           test.userID,
				Email:        test.email,
				PasswordHash: test.passwordHash,
			}
			expectedToken, err := testApi.GenerateJWT(&user)
			if test.successful {
				require.NoError(t, err)
				require.Contains(t, testApi.GetSessions(), expectedToken, "Expected token wasn't found")
			} else {
				require.NotContains(t, testApi.GetSessions(), expectedToken, "Expected token was found despite error")
			}
		})
	}
}
