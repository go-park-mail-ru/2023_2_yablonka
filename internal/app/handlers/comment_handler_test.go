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
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createCommentMux(
	mockCommentService *mock_service.MockICommentService,
) (http.Handler, error) {
	CommentHandler := *handlers.NewCommentHandler(mockCommentService)

	mux := chi.NewRouter()
	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/comment", func(r chi.Router) {
			r.Post("/create/", CommentHandler.Create)
		})
	})
	return mux, nil
}

func TestCommentHandler_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		newComment     dto.NewCommentInfo
		createdComment *entities.Comment
		expectations   func(cs *mock_service.MockICommentService, args args) *http.Request
	}
	tests := []struct {
		Text         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			Text: "Successful create",
			args: args{
				newComment: dto.NewCommentInfo{
					UserID: uint64(1),
					TaskID: uint64(1),
					Text:   "Mock new comment",
				},
				createdComment: &entities.Comment{
					ID:          uint64(1),
					UserID:      uint64(1),
					TaskID:      uint64(1),
					Text:        "Mock new comment",
					DateCreated: time.Now().Round(time.Second).UTC(),
				},
				expectations: func(cs *mock_service.MockICommentService, args args) *http.Request {
					cs.
						EXPECT().
						Create(gomock.Any(), args.newComment).
						Return(args.createdComment, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"user_id":%v, "text":"%s", "task_id":%v}`,
						args.newComment.UserID, args.newComment.Text, args.newComment.TaskID)))

					r := httptest.
						NewRequest("POST", "/api/v2/comment/create/", body).
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
			Text: "Bad request (invalid JSON)",
			args: args{
				newComment: dto.NewCommentInfo{
					UserID: uint64(1),
					TaskID: uint64(1),
					Text:   "Mock new comment",
				},
				expectations: func(cs *mock_service.MockICommentService, args args) *http.Request {
					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/comment/create/", body).
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
			Text: "Bad request (could not create Comment)",
			args: args{
				newComment: dto.NewCommentInfo{
					UserID: uint64(1),
					TaskID: uint64(1),
					Text:   "Mock new comment",
				},
				expectations: func(cs *mock_service.MockICommentService, args args) *http.Request {
					cs.
						EXPECT().
						Create(gomock.Any(), args.newComment).
						Return(&entities.Comment{}, apperrors.ErrCommentNotCreated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"user_id":%v, "text":"%s", "task_id":%v}`,
						args.newComment.UserID, args.newComment.Text, args.newComment.TaskID)))

					r := httptest.
						NewRequest("POST", "/api/v2/comment/create/", body).
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
		t.Run(tt.Text, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)

			mockCommentService := mock_service.NewMockICommentService(ctrl)

			testRequest := tt.args.expectations(mockCommentService, tt.args)

			mux, err := createCommentMux(mockCommentService)
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
				var jsonBody map[string]map[string]entities.Comment
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"]["comment"], *tt.args.createdComment),
					"Comment in JSON response doesn't match the comment returned by the service")
			}
		})
	}
}
