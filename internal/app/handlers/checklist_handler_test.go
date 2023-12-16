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
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createChecklistMux(
	mockChecklistService *mock_service.MockIChecklistService,
) (http.Handler, error) {
	ChecklistHandler := *handlers.NewChecklistHandler(mockChecklistService)

	mux := chi.NewRouter()
	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/checklist", func(r chi.Router) {
			r.Post("/create/", ChecklistHandler.Create)
			r.Post("/edit/", ChecklistHandler.Update)
			r.Delete("/delete/", ChecklistHandler.Delete)
		})
	})
	return mux, nil
}

func TestChecklistHandler_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		newChecklist     dto.NewChecklistInfo
		createdChecklist *dto.ChecklistInfo
		expectations     func(cls *mock_service.MockIChecklistService, args args) *http.Request
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
				newChecklist: dto.NewChecklistInfo{
					TaskID:       uint64(1),
					Name:         "Mock new checklist",
					ListPosition: uint64(0),
				},
				createdChecklist: &dto.ChecklistInfo{
					ID:           uint64(1),
					TaskID:       uint64(1),
					Name:         "Mock new checklist",
					ListPosition: uint64(0),
				},
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					cls.
						EXPECT().
						Create(gomock.Any(), args.newChecklist).
						Return(args.createdChecklist, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"task_id":%v, "name":"%s", "list_position":%v}`,
						args.newChecklist.TaskID, args.newChecklist.Name, args.newChecklist.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/create/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
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
				newChecklist: dto.NewChecklistInfo{
					TaskID:       uint64(1),
					Name:         "Mock new checklist",
					ListPosition: uint64(0),
				},
				createdChecklist: &dto.ChecklistInfo{
					ID:           uint64(1),
					TaskID:       uint64(1),
					Name:         "Mock new checklist",
					ListPosition: uint64(0),
				},
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/create/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (could not create checklist)",
			args: args{
				newChecklist: dto.NewChecklistInfo{
					TaskID:       uint64(1),
					Name:         "Mock new checklist",
					ListPosition: uint64(0),
				},
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					cls.
						EXPECT().
						Create(gomock.Any(), args.newChecklist).
						Return(&dto.ChecklistInfo{}, apperrors.ErrChecklistNotCreated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"task_id":%v, "name":"%s", "list_position":%v}`,
						args.newChecklist.TaskID, args.newChecklist.Name, args.newChecklist.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/create/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
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

			mockChecklistService := mock_service.NewMockIChecklistService(ctrl)

			testRequest := tt.args.expectations(mockChecklistService, tt.args)

			mux, err := createChecklistMux(mockChecklistService)
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
				var jsonBody map[string]map[string]dto.ChecklistInfo
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"]["checklist"], *tt.args.createdChecklist),
					"Board in JSON response doesn't match the board returned by the service")
			}
		})
	}
}

func TestChecklistHandler_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		info         dto.UpdatedChecklistInfo
		expectations func(cls *mock_service.MockIChecklistService, args args) *http.Request
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
				info: dto.UpdatedChecklistInfo{
					ID:           uint64(1),
					Name:         "Mock updated checklist info",
					ListPosition: uint64(1),
				},
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					cls.
						EXPECT().
						Update(gomock.Any(), args.info).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "name":"%s", "list_position":%v}`,
						args.info.ID, args.info.Name, args.info.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
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
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (could not update checklist)",
			args: args{
				info: dto.UpdatedChecklistInfo{
					ID:           uint64(1),
					Name:         "Mock updated checklist info",
					ListPosition: uint64(1),
				},
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					cls.
						EXPECT().
						Update(gomock.Any(), args.info).
						Return(apperrors.ErrChecklistNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "name":"%s", "list_position":%v}`,
						args.info.ID, args.info.Name, args.info.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/checklist/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
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

			mockChecklistService := mock_service.NewMockIChecklistService(ctrl)

			testRequest := tt.args.expectations(mockChecklistService, tt.args)

			mux, err := createChecklistMux(mockChecklistService)
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

func TestChecklistHandler_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		checklistID  dto.ChecklistID
		expectations func(cls *mock_service.MockIChecklistService, args args) *http.Request
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
				checklistID: dto.ChecklistID{
					Value: uint64(1),
				},
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					cls.
						EXPECT().
						Delete(gomock.Any(), args.checklistID).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`,
						args.checklistID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/checklist/delete/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
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
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/checklist/delete/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Bad request (could not delete checklist)",
			args: args{
				checklistID: dto.ChecklistID{
					Value: uint64(1),
				},
				expectations: func(cls *mock_service.MockIChecklistService, args args) *http.Request {
					cls.
						EXPECT().
						Delete(gomock.Any(), args.checklistID).
						Return(apperrors.ErrChecklistNotDeleted)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`,
						args.checklistID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/checklist/delete/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
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

			mockChecklistService := mock_service.NewMockIChecklistService(ctrl)

			testRequest := tt.args.expectations(mockChecklistService, tt.args)

			mux, err := createChecklistMux(mockChecklistService)
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
