package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"server/internal/app"
	"server/internal/app/handlers"
	"server/internal/apperrors"
	config "server/internal/config"
	jwt "server/internal/config/jwt"
	session "server/internal/config/session"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/service"
	authservice "server/internal/service/auth"
	boardservice "server/internal/service/board"
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
			name:           "[Session] Existing entry",
			userID:         1,
			email:          "test@email.com",
			password:       "12345678",
			serviceType:    "Session",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[Session] Wrong password",
			email:          "test@email.com",
			password:       "wrongpassword",
			serviceType:    "Session",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "[Session] User not found",
			email:          "notfound@email.com",
			password:       "doesntmatter",
			serviceType:    "Session",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "[JWT] User not found",
			email:          "notfound@email.com",
			password:       "doesntmatter",
			serviceType:    "JWT",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "[JWT] Existing entry",
			userID:         1,
			email:          "test@email.com",
			password:       "12345678",
			serviceType:    "JWT",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[JWT] Wrong password",
			email:          "test@email.com",
			password:       "wrongpassword",
			serviceType:    "JWT",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			userStorage := in_memory.NewUserStorage()
			boardStorage := in_memory.NewBoardStorage()

			authService := getAuthService(test.serviceType)
			userAuthService := userservice.NewAuthUserService(userStorage)
			boardService := boardservice.NewBoardService(boardStorage)

			mux, _ := app.GetChiMux(*handlers.NewHandlerManager(
				authService,
				userAuthService,
				//user.NewUserService(userStorage),
				boardService,
			))

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v1/auth/login/", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			status := w.Result().StatusCode

			require.EqualValuesf(t, test.expectedStatus, status,
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
			name:           "[Session] New user",
			email:          "newuser@email.com",
			password:       "100500600",
			serviceType:    "Session",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "[Session] New user",
			email:          "newuser@email.com",
			password:       "100500600",
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
			boardStorage := in_memory.NewBoardStorage()

			authService := getAuthService(test.serviceType)
			userAuthService := userservice.NewAuthUserService(userStorage)
			boardService := boardservice.NewBoardService(boardStorage)

			mux, _ := app.GetChiMux(*handlers.NewHandlerManager(
				authService,
				userAuthService,
				//user.NewUserService(userStorage),
				boardService,
			))

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v1/auth/signup/", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			status := w.Result().StatusCode

			require.EqualValuesf(t, test.expectedStatus, status,
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
				PasswordHash: "d40040163489d60c9adcbb768a6aa7a48ecc4b091bc8b43328fd51a46492fe75",
				Name:         "Никита",
				Surname:      "Архаров",
				ThumbnailURL: "https://sun1-27.userapi.com/s/v1/ig1/cAIfmwiDayww2WxVGPnIr5sHTSgXaf_567nuovSw_X4Cy9XAKrSVsAT2yAmJcJXDPkVOsXPW.jpg?size=50x50&quality=96&crop=351,248,540,540&ava=1",
			},
			password:       "12345678",
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
				PasswordHash: "72232ddaebd2ed90de93428069234c137b58f64f3d2ada273944fe5b52264310",
			},
			password:       "12345678",
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
				PasswordHash: "dd1ffd3fb76da152f41b103fb567910452708ad615b57876a63292797a041448",
				Name:         "Даниил",
				Surname:      "Капитанов",
				ThumbnailURL: "https://sun1-47.userapi.com/s/v1/ig2/aby-Y8KQ-yfQPLdvO-gq-ZenU63Iiw3ULbNlimdfaqLauSOj1cJ2jLxfBDtBMLpBW5T0UhaLFpyLVxAoYuVZiPB8.jpg?size=50x50&quality=95&crop=0,0,400,400&ava=1",
			},
			password:       "13579246",
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
			boardStorage := in_memory.NewBoardStorage()

			authService := getAuthService(test.serviceType)
			userAuthService := userservice.NewAuthUserService(userStorage)
			boardService := boardservice.NewBoardService(boardStorage)

			mux, _ := app.GetChiMux(*handlers.NewHandlerManager(
				authService,
				userAuthService,
				//user.NewUserService(userStorage),
				boardService,
			))

			body := bytes.NewReader([]byte(""))

			r := httptest.NewRequest("GET", "/api/v1/auth/verify/", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			rCtx := context.Background()

			if test.authorized {
				token, expiresAt, err := authService.AuthUser(rCtx, test.user)
				require.NoError(t, err)

				cookie := &http.Cookie{
					Name:     "tabula_user",
					Value:    token,
					HttpOnly: true,
					SameSite: http.SameSiteLaxMode,
					Expires:  expiresAt,
					Path:     "/api/v1/",
				}

				r.AddCookie(cookie)
			}

			mux.ServeHTTP(w, r)

			responseBody := w.Body.Bytes()
			status := w.Result().StatusCode

			require.EqualValuesf(t, test.expectedStatus, status,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))

			if test.authorized && test.successful {
				var jsonBody map[string]map[string]entities.User
				err := json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err)
				authedUser := jsonBody["body"]["user"]
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

