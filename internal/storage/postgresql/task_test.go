package postgresql

import (
	"context"
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
						Join("public.tag_task ON public.task.id = public.tag_task.id_task").
						Where(sq.Eq{"public.task.id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnRows(sqlmock.NewRows(allTaskFields2).
							AddRow(1, 1, time.Now(), "ame", "dsd", 1, time.Now(), time.Now(), pq.StringArray([]string{}), pq.StringArray([]string{}), pq.StringArray([]string{})))
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
						Join("public.tag_task ON public.task.id = public.tag_task.id_task").
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
						LeftJoin("public.tag_task ON public.task.id = public.tag_task.id_task").
						Where(sq.Eq{"public.task.id": args.id.Values}).
						GroupBy("public.task.id", "public.task.id_list").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					userIDS := pq.StringArray([]string{})
					commentIDs := pq.StringArray([]string{})
					checklistIDs := pq.StringArray([]string{})
					tagIDs := pq.StringArray([]string{})
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							"1", "2",
						).
						WillReturnRows(sqlmock.NewRows(allTaskFields).
							AddRow(1, 1, time.Now(), "ame", "dsd", 1, time.Now(), time.Now(), userIDS, commentIDs, checklistIDs, tagIDs))
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
						LeftJoin("public.tag_task ON public.task.id = public.tag_task.id_task").
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
		{
			name: "Building query failed",
			args: args{
				id:    &dto.TaskIDs{},
				query: func(mock sqlmock.Sqlmock, args args) {},
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

func TestPostgresTaskStorage_AttachFile(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.AttachedFileInfo
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
				info: dto.AttachedFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
					DateCreated:  time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					fileID := 1

					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.file").
						Columns(allFileInfoFields...).
						Values(args.info.OriginalName, args.info.FilePath, args.info.DateCreated).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
							args.info.DateCreated,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(fileID),
						)

					query2, _, _ := sq.
						Insert("public.task_file").
						Columns("id_task", "id_file").
						Values(args.info.TaskID, fileID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.TaskID,
							fileID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Building query 1 failed",
			args: args{
				info:  dto.AttachedFileInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Beginning transaction failed",
			args: args{
				info: dto.AttachedFileInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin().WillReturnError(apperrors.ErrCouldNotBeginTransaction)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBeginTransaction,
		},
		{
			name: "Executing query 1 failed",
			args: args{
				info: dto.AttachedFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
					DateCreated:  time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.file").
						Columns(allFileInfoFields...).
						Values(args.info.OriginalName, args.info.FilePath, args.info.DateCreated).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
							args.info.DateCreated,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Executing query 1 and rollback failed",
			args: args{
				info: dto.AttachedFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
					DateCreated:  time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.file").
						Columns(allFileInfoFields...).
						Values(args.info.OriginalName, args.info.FilePath, args.info.DateCreated).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
							args.info.DateCreated,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
		},
		{
			name: "Executing query 2 failed",
			args: args{
				info: dto.AttachedFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
					DateCreated:  time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					fileID := 1

					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.file").
						Columns(allFileInfoFields...).
						Values(args.info.OriginalName, args.info.FilePath, args.info.DateCreated).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
							args.info.DateCreated,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(fileID),
						)

					query2, _, _ := sq.
						Insert("public.task_file").
						Columns("id_task", "id_file").
						Values(args.info.TaskID, fileID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.TaskID,
							fileID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Executing query 2 and rollback failed",
			args: args{
				info: dto.AttachedFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
					DateCreated:  time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					fileID := 1

					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.file").
						Columns(allFileInfoFields...).
						Values(args.info.OriginalName, args.info.FilePath, args.info.DateCreated).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
							args.info.DateCreated,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(fileID),
						)

					query2, _, _ := sq.
						Insert("public.task_file").
						Columns("id_task", "id_file").
						Values(args.info.TaskID, fileID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.TaskID,
							fileID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
		},
		{
			name: "Commiting failed",
			args: args{
				info: dto.AttachedFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
					DateCreated:  time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					fileID := 1

					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.file").
						Columns(allFileInfoFields...).
						Values(args.info.OriginalName, args.info.FilePath, args.info.DateCreated).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
							args.info.DateCreated,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(fileID),
						)

					query2, _, _ := sq.
						Insert("public.task_file").
						Columns("id_task", "id_file").
						Values(args.info.TaskID, fileID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.TaskID,
							fileID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit().WillReturnError(apperrors.ErrCouldNotCommit)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotCommit,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
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

			if err := s.AttachFile(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.AttachFile() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_RemoveFile(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.RemoveFileInfo
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
				info: dto.RemoveFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					fileID := 1

					query1, _, _ := sq.
						Select("id").
						From("public.file").
						Where(sq.And{
							sq.Eq{"name": args.info.OriginalName},
							sq.Eq{"filepath": args.info.FilePath},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(fileID),
						)

					query2, _, _ := sq.
						Delete("public.file").
						Where(sq.Eq{"id": fileID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							fileID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Building query 1 failed",
			args: args{
				info:  dto.RemoveFileInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Executing query 1 failed",
			args: args{
				info: dto.RemoveFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select("id").
						From("public.file").
						Where(sq.And{
							sq.Eq{"name": args.info.OriginalName},
							sq.Eq{"filepath": args.info.FilePath},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Executing query 2 failed",
			args: args{
				info: dto.RemoveFileInfo{
					TaskID:       1,
					OriginalName: "sdvsd",
					FilePath:     "sdvsd",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					fileID := 1

					query1, _, _ := sq.
						Select("id").
						From("public.file").
						Where(sq.And{
							sq.Eq{"name": args.info.OriginalName},
							sq.Eq{"filepath": args.info.FilePath},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.OriginalName,
							args.info.FilePath,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(fileID),
						)

					query2, _, _ := sq.
						Delete("public.file").
						Where(sq.Eq{"id": fileID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							fileID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
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

			if err := s.RemoveFile(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.RemoveFile() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTaskStorage_GetFileList(t *testing.T) {
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
						Select(allPublicFileInfoFields...).
						From("public.file").
						Join("public.task_file ON public.task_file.id_file = public.file.id").
						Where(sq.Eq{"public.task_file.id_task": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnRows(sqlmock.NewRows(allPublicFileInfoFields).
							AddRow("ame", "dsd", time.Now()))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Executing query failed",
			args: args{
				id: &dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allPublicFileInfoFields...).
						From("public.file").
						Join("public.task_file ON public.task_file.id_file = public.file.id").
						Where(sq.Eq{"public.task_file.id_task": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Building query failed",
			args: args{
				id:    &dto.TaskID{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Scanning rows failed",
			args: args{
				id: &dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allPublicFileInfoFields...).
						From("public.file").
						Join("public.task_file ON public.task_file.id_file = public.file.id").
						Where(sq.Eq{"public.task_file.id_task": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnRows(sqlmock.NewRows(allPublicFileInfoFields).
							AddRow(nil, nil, nil))
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotScanRows,
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

			if _, err := s.GetFileList(ctx, *tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.GetFileList() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
