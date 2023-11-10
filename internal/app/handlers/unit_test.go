package handlers_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
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

func createConfig(envPath string) (*config.Config, error) {
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
			"SESSION_ID_LENGTH=32"+"\n"+
			"POSTGRES_USER='testuser'"+"\n"+
			"POSTGRES_PASSWORD='testpw'"+"\n"+
			"POSTGRES_DB='testdb'",
	)
	if err != nil {
		return nil, err
	}

	config, err := config.LoadConfig(envPath, configPath)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func createMux(mockAuthService *mock_service.MockIAuthService,
	mockUserService *mock_service.MockIUserService,
	mockBoardService *mock_service.MockIBoardService,
	mockWorkspaceService *mock_service.MockIWorkspaceService) (http.Handler, error) {

	mockHandlerManager := handlers.HandlerManager{
		UserHandler:      *handlers.NewUserHandler(mockUserService),
		WorkspaceHandler: *handlers.NewWorkspaceHandler(mockWorkspaceService),
		BoardHandler:     *handlers.NewBoardHandler(mockAuthService, mockBoardService),
	}

	cfg, err := createConfig(envPath)
	if err != nil {
		return nil, err
	}

	mux, _ := app.GetChiMux(mockHandlerManager, *cfg.CORS)
	return mux, nil
}

func TestUserHandler_LogIn(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		userID         dto.UserID
		email          string
		password       string
		successful     bool
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "Successful login",
			userID:         dto.UserID{Value: uint64(1)},
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

			authInfo := dto.AuthInfo{
				Email:    test.email,
				Password: test.password,
			}

			mockCheckPassword := mockUserService.EXPECT().CheckPassword(gomock.Any(), authInfo)
			if test.successful {
				mockCheckPassword.Return(
					&entities.User{
						ID:           test.userID.Value,
						Email:        test.email,
						PasswordHash: hashFromAuthInfo(authInfo),
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
				mockCheckPassword.Return(nil, test.expectedError)
			}

			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
			require.Equal(t, nil, err)

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
				userID, err := mockAuthService.VerifyAuth(ctx, dto.SessionToken{Value: generatedCookie})

				require.NoError(t, err)
				require.Equal(t, test.userID, userID, "Expected cookie wasn't found")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful authentication")
			}
		})
	}
}

func TestUserHandler_SignUp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		email          string
		password       string
		successful     bool
		newUserID      dto.UserID
		expectedStatus int
		expectedError  error
	}{
		{
			name:           "New user",
			email:          "newuser@email.com",
			password:       "100500600",
			successful:     true,
			expectedStatus: http.StatusOK,
			newUserID:      dto.UserID{Value: uint64(1)},
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

			authInfo := dto.AuthInfo{
				Email:    test.email,
				Password: test.password,
			}

			mockRegister := mockUserService.EXPECT().RegisterUser(gomock.Any(), authInfo)
			if test.successful {
				mockRegister.Return(
					&entities.User{
						ID:           test.newUserID.Value,
						Email:        test.email,
						PasswordHash: test.password,
					},
					nil,
				)
				mockAuthService.
					EXPECT().
					AuthUser(gomock.Any(), test.newUserID)
				mockAuthService.
					EXPECT().
					VerifyAuth(gomock.Any(), gomock.Any()).
					Return(test.newUserID, nil)
			} else {
				mockRegister.Return(nil, test.expectedError)
			}

			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
			require.Equal(t, nil, err)

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
				verifiedAuth, err := mockAuthService.VerifyAuth(ctx, dto.SessionToken{Value: generatedCookie})

				require.NoError(t, err)

				require.Equal(t, test.newUserID, verifiedAuth, "User wasn't saved correctly")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful registration")
			}
		})
	}
}

func TestUserHandler_LogOut(t *testing.T) {
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

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
					VerifyAuth(gomock.Any(), dto.SessionToken{Value: test.token}).
					Return(dto.UserID{Value: uint64(1)}, nil)

				mockAuthService.
					EXPECT().
					LogOut(gomock.Any(), dto.SessionToken{Value: test.token}).
					Return(nil)
			} else if test.hasCookie {
				var duration time.Duration
				mockAuthService.
					EXPECT().
					VerifyAuth(gomock.Any(), dto.SessionToken{Value: test.token}).
					Return(dto.UserID{Value: uint64(0)}, test.expectedError)

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

			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
			require.Equal(t, nil, err)

			mux.ServeHTTP(w, r)

			status := w.Result().StatusCode

			require.EqualValuesf(t, test.expectedStatus, status,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))
		})
	}
}

