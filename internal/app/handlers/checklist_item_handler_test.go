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
	"server/mocks/mock_service"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createChecklistItemMux(
	mockChecklistItemService *mock_service.MockIChecklistItemService,
) (http.Handler, error) {
	ChecklistItemHandler := *handlers.NewChecklistItemHandler(mockChecklistItemService)

	mux := chi.NewRouter()
	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/checklist/", func(r chi.Router) {
			r.Route("/item", func(r chi.Router) {
				r.Post("/create/", ChecklistItemHandler.Create)
				r.Post("/edit/", ChecklistItemHandler.Update)
				r.Delete("/delete/", ChecklistItemHandler.Delete)
			})
		})
	})
	return mux, nil
}

func TestChecklistItemHandler_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		newChecklistItem     dto.NewChecklistItemInfo
		createdChecklistItem *dto.ChecklistItemInfo
		expectations         func(clis *mock_service.MockIChecklistItemService, args args) *http.Request
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
				newChecklistItem: dto.NewChecklistItemInfo{
					ChecklistID:  uint64(1),
					Name:         "Mock new checklistItem",
					ListPosition: uint64(0),
				},
				createdChecklistItem: &dto.ChecklistItemInfo{
					ID:           uint64(1),
					ChecklistID:  uint64(1),
					Name:         "Mock new checklistItem",
					ListPosition: uint64(0),
				},
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					clis.
						EXPECT().
						Create(gomock.Any(), args.newChecklistItem).
						Return(args.createdChecklistItem, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"checklist_id":%v, "name":"%s", "list_position":%v}`,
						args.newChecklistItem.ChecklistID, args.newChecklistItem.Name, args.newChecklistItem.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/item/create/", body).
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
				newChecklistItem: dto.NewChecklistItemInfo{
					ChecklistID:  uint64(1),
					Name:         "Mock new checklistItem",
					ListPosition: uint64(0),
				},
				createdChecklistItem: &dto.ChecklistItemInfo{
					ID:           uint64(1),
					ChecklistID:  uint64(1),
					Name:         "Mock new checklistItem",
					ListPosition: uint64(0),
				},
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/item/create/", body).
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
			name: "Bad request (could not create checklistItem)",
			args: args{
				newChecklistItem: dto.NewChecklistItemInfo{
					ChecklistID:  uint64(1),
					Name:         "Mock new checklistItem",
					ListPosition: uint64(0),
				},
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					clis.
						EXPECT().
						Create(gomock.Any(), args.newChecklistItem).
						Return(&dto.ChecklistItemInfo{}, apperrors.ErrChecklistItemNotCreated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"checklist_id":%v, "name":"%s", "list_position":%v}`,
						args.newChecklistItem.ChecklistID, args.newChecklistItem.Name, args.newChecklistItem.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/item/create/", body).
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

			mockChecklistItemService := mock_service.NewMockIChecklistItemService(ctrl)

			testRequest := tt.args.expectations(mockChecklistItemService, tt.args)

			mux, err := createChecklistItemMux(mockChecklistItemService)
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
				var jsonBody map[string]map[string]dto.ChecklistItemInfo
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"]["checklistItem"], *tt.args.createdChecklistItem),
					"Board in JSON response doesn't match the board returned by the service")
			}
		})
	}
}

func TestChecklistItemHandler_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		info         dto.UpdatedChecklistItemInfo
		expectations func(clis *mock_service.MockIChecklistItemService, args args) *http.Request
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
				info: dto.UpdatedChecklistItemInfo{
					ID:           uint64(1),
					Name:         "Mock updated checklistItem info",
					ListPosition: uint64(1),
				},
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					clis.
						EXPECT().
						Update(gomock.Any(), args.info).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "name":"%s", "list_position":%v}`,
						args.info.ID, args.info.Name, args.info.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/item/edit/", body).
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
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/item/edit/", body).
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
			name: "Bad request (could not update checklistItem)",
			args: args{
				info: dto.UpdatedChecklistItemInfo{
					ID:           uint64(1),
					Name:         "Mock updated checklistItem info",
					ListPosition: uint64(1),
				},
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					clis.
						EXPECT().
						Update(gomock.Any(), args.info).
						Return(apperrors.ErrChecklistItemNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "name":"%s", "list_position":%v}`,
						args.info.ID, args.info.Name, args.info.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/item/edit/", body).
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

			mockChecklistItemService := mock_service.NewMockIChecklistItemService(ctrl)

			testRequest := tt.args.expectations(mockChecklistItemService, tt.args)

			mux, err := createChecklistItemMux(mockChecklistItemService)
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

func TestChecklistItemHandler_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		checklistItemID dto.ChecklistItemID
		expectations    func(clis *mock_service.MockIChecklistItemService, args args) *http.Request
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
				checklistItemID: dto.ChecklistItemID{
					Value: uint64(1),
				},
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					clis.
						EXPECT().
						Delete(gomock.Any(), args.checklistItemID).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`,
						args.checklistItemID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/checklist/item/delete/", body).
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
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/checklist/item/delete/", body).
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
			name: "Bad request (could not delete checklistItem)",
			args: args{
				checklistItemID: dto.ChecklistItemID{
					Value: uint64(1),
				},
				expectations: func(clis *mock_service.MockIChecklistItemService, args args) *http.Request {
					clis.
						EXPECT().
						Delete(gomock.Any(), args.checklistItemID).
						Return(apperrors.ErrChecklistItemNotDeleted)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`,
						args.checklistItemID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/checklist/item/delete/", body).
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

			mockChecklistItemService := mock_service.NewMockIChecklistItemService(ctrl)

			testRequest := tt.args.expectations(mockChecklistItemService, tt.args)

			mux, err := createChecklistItemMux(mockChecklistItemService)
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

			// if !tt.wantErr {
			// 	responseBody := w.Body.Bytes()
			// 	var jsonBody map[string]dto.FullBoardResult
			// 	err = json.Unmarshal(responseBody, &jsonBody)
			// 	require.NoError(t, err, "Error unmarshaling response")
			// 	require.True(t, reflect.DeepEqual(jsonBody["body"], tt.args.resultBoard), "Board in JSON response doesn't match the board returned by the service")
			// }
		})
	}
}
