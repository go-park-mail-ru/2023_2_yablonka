package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "server/internal/app/handlers"

	"github.com/stretchr/testify/require"
)

func Test_Signup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		email          string
		passwordHash   string
		successful     bool
		expectedStatus int
	}{
		{
			name:           "New user",
			email:          "newuser@email.com",
			passwordHash:   "qwerty",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Existing user",
			email:          "test@email.com",
			passwordHash:   "qwerty",
			successful:     false,
			expectedStatus: http.StatusConflict,
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

			r := httptest.NewRequest("POST", "/api/v1/signup/", body)
			w := httptest.NewRecorder()

			testApi.HandleSignupUser(w, r)

			require.EqualValuesf(t, test.expectedStatus, w.Code,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))

			if test.successful {
				require.Contains(t, testApi.GetUsers(), test.email, "User wasn't saved into storage")
			}
		})
	}
}
