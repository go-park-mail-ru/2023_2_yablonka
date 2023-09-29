package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "server/internal/app/handlers"
	"server/internal/app/utils"
	datatypes "server/internal/pkg/datatypes"
	authservice "server/internal/service/auth"
	userservice "server/internal/service/user"
	"server/internal/storage"

	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         uint64
		email          string
		password       string
		successful     bool
		expectedStatus int
	}{
		{
			name:           "Existing entry",
			userID:         1,
			email:          "test@email.com",
			password:       "123456",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User not found",
			email:          "notfound@email.com",
			password:       "doesntmatter",
			successful:     false,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Wrong password",
			email:          "test@email.com",
			password:       "wrongpassword",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			userStorage := storage.NewLocalUserStorage()

			authService := authservice.NewAuthJWTService()
			userAuthService := userservice.NewAuthUserService(userStorage)
			authHandler := handlers.NewAuthHandler(authService, userAuthService)

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v1/login/", body)
			w := httptest.NewRecorder()

			authHandler.LogIn(w, r)

			require.EqualValuesf(t, test.expectedStatus, w.Code,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))

			if test.successful {
				authInfo := datatypes.AuthInfo{
					Email:    test.email,
					Password: test.password,
				}

				passwordHash := utils.HashFromAuthInfo(authInfo)
				user := &datatypes.User{
					ID:           test.userID,
					Email:        test.email,
					PasswordHash: passwordHash,
				}
				ctx := context.Background()
				expectedToken, err := authService.AuthUser(ctx, user)

				generatedToken := w.Result().Cookies()[0].Value
				require.NoError(t, err)
				require.Equal(t, expectedToken, generatedToken, "Expected token wasn't found")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful authentication")
			}
		})
	}
}

func Test_Signup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		email          string
		password       string
		successful     bool
		expectedStatus int
	}{
		{
			name:           "New user",
			email:          "newuser@email.com",
			password:       "100500",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User already exists",
			email:          "test@email.com",
			password:       "coolpassword",
			successful:     false,
			expectedStatus: http.StatusConflict,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			userStorage := storage.NewLocalUserStorage()

			authService := authservice.NewAuthJWTService()
			userAuthService := userservice.NewAuthUserService(userStorage)
			authHandler := handlers.NewAuthHandler(authService, userAuthService)

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v1/signup/", body)
			w := httptest.NewRecorder()

			authHandler.SignUp(w, r)

			require.EqualValuesf(t, test.expectedStatus, w.Code,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))

			if test.successful {
				require.Equal(t, w.Result().Cookies()[0].Name, "user_jwt", "Expected token wasn't found")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful authentication")
			}
		})
	}
}
