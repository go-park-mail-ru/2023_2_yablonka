package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"server/internal/app/handlers"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/mocks/mock_service"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createBoardMux(
	mockAuthService *mock_service.MockIAuthService,
	mockBoardService *mock_service.MockIBoardService,
) (http.Handler, error) {
	BoardHandler := *handlers.NewBoardHandler(mockAuthService, mockBoardService)

	mux := chi.NewRouter()
	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/board", func(r chi.Router) {
			r.Post("/", BoardHandler.GetFullBoard)
			r.Post("/create/", BoardHandler.Create)
			r.Post("/update/", BoardHandler.UpdateData)
			r.Post("/update/change_thumbnail/", BoardHandler.UpdateThumbnail)
			r.Route("/user", func(r chi.Router) {
				r.Post("/add/", BoardHandler.AddUser)
				r.Post("/remove/", BoardHandler.RemoveUser)
			})
			r.Delete("/delete/", BoardHandler.Delete)
		})
	})
	return mux, nil
}

func TestBoardHandler_Unit_GetFullBoard(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		boardID      dto.BoardID
		resultBoard  dto.FullBoardResult
		expectations func(bs *mock_service.MockIBoardService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful get",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				boardID: dto.BoardID{
					Value: uint64(1),
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				resultBoard: dto.FullBoardResult{
					Board: dto.SingleBoardInfo{
						ID:               uint64(1),
						Name:             "Existing board return",
						WorkspaceID:      uint64(1),
						WorkspaceOwnerID: uint64(1),
						DateCreated:      time.Now().UTC(),
					},
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					requestedBoard := dto.IndividualBoardRequest{
						BoardID: args.boardID.Value,
						UserID:  args.user.ID,
					}

					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						GetFullBoard(gomock.Any(), requestedBoard).
						Return(&args.resultBoard, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v}`, args.boardID.Value)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/", body).
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
			name: "Bad request (unauthorized - no user object in context)",
			args: args{
				user: nil,
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v}`, args.boardID.Value)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/", body).
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
			name: "Bad request (invalid JSON)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
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
						NewRequest("POST", "/api/v2/board/", body).
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
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (board not found)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				boardID: dto.BoardID{
					Value: uint64(1),
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				resultBoard: dto.FullBoardResult{
					Board: dto.SingleBoardInfo{
						ID:               uint64(1),
						Name:             "Existing board return",
						WorkspaceID:      uint64(1),
						WorkspaceOwnerID: uint64(1),
						DateCreated:      time.Now().UTC(),
					},
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					requestedBoard := dto.IndividualBoardRequest{
						BoardID: args.boardID.Value,
						UserID:  args.user.ID,
					}

					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						GetFullBoard(gomock.Any(), requestedBoard).
						Return(&dto.FullBoardResult{}, apperrors.ErrCouldNotGetBoard)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v}`, args.boardID.Value)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/", body).
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
		{
			name: "Bad request (no access to board)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				boardID: dto.BoardID{
					Value: uint64(1),
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				resultBoard: dto.FullBoardResult{
					Board: dto.SingleBoardInfo{
						ID:               uint64(1),
						Name:             "Existing board return",
						WorkspaceID:      uint64(1),
						WorkspaceOwnerID: uint64(1),
						DateCreated:      time.Now().UTC(),
					},
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					requestedBoard := dto.IndividualBoardRequest{
						BoardID: args.boardID.Value,
						UserID:  args.user.ID,
					}

					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						GetFullBoard(gomock.Any(), requestedBoard).
						Return(&dto.FullBoardResult{}, apperrors.ErrNoBoardAccess)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v}`, args.boardID.Value)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/", body).
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
			expectedCode: http.StatusForbidden,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			testRequest := tt.args.expectations(mockBoardService, tt.args)

			mux, err := createBoardMux(mockAuthService, mockBoardService)
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

			if !tt.wantErr {
				responseBody := w.Body.Bytes()
				var jsonBody map[string]dto.FullBoardResult
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"], tt.args.resultBoard), "Board in JSON response doesn't match the board returned by the service")
			}
		})
	}
}

