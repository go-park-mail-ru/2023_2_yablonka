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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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

/*

func TestPostgresTaskStorage_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  dto.TaskID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresTaskStorage{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresTaskStorage_Read(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  dto.TaskID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.SingleTaskInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresTaskStorage{
				db: tt.fields.db,
			}
			got, err := s.Read(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresTaskStorage_ReadMany(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  dto.TaskIDs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.SingleTaskInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresTaskStorage{
				db: tt.fields.db,
			}
			got, err := s.ReadMany(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadMany() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresTaskStorage_RemoveUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.RemoveTaskUserInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresTaskStorage{
				db: tt.fields.db,
			}
			if err := s.RemoveUser(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresTaskStorage_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedTaskInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresTaskStorage{
				db: tt.fields.db,
			}
			if err := s.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

*/
