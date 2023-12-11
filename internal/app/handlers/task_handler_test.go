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

func createTaskMux(
	mockTaskService *mock_service.MockITaskService,
) (http.Handler, error) {
	TaskHandler := *handlers.NewTaskHandler(mockTaskService)

	mux := chi.NewRouter()
	mux.Route("/api/v2", func(r chi.Router) {
		r.Route("/task", func(r chi.Router) {
			r.Post("/", TaskHandler.Read)
			r.Post("/create/", TaskHandler.Create)
			r.Post("/edit/", TaskHandler.Update)
			r.Route("/user", func(r chi.Router) {
				r.Post("/add/", TaskHandler.AddUser)
				r.Post("/remove/", TaskHandler.RemoveUser)
			})
			r.Delete("/delete/", TaskHandler.Delete)
		})
	})
	return mux, nil
}

func TestTaskHandler_Unit_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		session      dto.SessionToken
		newTask      dto.NewTaskInfo
		resultTask   entities.Task
		expectations func(bs *mock_service.MockITaskService, args args) *http.Request
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
				session: dto.SessionToken{
					ID: "Mock session",
				},
				newTask: dto.NewTaskInfo{
					Name:         "Mock new Task",
					ListID:       uint64(1),
					ListPosition: uint64(0),
				},
				resultTask: entities.Task{
					ID:          uint64(1),
					ListID:      uint64(1),
					DateCreated: time.Now().Round(time.Second).UTC(),
					Name:        "Mock new task",
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					info := dto.NewTaskInfo{
						Name:         args.newTask.Name,
						ListID:       args.newTask.ListID,
						ListPosition: args.newTask.ListPosition,
					}

					bs.
						EXPECT().
						Create(gomock.Any(), info).
						Return(&args.resultTask, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "list_id":%v, "list_position":%v}`,
						args.newTask.Name, args.newTask.ListID, args.newTask.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/create/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/task/create/", body).
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
			name: "Bad request (could not create Task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				newTask: dto.NewTaskInfo{
					Name:   "Mock new Task",
					ListID: uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{
						Name:     "tabula_user",
						Value:    args.session.ID,
						HttpOnly: true,
						SameSite: http.SameSiteLaxMode,
						Expires:  args.session.ExpirationDate,
						Path:     "/api/v2/",
					}

					info := dto.NewTaskInfo{
						Name:         args.newTask.Name,
						ListID:       args.newTask.ListID,
						ListPosition: args.newTask.ListPosition,
					}

					bs.
						EXPECT().
						Create(gomock.Any(), info).
						Return(&entities.Task{}, apperrors.ErrTaskNotCreated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "list_id":%v, "list_position":%v}`,
						args.newTask.Name, args.newTask.ListID, args.newTask.ListPosition)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/create/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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

			mockTaskService := mock_service.NewMockITaskService(ctrl)

			testRequest := tt.args.expectations(mockTaskService, tt.args)

			mux, err := createTaskMux(mockTaskService)
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
				var jsonBody map[string]map[string]entities.Task
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"]["task"], tt.args.resultTask), "Task in JSON response doesn't match the Task returned by the service")
			}
		})
	}
}

func TestTaskHandler_Unit_Read(t *testing.T) {
	t.Parallel()

	type args struct {
		session      dto.SessionToken
		TaskID       dto.TaskID
		task         *dto.SingleTaskInfo
		expectations func(bs *mock_service.MockITaskService, args args) *http.Request
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		expectedCode int
	}{
		{
			name: "Successful read",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				task: &dto.SingleTaskInfo{
					ID:          uint64(1),
					ListID:      uint64(1),
					DateCreated: time.Now().Round(time.Second).UTC(),
					Name:        "Mock task",
				},
				TaskID: dto.TaskID{Value: uint64(1)},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Read(gomock.Any(), args.TaskID).
						Return(args.task, nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`, args.TaskID.Value)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/task/", body).
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
			name: "Bad request (could not read task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				TaskID: dto.TaskID{Value: uint64(1)},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Read(gomock.Any(), args.TaskID).
						Return(&dto.SingleTaskInfo{}, apperrors.ErrCouldNotGetTask)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`, args.TaskID.Value)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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

			mockTaskService := mock_service.NewMockITaskService(ctrl)

			testRequest := tt.args.expectations(mockTaskService, tt.args)

			mux, err := createTaskMux(mockTaskService)
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
				var jsonBody map[string]map[string]dto.SingleTaskInfo
				err = json.Unmarshal(responseBody, &jsonBody)
				require.NoError(t, err, "Error unmarshaling response")
				require.True(t, reflect.DeepEqual(jsonBody["body"]["task"], *tt.args.task), "Task in JSON response doesn't match the task returned by the service")
			}
		})
	}
}