// func TestUserHandler_ChangePassword(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name           string
// 		token          string
// 		password       string
// 		successful     bool
// 		hasCookie      bool
// 		expiredCookie  bool
// 		expectedStatus int
// 		expectedError  error
// 	}{}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()

// 			ctrl := gomock.NewController(t)

// 			mockAuthService := mock_service.NewMockIAuthService(ctrl)
// 			mockUserService := mock_service.NewMockIUserService(ctrl)
// 			mockBoardService := mock_service.NewMockIBoardService(ctrl)
// 			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

// 			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
// 			require.Equal(t, nil, err)
// 		})
// 	}
// }

func TestUserHandler_ChangeProfile(t *testing.T) {
	t.Parallel()

	newDescription := "New description"
	newSurname := "New surname"

	tests := []struct {
		name           string
		userID         uint64
		token          string
		oldUserName    string
		newUserName    string
		oldUserSurname string
		newUserSurname *string
		newDescription *string
		successful     bool
		expiredCookie  bool
		expectedStatus int
		expectedError  apperrors.ErrorResponse
	}{
		{
			name:           "Successful update",
			userID:         1,
			token:          "session",
			oldUserName:    "Old name",
			newUserName:    "New name",
			newUserSurname: &newSurname,
			newDescription: &newDescription,
			successful:     true,
			expectedStatus: http.StatusOK,
		},
		// {
		// 	name:           "Unsuccessful update (Unauthorized)",
		// 	userID:         2,
		// 	successful:     false,
		// 	expectedStatus: http.StatusUnauthorized,
		// 	expectedError:  apperrors.GenericUnauthorizedResponse,
		// },
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
			require.Equal(t, nil, err)

			newProfileInfo := dto.UserProfileInfo{
				UserID:      test.userID,
				Name:        test.newUserName,
				Surname:     test.newUserSurname,
				Description: test.newDescription,
			}

			body := bytes.NewReader([]byte(
				fmt.Sprintf(`{"user_id":%d, "name":"%s", "surname":"%s", "description":"%s"}`,
					test.userID, test.newUserName, test.newUserSurname, test.newDescription),
			))

			r := httptest.NewRequest("POST", "/api/v2/user/edit", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mockUpdateP := mockUserService.EXPECT().UpdateProfile(gomock.Any(), newProfileInfo)

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
					VerifyAuth(gomock.Any(), dto.SessionToken{Value: test.token}).
					Return(dto.UserID{Value: test.userID}, nil)

				mockUserService.
					EXPECT().
					GetWithID(gomock.Any(), dto.UserID{Value: test.userID}).
					Return(
						&entities.User{
							ID: test.userID,
						},
						nil,
					)

				mockUpdateP.Return(nil)
			} else {
				mockUpdateP.Return(apperrors.ErrUserNotUpdated)
			}

			mux.ServeHTTP(w, r)

			status := w.Result().StatusCode

			require.EqualValuesf(t, test.expectedStatus, status,
				"Expected code %d (%s), received code %d (%s)",
				test.expectedStatus, http.StatusText(test.expectedStatus),
				w.Code, http.StatusText(w.Code))

			if test.successful {
				resultBody := w.Body.Bytes()
				expectedBody := dto.JSONResponse{
					Body: dto.JSONMap{},
				}
				err := json.Unmarshal(resultBody, &expectedBody)
				require.NoError(t, err)
			}
		})
	}
}

// func TestUserHandler_ChangeAvatar(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name           string
// 		token          string
// 		password       string
// 		successful     bool
// 		hasCookie      bool
// 		expiredCookie  bool
// 		expectedStatus int
// 		expectedError  error
// 	}{}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()

// 			ctrl := gomock.NewController(t)

// 			mockAuthService := mock_service.NewMockIAuthService(ctrl)
// 			mockUserService := mock_service.NewMockIUserService(ctrl)
// 			mockBoardService := mock_service.NewMockIBoardService(ctrl)
// 			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

// 			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
// 			require.Equal(t, nil, err)
// 		})
// 	}
// }
