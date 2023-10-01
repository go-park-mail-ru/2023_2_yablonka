package handlers

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/internal/service"
	authservice "server/internal/service/auth"
	userservice "server/internal/service/user"
	"server/internal/storage/in_memory"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
	err := godotenv.Load("../../../cmd/app/.env")
	require.NoError(t, err)
	t.Parallel()

	tests := []struct {
		name           string
		userID         uint64
		email          string
		password       string
		serviceType    string
		successful     bool
		expectedStatus int
	}{
		{
			name:           "[JWT] Existing entry",
			userID:         1,
			email:          "test@email.com",
			password:       "123456",
			serviceType:    "JWT",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[JWT] User not found",
			email:          "notfound@email.com",
			password:       "doesntmatter",
			serviceType:    "JWT",
			successful:     false,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "[JWT] Wrong password",
			email:          "test@email.com",
			password:       "wrongpassword",
			serviceType:    "JWT",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "[Session] Existing entry",
			userID:         1,
			email:          "test@email.com",
			password:       "123456",
			serviceType:    "Session",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[Session] User not found",
			email:          "notfound@email.com",
			password:       "doesntmatter",
			serviceType:    "Session",
			successful:     false,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "[Session] Wrong password",
			email:          "test@email.com",
			password:       "wrongpassword",
			serviceType:    "Session",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			userStorage := in_memory.NewUserStorage()
			var authService service.IAuthService
			var initerr error

			switch test.serviceType {
			case "JWT":
				authService, initerr = authservice.NewAuthJWTService()
				require.NoError(t, initerr)
			case "Session":
				authStorage, initerr := in_memory.NewAuthStorage()
				require.NoError(t, initerr)
				authService, initerr = authservice.NewAuthSessionService(authStorage)
				require.NoError(t, initerr)
			}

			userAuthService := userservice.NewAuthUserService(userStorage)
			authHandler := NewAuthHandler(authService, userAuthService)

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
				generatedCookie := w.Result().Cookies()[0].Value

				ctx := context.Background()
				uid, err := authService.VerifyAuth(ctx, generatedCookie)

				require.NoError(t, err)
				require.Equal(t, test.userID, uid, "Expected cookie wasn't found")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful authentication")
			}
		})
	}
}

func Test_Signup(t *testing.T) {
	err := godotenv.Load("../../../cmd/app/.env")
	require.NoError(t, err)
	t.Parallel()

	tests := []struct {
		name           string
		email          string
		password       string
		serviceType    string
		successful     bool
		expectedStatus int
	}{
		{
			name:           "[JWT] New user",
			email:          "newuser@email.com",
			password:       "100500",
			serviceType:    "JWT",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[JWT] User already exists",
			email:          "test@email.com",
			password:       "coolpassword",
			serviceType:    "JWT",
			successful:     false,
			expectedStatus: http.StatusConflict,
		},
		{
			name:           "[Session] New user",
			email:          "newuser@email.com",
			password:       "100500",
			serviceType:    "Session",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[Session] User already exists",
			email:          "test@email.com",
			password:       "coolpassword",
			serviceType:    "Session",
			successful:     false,
			expectedStatus: http.StatusConflict,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			userStorage := in_memory.NewUserStorage()

			var authService service.IAuthService
			var initerr error

			switch test.serviceType {
			case "JWT":
				authService, initerr = authservice.NewAuthJWTService()
				require.NoError(t, initerr)
			case "Session":
				authStorage, initerr := in_memory.NewAuthStorage()
				require.NoError(t, initerr)
				authService, initerr = authservice.NewAuthSessionService(authStorage)
				require.NoError(t, initerr)
			}

			userAuthService := userservice.NewAuthUserService(userStorage)
			authHandler := NewAuthHandler(authService, userAuthService)

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
				require.Equal(t, w.Result().Cookies()[0].Name, "tabula_user", "Expected cookie wasn't found")
				generatedCookie := w.Result().Cookies()[0].Value

				ctx := context.Background()
				newID, err := authService.VerifyAuth(ctx, generatedCookie)

				require.NoError(t, err)

				expectedID := userStorage.GetHighestID()

				require.Equal(t, expectedID, newID, "User wasn't saved correctly")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful registration")
			}
		})
	}
}
