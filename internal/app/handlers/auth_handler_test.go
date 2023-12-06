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
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/mocks/mock_service"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const configPath string = "../../../config/config.yml"
const envPath string = ""

func hashFromAuthInfo(info dto.AuthInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(info.Email + info.Password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func getLogger() logger.ILogger {
	logger, _ := logger.NewLogrusLogger(&config.LoggingConfig{
		Level:                  "debug",
		DisableTimestamp:       false,
		FullTimestamp:          true,
		LevelBasedReport:       true,
		DisableLevelTruncation: true,
		ReportCaller:           true,
	})
	return &logger
}

func createConfig(envPath string) (*config.Config, error) {
	cfgFile, err := os.Create(envPath + ".env")
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()

	_, err = fmt.Fprint(cfgFile,
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
	mockCSRFService *mock_service.MockICSRFService) (http.Handler, error) {

	mockHandlerManager := handlers.Handlers{
		AuthHandler: *handlers.NewAuthHandler(mockAuthService, mockUserService, mockCSRFService),
		UserHandler: *handlers.NewUserHandler(mockUserService),
	}

	cfg, err := createConfig(envPath)
	if err != nil {
		return nil, err
	}

	mux, _ := app.GetChiMux(mockHandlerManager, *cfg, getLogger())
	return mux, nil
}

func TestUserHandler_Unit_LogIn(t *testing.T) {
	t.Parallel()

	type args struct {
		userID       dto.UserID
		authInfo     dto.AuthInfo
		session      dto.SessionToken
		csrf         dto.CSRFData
		expectations func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful login",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "mock@mail.com",
					Password: "test1234",
				},
				userID: dto.UserID{Value: uint64(1)},
				session: dto.SessionToken{
					ID:             "Mock session",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)).Round(time.Second).UTC(),
				},
				csrf: dto.CSRFData{
					Token:          "Mock CSRF token",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)).Round(time.Second).UTC(),
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						CheckPassword(gomock.Any(), args.authInfo).
						Return(&entities.User{
							ID:           args.userID.Value,
							Email:        args.authInfo.Email,
							PasswordHash: hashFromAuthInfo(args.authInfo),
						},
							nil)

					as.
						EXPECT().
						AuthUser(gomock.Any(), args.userID).
						Return(args.session, nil)

					cs.
						EXPECT().
						SetupCSRF(gomock.Any(), args.userID).
						Return(args.csrf, nil)

					as.
						EXPECT().
						VerifyAuth(gomock.Any(), args.session).
						Return(args.userID, nil)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (bad JSON)",
			args: args{
				authInfo: dto.AuthInfo{},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					return ""
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (invalid info in JSON)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "notmail",
					Password: "short",
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (wrong password)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "mock@mail.com",
					Password: "wrongpw123",
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						CheckPassword(gomock.Any(), args.authInfo).
						Return(&entities.User{},
							apperrors.ErrWrongPassword)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Bad request (failed to create session)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "mock@mail.com",
					Password: "test1234",
				},
				userID: dto.UserID{Value: uint64(1)},
				session: dto.SessionToken{
					ID:             "Mock session",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)),
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						CheckPassword(gomock.Any(), args.authInfo).
						Return(&entities.User{
							ID:           args.userID.Value,
							Email:        args.authInfo.Email,
							PasswordHash: hashFromAuthInfo(args.authInfo),
						},
							nil)

					as.
						EXPECT().
						AuthUser(gomock.Any(), args.userID).
						Return(dto.SessionToken{}, apperrors.ErrSessionNotCreated)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Bad request (CSRF token not created)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "mock@mail.com",
					Password: "test1234",
				},
				userID: dto.UserID{Value: uint64(1)},
				session: dto.SessionToken{
					ID:             "Mock session",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)),
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						CheckPassword(gomock.Any(), args.authInfo).
						Return(&entities.User{
							ID:           args.userID.Value,
							Email:        args.authInfo.Email,
							PasswordHash: hashFromAuthInfo(args.authInfo),
						},
							nil)

					as.
						EXPECT().
						AuthUser(gomock.Any(), args.userID).
						Return(args.session, nil)

					cs.
						EXPECT().
						SetupCSRF(gomock.Any(), args.userID).
						Return(dto.CSRFData{}, apperrors.ErrCSRFNotCreated)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockCSRFService := mock_service.NewMockICSRFService(ctrl)

			requestString := tt.args.expectations(mockAuthService, mockUserService, mockCSRFService, tt.args)

			mux, err := createMux(mockAuthService, mockUserService, mockCSRFService)
			require.Equal(t, nil, err)

			body := bytes.NewReader([]byte(requestString))

			r := httptest.NewRequest("POST", "/api/v2/auth/login/", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			status := w.Result().StatusCode

			require.EqualValuesf(t, tt.expectedCode, status,
				"Expected code %d (%s), received code %d (%s)",
				tt.expectedCode, http.StatusText(tt.expectedCode),
				w.Code, http.StatusText(w.Code))

			if !tt.wantErr {
				require.Equal(t, w.Result().Cookies()[0].Name, "tabula_user", "Expected cookie wasn't found")
				generatedCookie := w.Result().Cookies()[0].Value
				expiresAt := w.Result().Cookies()[0].Expires.UTC()
				csrfToken := w.Header().Get("X-Csrf-Token")

				ctx := context.Background()
				userID, err := mockAuthService.VerifyAuth(ctx, dto.SessionToken{
					ID:             generatedCookie,
					ExpirationDate: expiresAt,
				})

				require.NoError(t, err)
				require.Equal(t, tt.args.userID, userID, "Expected cookie wasn't found")
				require.Equal(t, tt.args.csrf.Token, csrfToken, "CSRF wasn't set correctly")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful authentication")
			}
		})
	}
}

