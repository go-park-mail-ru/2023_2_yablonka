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

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createListMux(
	mockListService *mock_service.MockIListService,
) (http.Handler, error) {
	ListHandler := *handlers.NewListHandler(mockListService)

	mux := chi.NewRouter()
	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/list", func(r chi.Router) {
			r.Post("/create/", ListHandler.Create)
			r.Post("/edit/", ListHandler.Update)
			r.Delete("/delete/", ListHandler.Delete)
		})
	})
	return mux, nil
}

func TestListHandler_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		newList      dto.NewListInfo
		createdList  *entities.List
		expectations func(cls *mock_service.MockIListService, args args) *http.Request
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
				newList: dto.NewListInfo{
					BoardID:      uint64(1),
					Name:         "Mock new List",
					ListPosition: uint64(0),
				},
				createdList: &entities.List{
					ID:           uint64(1),
					BoardID:      uint64(1),
					Name:         "Mock new List",
					ListPosition: uint64(0),
				},
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					cls.
						EXPECT().
						Create(gomock.Any(), args.newList).
						Return(args.createdList, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v, "name":"%s", "list_position":%v}`,
						args.newList.BoardID, args.newList.Name, args.newList.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/list/create/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				newList: dto.NewListInfo{
					BoardID:      uint64(1),
					Name:         "Mock new List",
					ListPosition: uint64(0),
				},
				createdList: &entities.List{
					ID:           uint64(1),
					BoardID:      uint64(1),
					Name:         "Mock new List",
					ListPosition: uint64(0),
				},
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/list/create/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (could not create List)",
			args: args{
				newList: dto.NewListInfo{
					BoardID:      uint64(1),
					Name:         "Mock new List",
					ListPosition: uint64(0),
				},
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					cls.
						EXPECT().
						Create(gomock.Any(), args.newList).
						Return(&entities.List{}, apperrors.ErrListNotCreated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"board_id":%v, "name":"%s", "list_position":%v}`,
						args.newList.BoardID, args.newList.Name, args.newList.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/list/create/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

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

			mockListService := mock_service.NewMockIListService(ctrl)

			testRequest := tt.args.expectations(mockListService, tt.args)

			mux, err := createListMux(mockListService)
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
				var jsonBody map[string]map[string]entities.List
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"]["list"], *tt.args.createdList),
					"Board in JSON response doesn't match the board returned by the service")
			}
		})
	}
}

func TestListHandler_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		info         dto.UpdatedListInfo
		expectations func(cls *mock_service.MockIListService, args args) *http.Request
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
				info: dto.UpdatedListInfo{
					ID:           uint64(1),
					Name:         "Mock updated List info",
					ListPosition: uint64(1),
				},
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					cls.
						EXPECT().
						Update(gomock.Any(), args.info).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "name":"%s", "list_position":%v}`,
						args.info.ID, args.info.Name, args.info.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/list/edit/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/list/edit/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (could not update List)",
			args: args{
				info: dto.UpdatedListInfo{
					ID:           uint64(1),
					Name:         "Mock updated List info",
					ListPosition: uint64(1),
				},
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					cls.
						EXPECT().
						Update(gomock.Any(), args.info).
						Return(apperrors.ErrListNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "name":"%s", "list_position":%v}`,
						args.info.ID, args.info.Name, args.info.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/list/edit/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

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

			mockListService := mock_service.NewMockIListService(ctrl)

			testRequest := tt.args.expectations(mockListService, tt.args)

			mux, err := createListMux(mockListService)
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

func TestListHandler_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		ListID       dto.ListID
		expectations func(cls *mock_service.MockIListService, args args) *http.Request
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
				ListID: dto.ListID{
					Value: uint64(1),
				},
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					cls.
						EXPECT().
						Delete(gomock.Any(), args.ListID).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`,
						args.ListID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/list/delete/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/list/delete/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (could not delete List)",
			args: args{
				ListID: dto.ListID{
					Value: uint64(1),
				},
				expectations: func(cls *mock_service.MockIListService, args args) *http.Request {
					cls.
						EXPECT().
						Delete(gomock.Any(), args.ListID).
						Return(apperrors.ErrListNotDeleted)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`,
						args.ListID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/list/delete/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)

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

			mockListService := mock_service.NewMockIListService(ctrl)

			testRequest := tt.args.expectations(mockListService, tt.args)

			mux, err := createListMux(mockListService)
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