func TestTaskHandler_Unit_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		session      dto.SessionToken
		updatedTask  dto.UpdatedTaskInfo
		expectations func(bs *mock_service.MockITaskService, args args) *http.Request
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
				updatedTask: dto.UpdatedTaskInfo{
					Name: "Mock new Task name",
					ID:   uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Update(gomock.Any(), args.updatedTask).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "id":%v}`, args.updatedTask.Name, args.updatedTask.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/edit/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/task/edit/", body).
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
			name: "Bad request (could not update task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				updatedTask: dto.UpdatedTaskInfo{
					Name: "Mock new Task name",
					ID:   uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Update(gomock.Any(), args.updatedTask).
						Return(apperrors.ErrTaskNotUpdated)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s", "id":%v}`, args.updatedTask.Name, args.updatedTask.ID)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/edit/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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

			mockTaskService := mock_service.NewMockITaskService(ctrl)

			testRequest := tt.args.expectations(mockTaskService, tt.args)

			mux, err := createTaskMux(mockTaskService)
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

func TestTaskHandler_Unit_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		session      dto.SessionToken
		TaskID       dto.TaskID
		expectations func(bs *mock_service.MockITaskService, args args) *http.Request
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
				TaskID: dto.TaskID{Value: uint64(1)},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Delete(gomock.Any(), args.TaskID).
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`, args.TaskID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/task/delete/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("DELETE", "/api/v2/task/delete/", body).
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
			name: "Bad request (could not delete task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				TaskID: dto.TaskID{Value: uint64(1)},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Delete(gomock.Any(), args.TaskID).
						Return(apperrors.ErrTaskNotDeleted)

					body := bytes.NewReader([]byte(fmt.Sprintf(`{"id":%v}`, args.TaskID.Value)))

					r := httptest.
						NewRequest("DELETE", "/api/v2/task/delete/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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

			mockTaskService := mock_service.NewMockITaskService(ctrl)

			testRequest := tt.args.expectations(mockTaskService, tt.args)

			mux, err := createTaskMux(mockTaskService)
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

func TestTaskHandler_Unit_AddUser(t *testing.T) {
	t.Parallel()

	type args struct {
		session      dto.SessionToken
		request      dto.AddTaskUserInfo
		expectations func(bs *mock_service.MockITaskService, args args) *http.Request
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
				request: dto.AddTaskUserInfo{
					UserID: uint64(1),
					TaskID: uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Return(nil)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "task_id":%v}`,
						args.request.UserID, args.request.TaskID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/add/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/add/", body).
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
			name: "Bad request (user already in task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				request: dto.AddTaskUserInfo{
					UserID: uint64(1),
					TaskID: uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Return(apperrors.ErrUserAlreadyInTask)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "task_id":%v}`,
						args.request.UserID, args.request.TaskID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/add/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusConflict,
		},
		{
			name: "Bad request (could not add user to task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				request: dto.AddTaskUserInfo{
					UserID: uint64(1),
					TaskID: uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Return(apperrors.ErrCouldNotAddTaskUser)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "task_id":%v}`,
						args.request.UserID, args.request.TaskID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/add/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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

			mockTaskService := mock_service.NewMockITaskService(ctrl)

			testRequest := tt.args.expectations(mockTaskService, tt.args)

			mux, err := createTaskMux(mockTaskService)
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

func TestTaskHandler_Unit_RemoveUser(t *testing.T) {
	t.Parallel()

	type args struct {
		session      dto.SessionToken
		info         dto.RemoveTaskUserInfo
		expectations func(bs *mock_service.MockITaskService, args args) *http.Request
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
				info: dto.RemoveTaskUserInfo{
					UserID: uint64(2),
					TaskID: uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						`{"user_id":%v, "task_id":%v}`,
						args.info.UserID, args.info.TaskID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/remove/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
					cookie := &http.Cookie{}

					body := bytes.NewReader([]byte(""))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/remove/", body).
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
			name: "Bad request (user not in task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				info: dto.RemoveTaskUserInfo{
					UserID: uint64(2),
					TaskID: uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Return(apperrors.ErrUserNotInTask)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "task_id":%v}`,
						args.info.UserID, args.info.TaskID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/remove/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
						)
					r.AddCookie(cookie)

					return r
				},
			},
			wantErr:      true,
			expectedCode: http.StatusConflict,
		},
		{
			name: "Bad request (could not remove user from task)",
			args: args{
				session: dto.SessionToken{
					ID: "Mock session",
				},
				info: dto.RemoveTaskUserInfo{
					UserID: uint64(2),
					TaskID: uint64(1),
				},
				expectations: func(bs *mock_service.MockITaskService, args args) *http.Request {
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
						Return(apperrors.ErrCouldNotRemoveTaskUser)

					body := bytes.NewReader([]byte(fmt.Sprintf(
						`{"user_id":%v, "task_id":%v}`,
						args.info.UserID, args.info.TaskID,
					)))

					r := httptest.
						NewRequest("POST", "/api/v2/task/user/remove/", body).
						WithContext(
							context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
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

			mockTaskService := mock_service.NewMockITaskService(ctrl)

			testRequest := tt.args.expectations(mockTaskService, tt.args)

			mux, err := createTaskMux(mockTaskService)
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