func TestUserHandler_SignUp(t *testing.T) {
	t.Parallel()

	type args struct {
		userID       dto.UserID
		authInfo     dto.AuthInfo
		session      dto.SessionToken
		csrf         dto.CSRFData
		expectations func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful signup",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "mock@mail.com",
					Password: "test1234",
				},
				userID: dto.UserID{Value: uint64(1)},
				session: dto.SessionToken{
					ID:             "Mock session",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)).Round(time.Second).UTC(),
				},
				csrf: dto.CSRFData{
					Token:          "Mock CSRF token",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)).Round(time.Second).UTC(),
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						RegisterUser(gomock.Any(), args.authInfo).
						Return(&entities.User{
							ID:           args.userID.Value,
							Email:        args.authInfo.Email,
							PasswordHash: hashFromAuthInfo(args.authInfo),
						},
							nil)

					as.
						EXPECT().
						AuthUser(gomock.Any(), args.userID).
						Return(args.session, nil)

					cs.
						EXPECT().
						SetupCSRF(gomock.Any(), args.userID).
						Return(args.csrf, nil)

					as.
						EXPECT().
						VerifyAuth(gomock.Any(), args.session).
						Return(args.userID, nil)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (bad JSON)",
			args: args{
				authInfo: dto.AuthInfo{},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					return ""
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (invalid info in JSON)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "notmail",
					Password: "short",
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (user already exists)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "exists@mail.com",
					Password: "wrongpw123",
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						RegisterUser(gomock.Any(), args.authInfo).
						Return(&entities.User{},
							apperrors.ErrUserAlreadyExists)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusConflict,
		},
		{
			name: "Bad request (failed to create session)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "mock@mail.com",
					Password: "test1234",
				},
				userID: dto.UserID{Value: uint64(1)},
				session: dto.SessionToken{
					ID:             "Mock session",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)),
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						RegisterUser(gomock.Any(), args.authInfo).
						Return(&entities.User{
							ID:           args.userID.Value,
							Email:        args.authInfo.Email,
							PasswordHash: hashFromAuthInfo(args.authInfo),
						},
							nil)

					as.
						EXPECT().
						AuthUser(gomock.Any(), args.userID).
						Return(dto.SessionToken{}, apperrors.ErrSessionNotCreated)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "Bad request (CSRF token not created)",
			args: args{
				authInfo: dto.AuthInfo{
					Email:    "mock@mail.com",
					Password: "test1234",
				},
				userID: dto.UserID{Value: uint64(1)},
				session: dto.SessionToken{
					ID:             "Mock session",
					ExpirationDate: time.Now().Add(time.Duration(14 * 24 * time.Hour)),
				},
				expectations: func(as *mock_service.MockIAuthService, us *mock_service.MockIUserService, cs *mock_service.MockICSRFService, args args) string {
					us.
						EXPECT().
						RegisterUser(gomock.Any(), args.authInfo).
						Return(&entities.User{
							ID:           args.userID.Value,
							Email:        args.authInfo.Email,
							PasswordHash: hashFromAuthInfo(args.authInfo),
						},
							nil)

					as.
						EXPECT().
						AuthUser(gomock.Any(), args.userID).
						Return(args.session, nil)

					cs.
						EXPECT().
						SetupCSRF(gomock.Any(), args.userID).
						Return(dto.CSRFData{}, apperrors.ErrCSRFNotCreated)

					return fmt.Sprintf(`{"email":"%s", "password":"%s"}`, args.authInfo.Email, args.authInfo.Password)
				},
			},
			wantErr:      true,
			expectedCode: http.StatusInternalServerError,
		},
		// {
		// 	name:           "User already exists",
		// 	email:          "test@email.com",
		// 	password:       "coolpassword",
		// 	successful:     false,
		// 	expectedCode: http.StatusConflict,
		// 	expectedError:  apperrors.ErrUserAlreadyExists,
		// },
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockCSRFService := mock_service.NewMockICSRFService(ctrl)

			requestString := tt.args.expectations(mockAuthService, mockUserService, mockCSRFService, tt.args)

			mux, err := createMux(mockAuthService, mockUserService, mockCSRFService)
			require.Equal(t, nil, err)

			body := bytes.NewReader([]byte(requestString))

			r := httptest.NewRequest("POST", "/api/v2/auth/signup/", body)
			r.Header.Add("Access-Control-Request-Headers", "content-type")
			r.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, r)

			status := w.Result().StatusCode

			require.EqualValuesf(t, tt.expectedCode, status,
				"Expected code %d (%s), received code %d (%s)",
				tt.expectedCode, http.StatusText(tt.expectedCode),
				w.Code, http.StatusText(w.Code))

			if !tt.wantErr {
				require.Equal(t, w.Result().Cookies()[0].Name, "tabula_user", "Expected cookie wasn't found")
				generatedCookie := w.Result().Cookies()[0].Value
				expiresAt := w.Result().Cookies()[0].Expires.UTC()
				csrfToken := w.Header().Get("X-Csrf-Token")

				ctx := context.Background()
				verifiedAuth, err := mockAuthService.VerifyAuth(ctx, dto.SessionToken{
					ID:             generatedCookie,
					ExpirationDate: expiresAt,
				})

				require.NoError(t, err)
				require.Equal(t, tt.args.userID, verifiedAuth, "User wasn't saved correctly")
				require.Equal(t, tt.args.csrf.Token, csrfToken, "CSRF wasn't set correctly")
			} else {
				require.Empty(t, w.Result().Cookies(), "Cookie was set despite unsuccessful registration")
			}
		})
	}
}

