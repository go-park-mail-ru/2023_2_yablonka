package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	config "server/internal/config"
	jwt "server/internal/config/jwt"
	session "server/internal/config/session"
	"server/internal/pkg/entities"
	"server/internal/service"
	authservice "server/internal/service/auth"
	userservice "server/internal/service/user"
	"server/internal/storage/in_memory"

	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
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
			userAuthService := userservice.NewAuthUserService(userStorage)

			switch test.serviceType {
			case "JWT":
				config := jwt.JWTServerConfig{
					Base: config.BaseServerConfig{
						SessionDuration: time.Duration(14 * 24 * time.Hour),
					},
					JWTSecret: "TESTJWTSECRET123",
				}
				authService = authservice.NewAuthJWTService(config.JWTSecret, config.Base.SessionDuration)
			case "Session":
				config := session.SessionServerConfig{
					Base: config.BaseServerConfig{
						SessionDuration: time.Duration(14 * 24 * time.Hour),
					},
					SessionIDLength: 32,
				}
				authStorage := in_memory.NewAuthStorage()
				authService = authservice.NewAuthSessionService(config.SessionIDLength, config.Base.SessionDuration, authStorage)
			}

			authHandler := NewAuthHandler(authService, userAuthService)
			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v1/auth/login/", body)
			w := httptest.NewRecorder()

			authHandler.LogIn(w, r)

			require.EqualValuesf(t, test.expectedStatus, w.Code,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))

			if test.successful {
				generatedCookie := w.Result().Cookies()[0].Value

				ctx := context.Background()
				verifiedAuth, err := authService.VerifyAuth(ctx, generatedCookie)

				require.NoError(t, err)
				require.Equal(t, test.userID, verifiedAuth.UserID, "Expected cookie wasn't found")
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
			userAuthService := userservice.NewAuthUserService(userStorage)

			switch test.serviceType {
			case "JWT":
				config := jwt.JWTServerConfig{
					Base: config.BaseServerConfig{
						SessionDuration: time.Duration(14 * 24 * time.Hour),
					},
					JWTSecret: "TESTJWTSECRET123",
				}
				authService = authservice.NewAuthJWTService(config.JWTSecret, config.Base.SessionDuration)
			case "Session":
				config := session.SessionServerConfig{
					Base: config.BaseServerConfig{
						SessionDuration: time.Duration(14 * 24 * time.Hour),
					},
					SessionIDLength: 32,
				}
				authStorage := in_memory.NewAuthStorage()
				authService = authservice.NewAuthSessionService(config.SessionIDLength, config.Base.SessionDuration, authStorage)
			}

			authHandler := NewAuthHandler(authService, userAuthService)

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v1/auth/signup/", body)
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
				verifiedAuth, err := authService.VerifyAuth(ctx, generatedCookie)

				require.NoError(t, err)

				expectedID := userStorage.GetHighestID()

				require.Equal(t, expectedID, verifiedAuth.UserID, "User wasn't saved correctly")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful registration")
			}
		})
	}
}

func Test_VerifyAuth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		email          string
		password       string
		serviceType    string
		user           *entities.User
		authorized     bool
		successful     bool
		expectedStatus int
	}{
		{
			name:        "[JWT] Valid authentification",
			serviceType: "JWT",
			user: &entities.User{
				ID:           1,
				Email:        "test@email.com",
				PasswordHash: "8a65f9232aec42190593cebe45067d14ade16eaf9aaefe0c2e9ec425b5b8ca73",
				Name:         "Никита",
				Surname:      "Архаров",
				ThumbnailURL: "https://sun1-27.userapi.com/s/v1/ig1/cAIfmwiDayww2WxVGPnIr5sHTSgXaf_567nuovSw_X4Cy9XAKrSVsAT2yAmJcJXDPkVOsXPW.jpg?size=50x50&quality=96&crop=351,248,540,540&ava=1",
			},
			authorized:     true,
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[JWT] No auth",
			serviceType:    "JWT",
			authorized:     false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:        "[JWT] Auth as nonexistent user",
			serviceType: "JWT",
			user: &entities.User{
				ID:           5,
				Email:        "fakeuser@email.com",
				PasswordHash: "8a65f9232aec42190593cebe45067d14ade16eaf9aaefe0c2e9ec425b5b8ca73",
			},
			authorized:     true,
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:        "[Session] Valid authentification",
			serviceType: "Session",
			user: &entities.User{
				ID:           2,
				Email:        "email@example.com",
				PasswordHash: "177e4fd1a8b22992e78145c3ba9c8781124e5c166d03b9c302cf8e100d77ad22",
				Name:         "Даниил",
				Surname:      "Капитанов",
				ThumbnailURL: "https://sun1-47.userapi.com/s/v1/ig2/aby-Y8KQ-yfQPLdvO-gq-ZenU63Iiw3ULbNlimdfaqLauSOj1cJ2jLxfBDtBMLpBW5T0UhaLFpyLVxAoYuVZiPB8.jpg?size=50x50&quality=95&crop=0,0,400,400&ava=1",
			},
			authorized:     true,
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[Session] No auth",
			serviceType:    "Session",
			authorized:     false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:        "[Session] Auth as nonexistent user",
			serviceType: "Session",
			user: &entities.User{
				ID:           5,
				Email:        "fakeuser@email.com",
				PasswordHash: "8a65f9232aec42190593cebe45067d14ade16eaf9aaefe0c2e9ec425b5b8ca73",
			},
			authorized:     true,
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
			userAuthService := userservice.NewAuthUserService(userStorage)

			switch test.serviceType {
			case "JWT":
				config := jwt.JWTServerConfig{
					Base: config.BaseServerConfig{
						SessionDuration: time.Duration(14 * 24 * time.Hour),
					},
					JWTSecret: "TESTJWTSECRET123",
				}
				authService = authservice.NewAuthJWTService(config.JWTSecret, config.Base.SessionDuration)
			case "Session":
				config := session.SessionServerConfig{
					Base: config.BaseServerConfig{
						SessionDuration: time.Duration(14 * 24 * time.Hour),
					},
					SessionIDLength: 32,
				}
				authStorage := in_memory.NewAuthStorage()
				authService = authservice.NewAuthSessionService(config.SessionIDLength, config.Base.SessionDuration, authStorage)
			}

			authHandler := NewAuthHandler(authService, userAuthService)

			body := bytes.NewReader([]byte(""))

			r := httptest.NewRequest("GET", "/api/v1/auth/verifyauth/", body)
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
			authHandler.VerifyAuth(w, r)
			status := w.Result().StatusCode
			responseBody := w.Body.Bytes()

			if !test.authorized || (test.authorized && !test.successful) {
				require.Equal(t, http.StatusUnauthorized, status)
			} else {
				require.Equal(t, http.StatusOK, status)
				var jsonBody map[string]entities.User
				err := json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err)
				authedUser := jsonBody["body"]
				require.Equal(t, test.user.ID, authedUser.ID)
				require.Equal(t, test.user.Email, authedUser.Email)
				require.Empty(t, authedUser.PasswordHash)
				require.Equal(t, test.user.Name, authedUser.Name)
				require.Equal(t, test.user.Surname, authedUser.Surname)
				require.Equal(t, test.user.ThumbnailURL, authedUser.ThumbnailURL)
			}
		})
	}
}
