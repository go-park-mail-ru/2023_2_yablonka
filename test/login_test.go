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

func TestSuccess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		userID       uint64
		email        string
		passwordHash string
	}{
		{
			name:         "Existing entry",
			userID:       1,
			email:        "test@email.com",
			passwordHash: "123456",
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

			user := datatypes.User{
				ID:           test.userID,
				Email:        test.email,
				PasswordHash: test.passwordHash,
			}
			expectedToken, err := testApi.GenerateJWT(&user)

			if w.Code != http.StatusOK {
				t.Error("Status is not ok")
			}

			require.NoError(t, err)
			require.Contains(t, testApi.GetSessions(), expectedToken, "Expected token wasn't found")
		})
	}
}

func TestFail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		email         string
		passwordHash  string
		expectedError int
	}{
		{
			name:          "User not found",
			email:         "notfound@email.com",
			passwordHash:  "123456",
			expectedError: http.StatusNotFound,
		},
		{
			name:          "Wrong password",
			email:         "test@email.com",
			passwordHash:  "654321",
			expectedError: http.StatusUnauthorized,
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

			require.EqualValues(t, w.Code, test.expectedError)
		})
	}
}
