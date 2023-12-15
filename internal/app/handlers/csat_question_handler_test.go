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
	"server/mocks/mock_service"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createCSATQuestionMux(
	mockCSATQuestionService *mock_service.MockICSATQuestionService,
) (http.Handler, error) {
	CSATQuestionHandler := *handlers.NewCSATQuestionHandler(mockCSATQuestionService)

	mux := chi.NewRouter()
	mux.Route("/csat", func(r chi.Router) {
		r.Route("/question", func(r chi.Router) {
			r.Get("/all", CSATQuestionHandler.GetQuestions)
			r.Get("/stats", CSATQuestionHandler.GetStats)
			r.Post("/create/", CSATQuestionHandler.Create)
			r.Post("/edit/", CSATQuestionHandler.Update)
		})
	})
	return mux, nil
}

func TestCSATQuestionHandler_Unit_GetStats(t *testing.T) {
	t.Parallel()

	type args struct {
		expectations func(cs *mock_service.MockICSATQuestionService, args args) *http.Request
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
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					r := httptest.
						NewRequest("GET", "/csat/question/stats", nil).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						GetStats(gomock.Any()).
						Return(&[]dto.QuestionWithStats{
							{
								ID:      uint64(1),
								Content: "Mock CSAT question 1",
								Type:    "5",
								Stats:   []dto.RatingStats{},
							},
							{
								ID:      uint64(2),
								Content: "Mock CSAT question 2",
								Type:    "5",
								Stats:   []dto.RatingStats{},
							},
						}, nil)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (could not get questions)",
			args: args{
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					r := httptest.
						NewRequest("GET", "/csat/question/stats", nil).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						GetStats(gomock.Any()).
						Return(&[]dto.QuestionWithStats{}, apperrors.ErrCouldNotGetQuestions)

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

			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

			testRequest := tt.args.expectations(mockCSATQuestionService, tt.args)

			mux, err := createCSATQuestionMux(mockCSATQuestionService)
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

func TestCSATQuestionHandler_Unit_GetQuestions(t *testing.T) {
	t.Parallel()

	type args struct {
		expectations func(cs *mock_service.MockICSATQuestionService, args args) *http.Request
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
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					r := httptest.
						NewRequest("GET", "/csat/question/all", nil).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						GetAll(gomock.Any()).
						Return(&[]dto.CSATQuestionFull{
							{
								ID:      uint64(1),
								Content: "Mock CSAT question 1",
								Type:    "5",
							},
							{
								ID:      uint64(2),
								Content: "Mock CSAT question 2",
								Type:    "5",
							},
						}, nil)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (could not get questions)",
			args: args{
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					r := httptest.
						NewRequest("GET", "/csat/question/all", nil).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						GetAll(gomock.Any()).
						Return(&[]dto.CSATQuestionFull{}, apperrors.ErrCouldNotGetQuestions)

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

			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

			testRequest := tt.args.expectations(mockCSATQuestionService, tt.args)

			mux, err := createCSATQuestionMux(mockCSATQuestionService)
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

func TestCSATQuestionHandler_Unit_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		newQuestion    dto.NewCSATQuestionInfo
		resultQuestion dto.CSATQuestionFull
		expectations   func(cs *mock_service.MockICSATQuestionService, args args) *http.Request
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
				newQuestion: dto.NewCSATQuestionInfo{
					Content: "Mock new question",
					Type:    "Mock new question type",
				},
				resultQuestion: dto.CSATQuestionFull{
					ID:      uint64(1),
					Content: "Mock new question",
					Type:    "Mock new question type",
				},
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					body := bytes.NewReader([]byte(fmt.Sprintf(`{"content":"%s", "type":"%s"}`,
						args.newQuestion.Content, args.newQuestion.Type)))

					r := httptest.
						NewRequest("POST", "/csat/question/create/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						Create(gomock.Any(), args.newQuestion).
						Return(&args.resultQuestion, nil)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/csat/question/create/", body).
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
			name: "Bad request (could not get questions)",
			args: args{
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					body := bytes.NewReader([]byte(fmt.Sprintf(`{"content":"%s", "type":"%s"}`,
						args.newQuestion.Content, args.newQuestion.Type)))

					r := httptest.
						NewRequest("POST", "/csat/question/create/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						Create(gomock.Any(), args.newQuestion).
						Return(&dto.CSATQuestionFull{}, apperrors.ErrCouldNotCreateQuestion)

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

			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

			testRequest := tt.args.expectations(mockCSATQuestionService, tt.args)

			mux, err := createCSATQuestionMux(mockCSATQuestionService)
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

func TestCSATQuestionHandler_Unit_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		updatedQuestion dto.UpdatedCSATQuestionInfo
		expectations    func(cs *mock_service.MockICSATQuestionService, args args) *http.Request
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
				updatedQuestion: dto.UpdatedCSATQuestionInfo{
					ID:      uint64(1),
					Content: "Mock updated question",
					Type:    "Mock updated question type",
				},
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "content":"%s", "type":"%s"}`,
						args.updatedQuestion.ID, args.updatedQuestion.Content, args.updatedQuestion.Type)))

					r := httptest.
						NewRequest("POST", "/csat/question/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						Update(gomock.Any(), args.updatedQuestion).
						Return(nil)

					return r
				},
			},
			wantErr:      false,
			expectedCode: http.StatusOK,
		},
		{
			name: "Bad request (invalid JSON)",
			args: args{
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/csat/question/edit/", body).
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
			name: "Bad request (could not get questions)",
			args: args{
				expectations: func(qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "content":"%s", "type":"%s"}`,
						args.updatedQuestion.ID, args.updatedQuestion.Content, args.updatedQuestion.Type)))

					r := httptest.
						NewRequest("POST", "/csat/question/edit/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.RequestIDKey, uuid.New(),
							),
						)

					qs.
						EXPECT().
						Update(gomock.Any(), args.updatedQuestion).
						Return(apperrors.ErrQuestionNotUpdated)

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

			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

			testRequest := tt.args.expectations(mockCSATQuestionService, tt.args)

			mux, err := createCSATQuestionMux(mockCSATQuestionService)
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