func TestBoardHandler_Unit_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		newBoard     dto.NewBoardRequest
		resultBoard  entities.Board
		expectations func(bs *mock_service.MockIBoardService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful create",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				newBoard: dto.NewBoardRequest{
					Name:        "Mock new board",
					WorkspaceID: uint64(1),
				},
				resultBoard: entities.Board{
					ID:          uint64(1),
					Name:        "Mock new board",
					Owner:       dto.UserID{Value: uint64(1)},
					DateCreated: time.Now().UTC(),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					info := dto.NewBoardInfo{
						Name:        args.newBoard.Name,
						WorkspaceID: args.newBoard.WorkspaceID,
						OwnerID:     args.user.ID,
						Thumbnail:   args.newBoard.Thumbnail,
					}

					bs.
						EXPECT().
						Create(gomock.Any(), info).
						Return(&args.resultBoard, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "workspace_id":%v}`, args.newBoard.Name, args.newBoard.WorkspaceID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/create/", body).
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
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/board/create/", body).
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
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized - no user object in context)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "workspace_id":%v}`, args.newBoard.Name, args.newBoard.WorkspaceID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/create/", body).
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
			name: "Bad request (could not create board)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				newBoard: dto.NewBoardRequest{
					Name:        "Mock new board",
					WorkspaceID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					info := dto.NewBoardInfo{
						Name:        args.newBoard.Name,
						WorkspaceID: args.newBoard.WorkspaceID,
						OwnerID:     args.user.ID,
						Thumbnail:   args.newBoard.Thumbnail,
					}

					bs.
						EXPECT().
						Create(gomock.Any(), info).
						Return(&entities.Board{}, apperrors.ErrBoardNotCreated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "workspace_id":%v}`, args.newBoard.Name, args.newBoard.WorkspaceID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/create/", body).
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			testRequest := tt.args.expectations(mockBoardService, tt.args)

			mux, err := createBoardMux(mockAuthService, mockBoardService)
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

			if !tt.wantErr {
				responseBody := w.Body.Bytes()
				var jsonBody map[string]map[string]entities.Board
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"]["board"], tt.args.resultBoard), "Board in JSON response doesn't match the board returned by the service")
			}
		})
	}
}

func TestBoardHandler_Unit_UpdateData(t *testing.T) {
	t.Parallel()

	type args struct {
		session      dto.SessionToken
		user         *entities.User
		updatedBoard dto.UpdatedBoardInfo
		expectations func(bs *mock_service.MockIBoardService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful update",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				updatedBoard: dto.UpdatedBoardInfo{
					Name: "Mock new board name",
					ID:   uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						UpdateData(gomock.Any(), args.updatedBoard).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "id":%v}`, args.updatedBoard.Name, args.updatedBoard.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/", body).
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
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/", body).
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
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized - no user object in context)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "d":%v}`, args.updatedBoard.Name, args.updatedBoard.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/", body).
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
			name: "Bad request (could not update board)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				updatedBoard: dto.UpdatedBoardInfo{
					Name: "Mock new board name",
					ID:   uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						UpdateData(gomock.Any(), args.updatedBoard).
						Return(apperrors.ErrBoardNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "id":%v}`, args.updatedBoard.Name, args.updatedBoard.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/", body).
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			testRequest := tt.args.expectations(mockBoardService, tt.args)

			mux, err := createBoardMux(mockAuthService, mockBoardService)
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

func TestBoardHandler_Unit_UpdateThumbnail(t *testing.T) {
	t.Parallel()

	type args struct {
		user             *entities.User
		session          dto.SessionToken
		updatedThumbnail dto.UpdatedBoardThumbnailInfo
		resultUrl        dto.UrlObj
		expectations     func(bs *mock_service.MockIBoardService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful update",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				updatedThumbnail: dto.UpdatedBoardThumbnailInfo{
					ID:        uint64(1),
					Thumbnail: []byte{},
				},
				resultUrl: dto.UrlObj{Value: "main_theme.jpg"},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						UpdateThumbnail(gomock.Any(), args.updatedThumbnail).
						Return(&args.resultUrl, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"thumbnail":%v, "id":%v}`, args.updatedThumbnail.Thumbnail, args.updatedThumbnail.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/change_thumbnail/", body).
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
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/change_thumbnail/", body).
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
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized - no user object in context)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"thumbnail":%v, "id":%v}`, args.updatedThumbnail.Thumbnail, args.updatedThumbnail.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/change_thumbnail/", body).
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
			name: "Bad request (could not update board)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				updatedThumbnail: dto.UpdatedBoardThumbnailInfo{
					ID:        uint64(1),
					Thumbnail: []byte{},
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						UpdateThumbnail(gomock.Any(), args.updatedThumbnail).
						Return(&dto.UrlObj{}, apperrors.ErrBoardNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"thumbnail":%v, "id":%v}`, args.updatedThumbnail.Thumbnail, args.updatedThumbnail.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/update/change_thumbnail/", body).
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			testRequest := tt.args.expectations(mockBoardService, tt.args)

			mux, err := createBoardMux(mockAuthService, mockBoardService)
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

func TestBoardHandler_Unit_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		boardID      dto.BoardID
		expectations func(bs *mock_service.MockIBoardService, args args) *http.Request
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
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				boardID: dto.BoardID{Value: uint64(1)},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						Delete(gomock.Any(), args.boardID).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v}`, args.boardID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/board/delete/", body).
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
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/board/delete/", body).
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
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized - no user object in context)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v}`, args.boardID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/board/delete/", body).
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
			name: "Bad request (could not delete board)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				boardID: dto.BoardID{Value: uint64(1)},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						Delete(gomock.Any(), args.boardID).
						Return(apperrors.ErrBoardNotDeleted)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v}`, args.boardID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/board/delete/", body).
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			testRequest := tt.args.expectations(mockBoardService, tt.args)

			mux, err := createBoardMux(mockAuthService, mockBoardService)
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

