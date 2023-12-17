package postgresql

import (
	"context"
	"database/sql/driver"
	"fmt"
	"regexp"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func TestPostgresTaskStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.NewTaskInfo
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				info: &dto.NewTaskInfo{
					ListID:       1,
					Name:         "sdfgsdfgsdfgsdfgsdfgsdf",
					ListPosition: 0,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.task").
						Columns(newTaskFields...).
						Values(args.info.ListID, args.info.Name, args.info.ListPosition).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.ListID,
							args.info.Name,
							args.info.ListPosition,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
							AddRow(1, time.Now()))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: &dto.NewTaskInfo{
					ListID:       1,
					Name:         "sdfgsdfgsdfgsdfgsdfgsdf",
					ListPosition: 0,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.task").
						Columns(newTaskFields...).
						Values(args.info.ListID, args.info.Name, args.info.ListPosition).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.ListID,
							args.info.Name,
							args.info.ListPosition,
						).
						WillReturnError(apperrors.ErrTaskNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrTaskNotCreated,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if _, err := s.Create(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_CheckAccess(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.CheckTaskAccessInfo
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				info: &dto.CheckTaskAccessInfo{
					UserID: 1,
					TaskID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.Select("count(*)").
						From("public.task_user").
						Where(sq.And{
							sq.Eq{"id_task": args.info.TaskID},
							sq.Eq{"id_user": args.info.UserID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.TaskID,
							args.info.UserID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"count"}).
							AddRow(10))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "query fail",
			args: args{
				info: &dto.CheckTaskAccessInfo{
					UserID: 1,
					TaskID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.Select("count(*)").
						From("public.task_user").
						Where(sq.And{
							sq.Eq{"id_task": args.info.TaskID},
							sq.Eq{"id_user": args.info.UserID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.TaskID,
							args.info.UserID,
						).
						WillReturnError(apperrors.ErrCouldNotGetUser)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetUser,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if _, err := s.CheckAccess(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.CheckAccess() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_RemoveUser(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.RemoveTaskUserInfo
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				info: &dto.RemoveTaskUserInfo{
					UserID: 1,
					TaskID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("task_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_task": args.info.TaskID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.UserID,
							args.info.TaskID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query failed",
			args: args{
				info: &dto.RemoveTaskUserInfo{
					UserID: 1,
					TaskID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("task_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_task": args.info.TaskID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.UserID,
							args.info.TaskID,
						).
						WillReturnError(apperrors.ErrCouldNotRemoveTaskUser)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRemoveTaskUser,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if err := s.RemoveUser(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.RemoveUser() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_AddUser(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.AddTaskUserInfo
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				info: &dto.AddTaskUserInfo{
					UserID: 1,
					TaskID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.task_user").
						Columns("id_task", "id_user").
						Values(args.info.TaskID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.TaskID,
							args.info.UserID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: &dto.AddTaskUserInfo{
					UserID: 1,
					TaskID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.task_user").
						Columns("id_task", "id_user").
						Values(args.info.TaskID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.TaskID,
							args.info.UserID,
						).
						WillReturnError(apperrors.ErrCouldNotAddTaskUser)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotAddTaskUser,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if err := s.AddUser(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.AddUser() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		id    *dto.TaskID
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				id: &dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.task").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Happy path",
			args: args{
				id: &dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.task").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnError(apperrors.ErrTaskNotDeleted)
				},
			},
			wantErr: true,
			err:     apperrors.ErrTaskNotDeleted,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if err := s.Delete(ctx, *tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.AddUser() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_Read(t *testing.T) {
	t.Parallel()
	type args struct {
		id    *dto.TaskID
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				id: &dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allTaskFields2...).
						From("public.task").
						Join("public.task_user ON public.task.id = public.task_user.id_task").
						Join("public.comment ON public.task.id = public.comment.id_task").
						Where(sq.Eq{"public.task.id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					userIDS := pq.StringArray([]string{})
					commentIDs := pq.StringArray([]string{})
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnRows(sqlmock.NewRows(allTaskFields2).
							AddRow(1, 1, time.Now(), "ame", "dsd", 1, time.Now(), time.Now(), userIDS, commentIDs))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				id: &dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allTaskFields2...).
						From("public.task").
						Join("public.task_user ON public.task.id = public.task_user.id_task").
						Join("public.comment ON public.task.id = public.comment.id_task").
						Where(sq.Eq{"public.task.id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnError(apperrors.ErrCouldNotGetTask)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetTask,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if _, err := s.Read(ctx, *tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Read() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_ReadMany(t *testing.T) {
	t.Parallel()
	type args struct {
		id    *dto.TaskIDs
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				id: &dto.TaskIDs{
					Values: []string{"1", "2"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allTaskFields...).
						From("public.task").
						LeftJoin("public.task_user ON public.task.id = public.task_user.id_task").
						LeftJoin("public.comment ON public.task.id = public.comment.id_task").
						LeftJoin("public.checklist ON public.task.id = public.checklist.id_task").
						Where(sq.Eq{"public.task.id": args.id.Values}).
						GroupBy("public.task.id", "public.task.id_list").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					userIDS := pq.StringArray([]string{})
					commentIDs := pq.StringArray([]string{})
					checklistIDs := pq.StringArray([]string{})
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							"1", "2",
						).
						WillReturnRows(sqlmock.NewRows(allTaskFields).
							AddRow(1, 1, time.Now(), "ame", "dsd", 1, time.Now(), time.Now(), userIDS, commentIDs, checklistIDs))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				id: &dto.TaskIDs{
					Values: []string{"1", "2"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allTaskFields...).
						From("public.task").
						LeftJoin("public.task_user ON public.task.id = public.task_user.id_task").
						LeftJoin("public.comment ON public.task.id = public.comment.id_task").
						LeftJoin("public.checklist ON public.task.id = public.checklist.id_task").
						Where(sq.Eq{"public.task.id": args.id.Values}).
						GroupBy("public.task.id", "public.task.id_list").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							"1", "2",
						).
						WillReturnError(apperrors.ErrCouldNotGetBoard)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetBoard,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if _, err := s.ReadMany(ctx, *tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.ReadMany() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_Update(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.UpdatedTaskInfo
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				info: &dto.UpdatedTaskInfo{
					ID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.task").
						Set("name", args.info.Name).
						Set("description", args.info.Description).
						Set("list_position", args.info.ListPosition).
						Set("task_start", &args.info.Start).
						Set("task_end", &args.info.End).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.ListPosition,
							args.info.Start,
							args.info.End,
							args.info.ID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: &dto.UpdatedTaskInfo{
					ID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.task").
						Set("name", args.info.Name).
						Set("description", args.info.Description).
						Set("list_position", args.info.ListPosition).
						Set("task_start", &args.info.Start).
						Set("task_end", &args.info.End).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.ListPosition,
							args.info.Start,
							args.info.End,
							args.info.ID,
						).
						WillReturnError(apperrors.ErrTaskNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrTaskNotUpdated,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if err := s.Update(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Update() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_Move(t *testing.T) {
	type args struct {
		info  *dto.TaskMoveInfo
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path, different lists",
			args: args{
				info: &dto.TaskMoveInfo{
					TaskID: 1,
					OldList: dto.TaskMoveListInfo{
						ListID:  1,
						TaskIDs: []uint64{1, 2, 3},
					},
					NewList: dto.TaskMoveListInfo{
						ListID:  2,
						TaskIDs: []uint64{1, 2, 3},
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					queryArgsFinal := []driver.Value{}
					caseBuilder := sq.Case()
					for i, id := range []uint64{1, 2, 3} {
						queryArgsFinal = append(queryArgsFinal, fmt.Sprintf("%v", id))
						caseBuilder = caseBuilder.
							When(sq.Eq{"id": fmt.Sprintf("%v", id)}, fmt.Sprintf("%v", i)).
							Else("list_position")
					}

					query, _, _ := sq.
						Update("public.task").
						Set("list_position", caseBuilder).
						Set("id_list", sq.Case().
							When(sq.Eq{"id": fmt.Sprintf("%v", args.info.TaskID)}, fmt.Sprintf("%v", args.info.NewList.ListID)).
							Else("id_list")).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					queryArgsFinal = append(queryArgsFinal, fmt.Sprintf("%v", args.info.TaskID))
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							queryArgsFinal...,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Happy path, same list",
			args: args{
				info: &dto.TaskMoveInfo{
					TaskID: 1,
					OldList: dto.TaskMoveListInfo{
						ListID:  1,
						TaskIDs: []uint64{1, 2, 3},
					},
					NewList: dto.TaskMoveListInfo{
						ListID:  1,
						TaskIDs: []uint64{1, 2, 3},
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					queryArgsFinal := []driver.Value{}
					caseBuilder := sq.Case()
					for i, id := range []uint64{1, 2, 3} {
						queryArgsFinal = append(queryArgsFinal, fmt.Sprintf("%v", id))
						caseBuilder = caseBuilder.
							When(sq.Eq{"id": fmt.Sprintf("%v", id)}, fmt.Sprintf("%v", i)).
							Else("list_position")
					}

					query, _, _ := sq.
						Update("public.task").
						Set("list_position", caseBuilder).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							queryArgsFinal...,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query not executed, same list",
			args: args{
				info: &dto.TaskMoveInfo{
					TaskID: 1,
					OldList: dto.TaskMoveListInfo{
						ListID:  1,
						TaskIDs: []uint64{1, 2, 3},
					},
					NewList: dto.TaskMoveListInfo{
						ListID:  1,
						TaskIDs: []uint64{1, 2, 3},
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					queryArgsFinal := []driver.Value{}
					caseBuilder := sq.Case()
					for i, id := range []uint64{1, 2, 3} {
						queryArgsFinal = append(queryArgsFinal, fmt.Sprintf("%v", id))
						caseBuilder = caseBuilder.
							When(sq.Eq{"id": fmt.Sprintf("%v", id)}, fmt.Sprintf("%v", i)).
							Else("list_position")
					}

					query, _, _ := sq.
						Update("public.task").
						Set("list_position", caseBuilder).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							queryArgsFinal...,
						).
						WillReturnError(apperrors.ErrCouldNotChangeTaskOrder)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotChangeTaskOrder,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			tt.args.query(mock, tt.args)

			s := NewTaskStorage(db)

			if err := s.Move(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Move() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
