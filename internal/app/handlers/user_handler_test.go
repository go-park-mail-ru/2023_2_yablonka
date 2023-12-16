package handlers_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/internal/app/handlers"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/mocks/mock_service"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createUserMux(
	mockUserService *mock_service.MockIUserService,
	mockWorkspaceService *mock_service.MockIWorkspaceService,
) (http.Handler, error) {
	UserHandler := *handlers.NewUserHandler(mockUserService)
	WorkspaceHandler := *handlers.NewWorkspaceHandler(mockWorkspaceService)

	mux := chi.NewRouter()
	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/workspaces", WorkspaceHandler.GetUserWorkspaces)
			r.Post("/edit/", UserHandler.ChangeProfile)
			r.Post("/edit/change_password/", UserHandler.ChangePassword)
			r.Post("/edit/change_avatar/", UserHandler.ChangeAvatar)
			r.Delete("/edit/delete_avatar/", UserHandler.DeleteAvatar)
		})
	})
	return mux, nil
}

func TestUserHandler_Unit_ChangePassword(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		passwords    dto.PasswordChangeInfo
		expectations func(bs *mock_service.MockIUserService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful change",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				passwords: dto.PasswordChangeInfo{
					UserID:      uint64(1),
					OldPassword: "Mock old password",
					NewPassword: "Mock new password",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						UpdatePassword(gomock.Any(), args.passwords).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"old_password":"%s", "new_password":"%s"}`,
						args.passwords.OldPassword, args.passwords.NewPassword)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_password/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_password/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (invalid data)",
			args: args{
				passwords: dto.PasswordChangeInfo{
					UserID:      uint64(1),
					OldPassword: "shortpw",
					NewPassword: "Mock new password that is definitely way over the string length valid limit",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"old_password":"%s", "new_password":"%s"}`,
						args.passwords.OldPassword, args.passwords.NewPassword)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_password/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized -- no user object in context)",
			args: args{
				passwords: dto.PasswordChangeInfo{
					UserID:      uint64(1),
					OldPassword: "Mock old password",
					NewPassword: "Mock new password",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"old_password":"%s", "new_password":"%s"}`,
						args.passwords.OldPassword, args.passwords.NewPassword)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_password/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Bad request (wrong password)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				passwords: dto.PasswordChangeInfo{
					UserID:      uint64(1),
					OldPassword: "Mock old password",
					NewPassword: "Mock new password",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						UpdatePassword(gomock.Any(), args.passwords).
						Return(apperrors.ErrWrongPassword)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"old_password":"%s", "new_password":"%s"}`,
						args.passwords.OldPassword, args.passwords.NewPassword)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_password/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Bad request (failed to update password)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				passwords: dto.PasswordChangeInfo{
					UserID:      uint64(1),
					OldPassword: "Mock old password",
					NewPassword: "Mock new password",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						UpdatePassword(gomock.Any(), args.passwords).
						Return(apperrors.ErrUserNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"old_password":"%s", "new_password":"%s"}`,
						args.passwords.OldPassword, args.passwords.NewPassword)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_password/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
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

			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

			testRequest := tt.args.expectations(mockUserService, tt.args)

			mux, err := createUserMux(mockUserService, mockWorkspaceService)
			require.Equal(t, nil, err)

			testRequest.Header.Add("Access-Control-Request-Headers", "content-type")
			testRequest.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, testRequest)

			status := w.Result().StatusCode

			require.EqualValuesf(t, tt.expectedCode, status,
				"Expected code %d (%s), received code %d (%s)",
				tt.expectedCode, http.StatusText(tt.expectedCode),
				w.Code, http.StatusText(w.Code))
		})
	}
}

func TestUserHandler_Unit_ChangeAvatar(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		info         dto.AvatarChangeInfo
		expectations func(bs *mock_service.MockIUserService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful change",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.AvatarChangeInfo{
					UserID:   uint64(1),
					Avatar:   []byte{},
					Filename: "mock_avatar.png",
					Mimetype: "image/png",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						UpdateAvatar(gomock.Any(), args.info).
						Return(&dto.UrlObj{
							Value: gomock.Any().String(),
						}, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"avatar":%v, "filename":"%s", "mimetype":"%s"}`,
						args.info.Avatar, args.info.Filename, args.info.Mimetype)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_avatar/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_password/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized -- no user object in context)",
			args: args{
				info: dto.AvatarChangeInfo{
					UserID:   uint64(1),
					Avatar:   []byte{},
					Filename: "mock_avatar.png",
					Mimetype: "image/png",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"avatar":%v, "filename":"%s", "mimetype":"%s"}`,
						args.info.Avatar, args.info.Filename, args.info.Mimetype)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_avatar/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Bad request (failed to change avatar)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.AvatarChangeInfo{
					UserID:   uint64(1),
					Avatar:   []byte{},
					Filename: "mock_avatar.png",
					Mimetype: "image/png",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						UpdateAvatar(gomock.Any(), args.info).
						Return(&dto.UrlObj{}, apperrors.ErrUserNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"avatar":%v, "filename":"%s", "mimetype":"%s"}`,
						args.info.Avatar, args.info.Filename, args.info.Mimetype)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/change_avatar/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
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

			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

			testRequest := tt.args.expectations(mockUserService, tt.args)

			mux, err := createUserMux(mockUserService, mockWorkspaceService)
			require.Equal(t, nil, err)

			testRequest.Header.Add("Access-Control-Request-Headers", "content-type")
			testRequest.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, testRequest)

			status := w.Result().StatusCode

			require.EqualValuesf(t, tt.expectedCode, status,
				"Expected code %d (%s), received code %d (%s)",
				tt.expectedCode, http.StatusText(tt.expectedCode),
				w.Code, http.StatusText(w.Code))
		})
	}
}

