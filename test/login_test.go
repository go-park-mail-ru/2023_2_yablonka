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
		name          string
		email         string
		password_hash string
	}{
		{
			name:          "Existing entry",
			email:         "test@email.com",
			password_hash: "123456",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			testApi := handlers.TestApi()

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password_hash":"%s"}`, test.email, test.password_hash),
			))

			r := httptest.NewRequest("POST", "/api/v1/login/", body)
			w := httptest.NewRecorder()

			testApi.HandleLoginUser(w, r)

			login := datatypes.LoginInfo{
				Email:        test.email,
				PasswordHash: test.password_hash,
			}
			expectedToken, err := testApi.GenerateJWT(login)

			if w.Code != http.StatusOK {
				t.Error("Status is not ok")
			}

			require.NoError(t, err)
			require.Contains(t, testApi.GetSessions(), expectedToken, "Expected token wasn't found")
		})
	}
}