// func TestUserHandler_LogOut(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name           string
// 		token          string
// 		password       string
// 		successful     bool
// 		hasCookie      bool
// 		expiredCookie  bool
// 		expectedCode int
// 		expectedError  error
// 	}{
// 		{
// 			name:           "Successful logout",
// 			token:          "existing session",
// 			successful:     true,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:           "No cookie",
// 			successful:     false,
// 			hasCookie:      false,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:           "Empty cookie",
// 			successful:     false,
// 			hasCookie:      true,
// 			expectedCode: http.StatusUnauthorized,
// 			expectedError:  apperrors.ErrSessionNotFound,
// 		},
// 		{
// 			name:           "Expired cookie",
// 			successful:     false,
// 			hasCookie:      true,
// 			expiredCookie:  true,
// 			expectedCode: http.StatusUnauthorized,
// 			expectedError:  apperrors.ErrSessionExpired,
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			ctrl := gomock.NewController(t)

// 			mockAuthService := mock_service.NewMockIAuthService(ctrl)
// 			mockUserService := mock_service.NewMockIUserService(ctrl)
// 			mockBoardService := mock_service.NewMockIBoardService(ctrl)
// 			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

// 			body := bytes.NewReader([]byte(""))

// 			r := httptest.NewRequest("DELETE", "/api/v2/auth/logout/", body)
// 			r.Header.Add("Access-Control-Request-Headers", "content-type")
// 			r.Header.Add("Origin", "localhost:8081")
// 			w := httptest.NewRecorder()

