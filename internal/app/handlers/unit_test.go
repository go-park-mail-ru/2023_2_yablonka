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

func Test_ULogin(t *testing.T) {
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

			cfgFile, err := os.Create(envPath + ".env")
			require.NoError(t, err)
			defer cfgFile.Close()

			_, err = fmt.Fprint(cfgFile,
				"JWT_SECRET='test secret'"+"\n"+
					"SESSION_DURATION_DAYS=14"+"\n"+
					"SESSION_DURATION_HOURS=0"+"\n"+
					"SESSION_DURATION_MINUTES=0"+"\n"+
					"SESSION_DURATION_SECONDS=0"+"\n"+
					"SESSION_ID_LENGTH=32"+"\n",
			)
			require.NoError(t, err)

			config, err := config.NewBaseEnvConfig(envPath, configPath)
			require.Equal(t, nil, err)

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"email":"%s", "password":"%s"}`, test.email, test.password),
			))

			r := httptest.NewRequest("POST", "/api/v2/auth/login/", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

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
				mockUA.
					Return(
						&entities.User{
							ID:           test.userID,
							Email:        test.email,
							PasswordHash: test.password,
						},
						nil,
					).
					AnyTimes()
				mockAuthService.
					EXPECT().
					AuthUser(gomock.Any(), test.userID)
				mockAuthService.
					EXPECT().
					VerifyAuth(gomock.Any(), gomock.Any()).
					Return(test.userID, nil)
			} else {
				mockUA.
					Return(nil, test.expectedError).
					AnyTimes()
			}

			mux, _ := app.GetChiMux(*handlers.NewHandlerManager(
				mockAuthService,
				mockUserAuthService,
				mockBoardService,
			),
				*config,
			)

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