func TestBoardHandler_Unit_AddUser(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		request      dto.AddBoardUserRequest
		addedUser    dto.UserPublicInfo
		expectations func(bs *mock_service.MockIBoardService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful user add",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				request: dto.AddBoardUserRequest{
					UserEmail:   "seconduser@email.com",
					BoardID:     uint64(1),
					WorkspaceID: uint64(1),
				},
				addedUser: dto.UserPublicInfo{
					ID:    uint64(2),
					Email: "seconduser@email.com",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						AddUser(gomock.Any(), args.request).
						Return(args.addedUser, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_email":"%s", "board_id":%v, "workspace_id":%v}`,
						args.request.UserEmail, args.request.BoardID, args.request.WorkspaceID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/add/", body).
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
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/add/", body).
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
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized - no user object in context)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_email":"%s", "board_id":%v, "workspace_id":%v}`,
						args.request.UserEmail, args.request.BoardID, args.request.WorkspaceID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/add/", body).
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
			name: "Bad request (no board access)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				request: dto.AddBoardUserRequest{
					UserEmail:   "seconduser@email.com",
					BoardID:     uint64(1),
					WorkspaceID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						AddUser(gomock.Any(), args.request).
						Return(apperrors.ErrNoBoardAccess)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_email":"%s", "board_id":%v, "workspace_id":%v}`,
						args.request.UserEmail, args.request.BoardID, args.request.WorkspaceID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/add/", body).
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
			expectedCode: http.StatusForbidden,
		},
		{
			name: "Bad request (user already in board)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				request: dto.AddBoardUserRequest{
					UserEmail:   "seconduser@email.com",
					BoardID:     uint64(1),
					WorkspaceID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						AddUser(gomock.Any(), args.request).
						Return(apperrors.ErrUserAlreadyInBoard)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_email":"%s", "board_id":%v, "workspace_id":%v}`,
						args.request.UserEmail, args.request.BoardID, args.request.WorkspaceID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/add/", body).
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
			expectedCode: http.StatusConflict,
		},
		{
			name: "Bad request (could not add user to board)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				request: dto.AddBoardUserRequest{
					UserEmail:   "seconduser@email.com",
					BoardID:     uint64(1),
					WorkspaceID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						AddUser(gomock.Any(), args.request).
						Return(apperrors.ErrCouldNotAddBoardUser)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_email":"%s", "board_id":%v, "workspace_id":%v}`,
						args.request.UserEmail, args.request.BoardID, args.request.WorkspaceID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/add/", body).
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			testRequest := tt.args.expectations(mockBoardService, tt.args)

			mux, err := createBoardMux(mockAuthService, mockBoardService)
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

func TestBoardHandler_Unit_RemoveUser(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		info         dto.RemoveBoardUserInfo
		expectations func(bs *mock_service.MockIBoardService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful user remove",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.RemoveBoardUserInfo{
					UserID:  uint64(2),
					BoardID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						RemoveUser(gomock.Any(), args.info).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "board_id":%v}`,
						args.info.UserID, args.info.BoardID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/remove/", body).
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
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/remove/", body).
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
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (unauthorized - no user object in context)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "board_id":%v}`,
						args.info.UserID, args.info.BoardID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/remove/", body).
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
			name: "Bad request (no board access)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.RemoveBoardUserInfo{
					UserID:  uint64(2),
					BoardID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						RemoveUser(gomock.Any(), args.info).
						Return(apperrors.ErrNoBoardAccess)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "board_id":%v}`,
						args.info.UserID, args.info.BoardID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/remove/", body).
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
			expectedCode: http.StatusForbidden,
		},
		{
			name: "Bad request (user not in board)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.RemoveBoardUserInfo{
					UserID:  uint64(2),
					BoardID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						RemoveUser(gomock.Any(), args.info).
						Return(apperrors.ErrUserNotInBoard)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "board_id":%v}`,
						args.info.UserID, args.info.BoardID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/remove/", body).
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
			expectedCode: http.StatusConflict,
		},
		{
			name: "Bad request (could not remove user from board)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				info: dto.RemoveBoardUserInfo{
					UserID:  uint64(2),
					BoardID: uint64(1),
				},
				expectations: func(bs *mock_service.MockIBoardService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					bs.
						EXPECT().
						RemoveUser(gomock.Any(), args.info).
						Return(apperrors.ErrCouldNotRemoveBoardUser)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "board_id":%v}`,
						args.info.UserID, args.info.BoardID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/board/user/remove/", body).
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

			mockAuthService := mock_service.NewMockIAuthService(ctrl)
			mockBoardService := mock_service.NewMockIBoardService(ctrl)

			testRequest := tt.args.expectations(mockBoardService, tt.args)

			mux, err := createBoardMux(mockAuthService, mockBoardService)
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
