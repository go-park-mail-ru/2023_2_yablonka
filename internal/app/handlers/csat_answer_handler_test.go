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
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createCSATAnswerMux(
	mockCSATAnswerService *mock_service.MockICSATSAnswerService,
	mockCSATQuestionService *mock_service.MockICSATQuestionService,
) (http.Handler, error) {
	CSATAnswerHandler := *handlers.NewCSATAnswerHandler(mockCSATAnswerService, mockCSATQuestionService)

	mux := chi.NewRouter()
	mux.Route("/csat", func(r chi.Router) {
		r.Post("/answer/", CSATAnswerHandler.Create)
	})
	return mux, nil
}

func TestCSATAnswerHandler_Unit_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		user         *entities.User
		session      dto.SessionToken
		newAnswer    dto.NewCSATAnswerInfo
		expectations func(as *mock_service.MockICSATSAnswerService, qs *mock_service.MockICSATQuestionService, args args) *http.Request
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
				newAnswer: dto.NewCSATAnswerInfo{
					ID:         uint64(1),
					QuestionID: uint64(1),
					Rating:     uint64(5),
				},
				expectations: func(as *mock_service.MockICSATSAnswerService, qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					CSATAnswer := dto.NewCSATAnswer{
						UserID:     args.user.ID,
						QuestionID: args.newAnswer.QuestionID,
						Rating:     args.newAnswer.Rating,
					}

					qs.
						EXPECT().
						CheckRating(gomock.Any(), args.newAnswer).
						Return(nil)

					as.
						EXPECT().
						Create(gomock.Any(), CSATAnswer).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "id_question":%v, "rating":%v}`,
						args.newAnswer.ID, args.newAnswer.QuestionID, args.newAnswer.Rating)))

					r := httptest.
						NewRequest("POST", "/csat/answer/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.UserObjKey, args.user,
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
				expectations: func(as *mock_service.MockICSATSAnswerService, qs *mock_service.MockICSATQuestionService, args args) *http.Request {
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
						NewRequest("POST", "/csat/answer/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
				newAnswer: dto.NewCSATAnswerInfo{
					ID:         uint64(1),
					QuestionID: uint64(1),
					Rating:     uint64(5),
				},
				expectations: func(as *mock_service.MockICSATSAnswerService, qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "id_question":%v, "rating":%v}`,
						args.newAnswer.ID, args.newAnswer.QuestionID, args.newAnswer.Rating)))

					r := httptest.
						NewRequest("POST", "/csat/answer/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Bad request (answer rating too big)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				newAnswer: dto.NewCSATAnswerInfo{
					ID:         uint64(1),
					QuestionID: uint64(1),
					Rating:     uint64(5),
				},
				expectations: func(as *mock_service.MockICSATSAnswerService, qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					qs.
						EXPECT().
						CheckRating(gomock.Any(), args.newAnswer).
						Return(apperrors.ErrAnswerRatingTooBig)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "id_question":%v, "rating":%v}`,
						args.newAnswer.ID, args.newAnswer.QuestionID, args.newAnswer.Rating)))

					r := httptest.
						NewRequest("POST", "/csat/answer/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.UserObjKey, args.user,
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
			name: "Bad request (could not store answer)",
			args: args{
				user: &entities.User{
					ID:           uint64(1),
					Email:        "mock@mail.com",
					PasswordHash: "Mock hash",
				},
				session: dto.SessionToken{
					ID: "Mock session",
				},
				newAnswer: dto.NewCSATAnswerInfo{
					ID:         uint64(1),
					QuestionID: uint64(1),
					Rating:     uint64(5),
				},
				expectations: func(as *mock_service.MockICSATSAnswerService, qs *mock_service.MockICSATQuestionService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					CSATAnswer := dto.NewCSATAnswer{
						UserID:     args.user.ID,
						QuestionID: args.newAnswer.QuestionID,
						Rating:     args.newAnswer.Rating,
					}

					qs.
						EXPECT().
						CheckRating(gomock.Any(), args.newAnswer).
						Return(nil)

					as.
						EXPECT().
						Create(gomock.Any(), CSATAnswer).
						Return(apperrors.ErrCouldNotStoreAnswer)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v, "id_question":%v, "rating":%v}`,
						args.newAnswer.ID, args.newAnswer.QuestionID, args.newAnswer.Rating)))

					r := httptest.
						NewRequest("POST", "/csat/answer/", body).
						WithContext(
							context.WithValue(
								context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
								dto.UserObjKey, args.user,
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

			mockCSATAnswerService := mock_service.NewMockICSATSAnswerService(ctrl)
			mockCSATQuestionService := mock_service.NewMockICSATQuestionService(ctrl)

			testRequest := tt.args.expectations(mockCSATAnswerService, mockCSATQuestionService, tt.args)

			mux, err := createCSATAnswerMux(mockCSATAnswerService, mockCSATQuestionService)
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