// 			if !tt.wantErr {
// 				cookie := &http.Cookie{
// 					Name:     "tabula_user",
// 					Value:    tt.token,
// 					HttpOnly: true,
// 					SameSite: http.SameSiteLaxMode,
// 					Expires:  time.Now().Add(time.Duration(time.Hour)),
// 					Path:     "/api/v2/",
// 				}

// 				r.AddCookie(cookie)

// 				mockAuthService.
// 					EXPECT().
// 					VerifyAuth(gomock.Any(), dto.SessionToken{ID: tt.token}).
// 					Return(dto.UserID{Value: uint64(1)}, nil)

// 				mockAuthService.
// 					EXPECT().
// 					LogOut(gomock.Any(), dto.SessionToken{ID: tt.token}).
// 					Return(nil)
// 			} else if tt.hasCookie {
// 				var duration time.Duration
// 				mockAuthService.
// 					EXPECT().
// 					VerifyAuth(gomock.Any(), dto.SessionToken{ID: tt.token}).
// 					Return(dto.UserID{Value: uint64(time.Second)}, tt.expectedError)

// 				if tt.expiredCookie {
// 					duration = 0 * time.Hour
// 				} else {
// 					duration = 1 * time.Hour
// 				}

// 				cookie := &http.Cookie{
// 					Name:     "tabula_user",
// 					Value:    tt.token,
// 					HttpOnly: true,
// 					SameSite: http.SameSiteLaxMode,
// 					Expires:  time.Now().Add(duration),
// 					Path:     "/api/v2/",
// 				}

// 				r.AddCookie(cookie)
// 			}

// 			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
// 			require.Equal(t, nil, err)

// 			mux.ServeHTTP(w, r)

// 			status := w.Result().StatusCode

// 			require.EqualValuesf(t, tt.expectedCode, status,
// 				"Expected code %d (%s), received code %d (%s)",
// 				tt.expectedCode, http.StatusText(tt.expectedCode),
// 				w.Code, http.StatusText(w.Code))
// 		})
// 	}
// }

// func TestUserHandler_ChangePassword(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name           string
// 		token          string
// 		password       string
// 		successful     bool
// 		hasCookie      bool
// 		expiredCookie  bool
// 		expectedCode int
// 		expectedError  error
// 	}{}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(tt.name, func(t *testing.T) {
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

// func TestUserHandler_ChangeProfile(t *testing.T) {
// 	t.Parallel()

// 	newDescription := "New description"
// 	newSurname := "New surname"

