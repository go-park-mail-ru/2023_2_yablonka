package handlers_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"server/internal/app"
	"server/internal/app/handlers"
	"server/internal/apperrors"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/mocks/mock_service"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const configPath string = "../../config/config.yml"
const envPath string = ""

func hashFromAuthInfo(info dto.AuthInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(info.Email + info.Password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func createConfig(envPath string) (*config.BaseServerConfig, error) {
	cfgFile, err := os.Create(envPath + ".env")
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()

	_, err = fmt.Fprint(cfgFile,
		"JWT_SECRET='test secret'"+"\n"+
			"SESSION_DURATION_DAYS=14"+"\n"+
			"SESSION_DURATION_HOURS=0"+"\n"+
			"SESSION_DURATION_MINUTES=0"+"\n"+
			"SESSION_DURATION_SECONDS=0"+"\n"+
			"SESSION_ID_LENGTH=32"+"\n",
	)
	if err != nil {
		return nil, err
	}

	config, err := config.NewBaseEnvConfig(envPath, configPath)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func TestAuthHandler_LogIn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         uint64
		email          string
		password       string
		successful     bool
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "Successful login",
			userID:         1,
			email:          "test@email.com",
			password:       "a1234567",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Wrong password",
			email:          "test@email.com",
			password:       "wrongpassword",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
			expectedError:  apperrors.ErrWrongPassword,
		},
		{
			name:           "User not found",
			email:          "notfound@email.com",
			password:       "doesntmatter",
			successful:     false,
			expectedStatus: http.StatusUnauthorized,
			expectedError:  apperrors.ErrUserNotFound,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			cfg, err := createConfig(envPath)
			require.Equal(t, nil, err)

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserAuthService := mock_service.NewMockIUserAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			authInfo := dto.AuthInfo{
				Email:    test.email,
				Password: test.password,
			}

			loginInfo := dto.LoginInfo{
				Email:        test.email,
				PasswordHash: hashFromAuthInfo(authInfo),
			}
			mockUA := mockUserAuthService.EXPECT().Login(gomock.Any(), loginInfo)
			if test.successful {
				mockUA.Return(
					&entities.User{
						ID:           test.userID,
						Email:        test.email,
						PasswordHash: test.password,
					},
					nil,
				)
				mockAuthService.
					EXPECT().
					AuthUser(gomock.Any(), test.userID)
				mockAuthService.
					EXPECT().
					VerifyAuth(gomock.Any(), gomock.Any()).
					Return(test.userID, nil)
			} else {
				mockUA.Return(nil, test.expectedError)
			}

			mux, _ := app.GetChiMux(*handlers.NewHandlerManager(
				mockAuthService,
				mockUserAuthService,
				mockBoardService,
			),
				*cfg,
			)

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v2/auth/login/", body)
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
				userID, err := mockAuthService.VerifyAuth(ctx, generatedCookie)

				require.NoError(t, err)
				require.Equal(t, test.userID, userID, "Expected cookie wasn't found")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful authentication")
			}
		})
	}
}

func TestAuthHandler_SignUp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		email          string
		password       string
		successful     bool
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "New user",
			email:          "newuser@email.com",
			password:       "100500600",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "User already exists",
			email:          "test@email.com",
			password:       "coolpassword",
			successful:     false,
			expectedStatus: http.StatusConflict,
			expectedError:  apperrors.ErrUserAlreadyExists,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			cfg, err := createConfig(envPath)
			require.Equal(t, nil, err)

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserAuthService := mock_service.NewMockIUserAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			authInfo := dto.AuthInfo{
				Email:    test.email,
				Password: test.password,
			}

			signupInfo := dto.SignupInfo{
				Email:        test.email,
				PasswordHash: hashFromAuthInfo(authInfo),
			}
			mockUA := mockUserAuthService.EXPECT().RegisterUser(gomock.Any(), signupInfo)
			if test.successful {
				mockUA.Return(
					&entities.User{
						ID:           uint64(1),
						Email:        test.email,
						PasswordHash: test.password,
					},
					nil,
				)
				mockAuthService.
					EXPECT().
					AuthUser(gomock.Any(), uint64(1))
				mockAuthService.
					EXPECT().
					VerifyAuth(gomock.Any(), gomock.Any()).
					Return(uint64(1), nil)
			} else {
				mockUA.Return(nil, test.expectedError)
			}

			mux, _ := app.GetChiMux(*handlers.NewHandlerManager(
				mockAuthService,
				mockUserAuthService,
				mockBoardService,
			),
				*cfg,
			)

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v2/auth/signup/", body)
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
				verifiedAuth, err := mockAuthService.VerifyAuth(ctx, generatedCookie)

				require.NoError(t, err)

				require.Equal(t, uint64(1), verifiedAuth, "User wasn't saved correctly")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful registration")
			}
		})
	}
}

func TestAuthHandler_LogOut(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		token          string
		password       string
		successful     bool
		hasCookie      bool
		expiredCookie  bool
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "Successful logout",
			token:          "existing session",
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "No cookie",
			successful:     false,
			hasCookie:      false,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Empty cookie",
			successful:     false,
			hasCookie:      true,
			expectedStatus: http.StatusUnauthorized,
			expectedError:  apperrors.ErrSessionNotFound,
		},
		{
			name:           "Expired cookie",
			successful:     false,
			hasCookie:      true,
			expiredCookie:  true,
			expectedStatus: http.StatusUnauthorized,
			expectedError:  apperrors.ErrSessionExpired,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			cfg, err := createConfig(envPath)
			require.Equal(t, nil, err)

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserAuthService := mock_service.NewMockIUserAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			body := bytes.NewReader([]byte(""))

			r := httptest.NewRequest("DELETE", "/api/v2/auth/logout/", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			if test.successful {
				cookie := &http.Cookie{
					Name:     "tabula_user",
					Value:    test.token,
					HttpOnly: true,
					SameSite: http.SameSiteLaxMode,
					Expires:  time.Now().Add(time.Duration(time.Hour)),
					Path:     "/api/v2/",
				}

				r.AddCookie(cookie)

				mockAuthService.
					EXPECT().
					VerifyAuth(gomock.Any(), test.token).
					Return(uint64(1), nil)

				mockAuthService.
					EXPECT().
					LogOut(gomock.Any(), test.token).
					Return(nil)
			} else if test.hasCookie {
				var duration time.Duration
				mockAuthService.
					EXPECT().
					VerifyAuth(gomock.Any(), test.token).
					Return(uint64(0), test.expectedError)

				if test.expiredCookie {
					duration = 0 * time.Hour
				} else {
					duration = 1 * time.Hour
				}

				cookie := &http.Cookie{
					Name:     "tabula_user",
					Value:    test.token,
					HttpOnly: true,
					SameSite: http.SameSiteLaxMode,
					Expires:  time.Now().Add(duration),
					Path:     "/api/v2/",
				}

				r.AddCookie(cookie)
			}

			mux, _ := app.GetChiMux(*handlers.NewHandlerManager(
				mockAuthService,
				mockUserAuthService,
				mockBoardService,
			),
				*cfg,
			)

			mux.ServeHTTP(w, r)

			status := w.Result().StatusCode

			require.EqualValuesf(t, test.expectedStatus, status,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))
		})
	}
}