func Test_GetUserBoards(t *testing.T) {
	// t.Parallel()

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
				Email: "test@email.com",
			},
			ownedBoards: 2,
			guestBoards: 1,
		},
		{
			name:       "Valid response (only owned boards)",
			authorized: true,
			user: &entities.User{
				ID:    2,
				Email: "email@example.com",
			},
			ownedBoards: 1,
			guestBoards: 0,
		},
		{
			name:       "Valid response (only guest boards)",
			authorized: true,
			user: &entities.User{
				ID:    3,
				Email: "newchallenger@email.com",
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
			// t.Parallel()

			userStorage := in_memory.NewUserStorage()
			boardStorage := in_memory.NewBoardStorage()

			authService := getAuthService("Session")
			userAuthService := userservice.NewAuthUserService(userStorage)
			boardService := boardservice.NewBoardService(boardStorage)

			handlerManager := handlers.NewHandlerManager(
				authService,
				userAuthService,
				boardService,
			)

			body := bytes.NewReader([]byte(""))

			r := httptest.NewRequest("GET", "/api/v1/auth/login/", body)
			w := httptest.NewRecorder()

			if test.authorized {
				ctx := context.Background()
				token, _, err := authService.AuthUser(ctx, test.user)
				require.NoError(t, err)
				userInfo, err := authService.VerifyAuth(ctx, token)
				require.NoError(t, err)
				userObj, err := userAuthService.GetUserByID(ctx, userInfo.UserID)
				require.NoError(t, err)
				*r = *r.WithContext(context.WithValue(r.Context(), dto.UserObjKey, userObj))
			}

			w.Header().Set("Content-Type", "application/json")

			handlerManager.BoardHandler.GetUserBoards(w, r)
			responseBody := w.Body.Bytes()

			if !test.authorized {
				err := r.Context().Value(dto.ErrorKey).(apperrors.ErrorResponse)
				status := err.Code
				require.Equal(t, http.StatusUnauthorized, status)
			} else {
				status := w.Result().StatusCode
				require.Equal(t, http.StatusOK, status)
				var jsonBody = dto.JSONResponse{
					Body: dto.JSONMap{
						"user": entities.User{},
						"boards": map[string]interface{}{
							"user_owned_boards": []dto.UserOwnedBoardInfo{},
							"user_guest_boards": []dto.UserGuestBoardInfo{},
						},
					},
				}
				err := json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err)
				bodyMap := jsonBody.Body.(map[string]interface{})
				userMap := bodyMap["user"].(map[string]interface{})
				require.Equal(t, test.user.ID, uint64(userMap["user_id"].(float64)))
				require.Equal(t, test.user.Email, userMap["email"])
				boardMap := bodyMap["boards"].(map[string]interface{})
				ownedBoardsI := boardMap["user_owned_boards"]
				guestBoardsI := boardMap["user_guest_boards"]
				var ownedBoards []interface{}
				var guestBoards []interface{}
				if ownedBoardsI != nil {
					ownedBoards = ownedBoardsI.([]interface{})
				}
				if guestBoardsI != nil {
					guestBoards = guestBoardsI.([]interface{})
				}
				require.Equal(t, test.ownedBoards, len(ownedBoards))
				require.Equal(t, test.guestBoards, len(guestBoards))
			}
		})
	}
}

func getAuthService(serviceType string) service.IAuthService {
	switch serviceType {
	case "JWT":
		config := jwt.JWTServerConfig{
			Base: config.BaseServerConfig{
				SessionDuration: time.Duration(14 * 24 * time.Hour),
			},
			JWTSecret: "TESTJWTSECRET123",
		}
		return authservice.NewAuthJWTService(config)
	case "Session":
		config := session.SessionServerConfig{
			Base: config.BaseServerConfig{
				SessionDuration: time.Duration(14 * 24 * time.Hour),
			},
			SessionIDLength: 32,
		}
		authStorage := in_memory.NewAuthStorage()
		return authservice.NewAuthSessionService(config, authStorage)
	default:
		config := session.SessionServerConfig{
			Base: config.BaseServerConfig{
				SessionDuration: time.Duration(1),
			},
			SessionIDLength: 32,
		}
		authStorage := in_memory.NewAuthStorage()
		return authservice.NewAuthSessionService(config, authStorage)
	}
}