// 	tests := []struct {
// 		name           string
// 		userID         uint64
// 		token          string
// 		oldUserName    string
// 		newUserName    string
// 		oldUserSurname string
// 		newUserSurname string
// 		newDescription string
// 		successful     bool
// 		expiredCookie  bool
// 		expectedCode int
// 		expectedError  apperrors.ErrorResponse
// 	}{
// 		{
// 			name:           "Successful update",
// 			userID:         1,
// 			token:          "session",
// 			oldUserName:    "Old name",
// 			newUserName:    "New name",
// 			newUserSurname: newSurname,
// 			newDescription: newDescription,
// 			successful:     true,
// 			expectedCode: http.StatusOK,
// 		},
// 		// {
// 		// 	name:           "Unsuccessful update (Unauthorized)",
// 		// 	userID:         2,
// 		// 	successful:     false,
// 		// 	expectedCode: http.StatusUnauthorized,
// 		// 	expectedError:  apperrors.GenericUnauthorizedResponse,
// 		// },
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			ctrl := gomock.NewController(t)

// 			mockAuthService := mock_service.NewMockIAuthService(ctrl)
// 			mockUserService := mock_service.NewMockIUserService(ctrl)
// 			mockBoardService := mock_service.NewMockIBoardService(ctrl)
// 			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

// 			mux, err := createMux(mockAuthService, mockUserService, mockBoardService, mockWorkspaceService)
// 			require.Equal(t, nil, err)

// 			newProfileInfo := dto.UserProfileInfo{
// 				UserID:      tt.userID,
// 				Name:        tt.newUserName,
// 				Surname:     tt.newUserSurname,
// 				Description: tt.newDescription,
// 			}

// 			body := bytes.NewReader([]byte(
// 				fmt.Sprintf(`{"user_id":%d, "name":"%s", "surname":"%s", "description":"%s"}`,
// 					tt.userID, tt.newUserName, tt.newUserSurname, tt.newDescription),
// 			))

// 			r := httptest.NewRequest("POST", "/api/v2/user/edit", body)
// 			r.Header.Add("Access-Control-Request-Headers", "content-type")
// 			r.Header.Add("Origin", "localhost:8081")
// 			w := httptest.NewRecorder()

// 			mockUpdateP := mockUserService.EXPECT().UpdateProfile(gomock.Any(), newProfileInfo)

// 			if !tt.wantErr {
// 				cookie := &http.Cookie{
// 					Name:     "tabula_user",
// 					Value:    tt.token,
// 					HttpOnly: true,
// 					SameSite: http.SameSiteLaxMode,
// 					Expires:  time.Now().Add(time.Duration(time.Hour)),
// 					Path:     "/api/v2/",
// 				}

// 				r.AddCookie(cookie)

// 				mockAuthService.
// 					EXPECT().
// 					VerifyAuth(gomock.Any(), dto.SessionToken{ID: tt.token}).
// 					Return(dto.UserID{Value: tt.userID}, nil)

// 				mockUserService.
// 					EXPECT().
// 					GetWithID(gomock.Any(), dto.UserID{Value: tt.userID}).
// 					Return(
// 						&entities.User{
// 							ID: tt.userID,
// 						},
// 						nil,
// 					)

// 				mockUpdateP.Return(nil)
// 			} else {
// 				mockUpdateP.Return(apperrors.ErrUserNotUpdated)
// 			}

// 			mux.ServeHTTP(w, r)

// 			status := w.Result().StatusCode

// 			require.EqualValuesf(t, tt.expectedCode, status,
// 				"Expected code %d (%s), received code %d (%s)",
// 				tt.expectedCode, http.StatusText(tt.expectedCode),
// 				w.Code, http.StatusText(w.Code))

// 			if !tt.wantErr {
// 				resultBody := w.Body.Bytes()
// 				expectedBody := dto.JSONResponse{
// 					Body: dto.JSONMap{},
// 				}
// 				err := json.Unmarshal(resultBody, &expectedBody)
// 				require.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestUserHandler_ChangeAvatar(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name           string
// 		token          string
// 		password       string
// 		successful     bool
// 		hasCookie      bool
// 		expiredCookie  bool
// 		expectedCode int
// 		expectedError  error
// 	}{}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(tt.name, func(t *testing.T) {
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
