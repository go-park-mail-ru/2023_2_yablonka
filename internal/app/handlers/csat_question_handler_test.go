package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"server/internal/app/handlers"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/mocks/mock_service"
	"testing"

	"github.com/go-chi/chi/v5"
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
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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

// func TestCSATQuestionHandler_Unit_GetQuestions(t *testing.T) {
// 	t.Parallel()

// 	type args struct {
// 		expectations func(cs *mock_service.MockICSATQuestionService, args args) *http.Request
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		wantErr      bool
// 		expectedCode int
// 	}{
// 		// TODO Add cases
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			ctrl := gomock.NewController(t)

// 			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

// 			testRequest := tt.args.expectations(mockCSATQuestionService, tt.args)

// 			mux, err := createCSATQuestionMux(mockCSATQuestionService)
// 			require.Equal(t, nil, err)

// 			testRequest.Header.Add("Access-Control-Request-Headers", "content-type")
// 			testRequest.Header.Add("Origin", "localhost:8081")
// 			w := httptest.NewRecorder()

// 			mux.ServeHTTP(w, testRequest)

// 			status := w.Result().StatusCode

// 			require.EqualValuesf(t, tt.expectedCode, status,
// 				"Expected code %d (%s), received code %d (%s)",
// 				tt.expectedCode, http.StatusText(tt.expectedCode),
// 				w.Code, http.StatusText(w.Code))
// 		})
// 	}
// }

// func TestCSATQuestionHandler_Unit_Create(t *testing.T) {
// 	t.Parallel()

// 	type args struct {
// 		expectations func(cs *mock_service.MockICSATQuestionService, args args) *http.Request
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		wantErr      bool
// 		expectedCode int
// 	}{
// 		// TODO Add cases
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			ctrl := gomock.NewController(t)

// 			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

// 			testRequest := tt.args.expectations(mockCSATQuestionService, tt.args)

// 			mux, err := createCSATQuestionMux(mockCSATQuestionService)
// 			require.Equal(t, nil, err)

// 			testRequest.Header.Add("Access-Control-Request-Headers", "content-type")
// 			testRequest.Header.Add("Origin", "localhost:8081")
// 			w := httptest.NewRecorder()

// 			mux.ServeHTTP(w, testRequest)

// 			status := w.Result().StatusCode

// 			require.EqualValuesf(t, tt.expectedCode, status,
// 				"Expected code %d (%s), received code %d (%s)",
// 				tt.expectedCode, http.StatusText(tt.expectedCode),
// 				w.Code, http.StatusText(w.Code))
// 		})
// 	}
// }

// func TestCSATQuestionHandler_Unit_Update(t *testing.T) {
// 	t.Parallel()

// 	type args struct {
// 		expectations func(cs *mock_service.MockICSATQuestionService, args args) *http.Request
// 	}
// 	tests := []struct {
// 		name         string
// 		args         args
// 		wantErr      bool
// 		expectedCode int
// 	}{
// 		// TODO Add cases
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			ctrl := gomock.NewController(t)

// 			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

// 			testRequest := tt.args.expectations(mockCSATQuestionService, tt.args)

// 			mux, err := createCSATQuestionMux(mockCSATQuestionService)
// 			require.Equal(t, nil, err)

// 			testRequest.Header.Add("Access-Control-Request-Headers", "content-type")
// 			testRequest.Header.Add("Origin", "localhost:8081")
// 			w := httptest.NewRecorder()

// 			mux.ServeHTTP(w, testRequest)

// 			status := w.Result().StatusCode

// 			require.EqualValuesf(t, tt.expectedCode, status,
// 				"Expected code %d (%s), received code %d (%s)",
// 				tt.expectedCode, http.StatusText(tt.expectedCode),
// 				w.Code, http.StatusText(w.Code))
// 		})
// 	}
// }