func TestUserHandler_Unit_ChangeProfile(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		info         dto.UserProfileInfo
		expectations func(bs *mock_service.MockIUserService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful change",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.UserProfileInfo{
					UserID:      uint64(1),
					Name:        "Mock new name",
					Surname:     "Mock new surname",
					Description: "Mock new description",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						UpdateProfile(gomock.Any(), args.info).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "surname":"%s", "description":"%s"}`,
						args.info.Name, args.info.Surname, args.info.Description)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized -- no user object in context)",
			args: args{
				info: dto.UserProfileInfo{
					UserID:      uint64(1),
					Name:        "Mock new name",
					Surname:     "Mock new surname",
					Description: "Mock new description",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "surname":"%s", "description":"%s"}`,
						args.info.Name, args.info.Surname, args.info.Description)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Bad request (failed to change profile)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.UserProfileInfo{
					UserID:      uint64(1),
					Name:        "Mock new name",
					Surname:     "Mock new surname",
					Description: "Mock new description",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						UpdateProfile(gomock.Any(), args.info).
						Return(apperrors.ErrUserNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "surname":"%s", "description":"%s"}`,
						args.info.Name, args.info.Surname, args.info.Description)))

					r := httptest.
						NewRequest("POST", "/api/v2/user/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
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

			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

			testRequest := tt.args.expectations(mockUserService, tt.args)

			mux, err := createUserMux(mockUserService, mockWorkspaceService)
			require.Equal(t, nil, err)

			testRequest.Header.Add("Access-Control-Request-Headers", "content-type")
			testRequest.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, testRequest)

			status := w.Result().StatusCode

			require.EqualValuesf(t, tt.expectedCode, status,
				"Expected code %d (%s), received code %d (%s)",
				tt.expectedCode, http.StatusText(tt.expectedCode),
				w.Code, http.StatusText(w.Code))
		})
	}
}

func TestUserHandler_Unit_DeleteAvatar(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		info         dto.AvatarRemovalInfo
		expectations func(bs *mock_service.MockIUserService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful delete",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
					AvatarURL: func() *string {
						url := "mock_url.jpg"
						return &url
					}(),
				},
				info: dto.AvatarRemovalInfo{
					UserID:    uint64(1),
					AvatarUrl: "mock_url.jpg",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						DeleteAvatar(gomock.Any(), args.info).
						Return(&dto.UrlObj{
							Value: gomock.Any().String(),
						}, nil)

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/user/edit/delete_avatar/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (unauthorized -- no user object in context)",
			args: args{
				info: dto.AvatarRemovalInfo{
					UserID:    uint64(1),
					AvatarUrl: "mockurl.jpg",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/user/edit/delete_avatar/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Bad request (failed to change avatar)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
					AvatarURL: func() *string {
						url := "mock_url.jpg"
						return &url
					}(),
				},
				info: dto.AvatarRemovalInfo{
					UserID:    uint64(1),
					AvatarUrl: "mock_url.jpg",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(us *mock_service.MockIUserService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					us.
						EXPECT().
						DeleteAvatar(gomock.Any(), args.info).
						Return(&dto.UrlObj{}, apperrors.ErrUserNotUpdated)

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/user/edit/delete_avatar/", body).
						WithContext(
							context.WithValue(
								context.WithValue(
									context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
									dto.UserObjKey, args.user,
								),
								dto.RequestIDKey, uuid.New(),
							),
						)
					r.AddCookie(cookie)

					return r
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

			mockUserService := mock_service.NewMockIUserService(ctrl)
			mockWorkspaceService := mock_service.NewMockIWorkspaceService(ctrl)

			testRequest := tt.args.expectations(mockUserService, tt.args)

			mux, err := createUserMux(mockUserService, mockWorkspaceService)
			require.Equal(t, nil, err)

			testRequest.Header.Add("Access-Control-Request-Headers", "content-type")
			testRequest.Header.Add("Origin", "localhost:8081")
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, testRequest)

			status := w.Result().StatusCode

			require.EqualValuesf(t, tt.expectedCode, status,
				"Expected code %d (%s), received code %d (%s)",
				tt.expectedCode, http.StatusText(tt.expectedCode),
				w.Code, http.StatusText(w.Code))
		})
	}
}
