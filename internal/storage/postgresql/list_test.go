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
	"github.com/lib/pq"
)

func TestPostgresListStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.NewListInfo
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
				info: &dto.NewListInfo{
					BoardID:      1,
					Name:         "sdfgsdfgsdfgsdfgsdfgsdf",
					ListPosition: 0,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.list").
						Columns("name", "list_position", "description", "id_board").
						Values(args.info.Name, args.info.ListPosition, args.info.Description, args.info.BoardID).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.ListPosition,
							args.info.Description,
							args.info.BoardID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: &dto.NewListInfo{
					BoardID:      1,
					Name:         "sdfgsdfgsdfgsdfgsdfgsdf",
					ListPosition: 0,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.list").
						Columns("name", "list_position", "description", "id_board").
						Values(args.info.Name, args.info.ListPosition, args.info.Description, args.info.BoardID).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.ListPosition,
							args.info.Description,
							args.info.BoardID,
						).
						WillReturnError(apperrors.ErrListNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrListNotCreated,
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

			s := NewListStorage(db)

			if _, err := s.Create(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresListStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.ListID
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
				info: &dto.ListID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.list").
						Where(sq.Eq{"id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
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
				info: &dto.ListID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.list").
						Where(sq.Eq{"id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrListNotDeleted)
				},
			},
			wantErr: true,
			err:     apperrors.ErrListNotDeleted,
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

			s := NewListStorage(db)

			if err := s.Delete(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresListStorage_Update(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.UpdatedListInfo
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
				info: &dto.UpdatedListInfo{
					ID:   1,
					Name: "sdsd",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.list").
						Set("name", args.info.Name).
						Set("description", args.info.Description).
						Set("list_position", args.info.ListPosition).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.ListPosition,
							args.info.ID,
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
				info: &dto.UpdatedListInfo{
					ID:   1,
					Name: "sdsd",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.list").
						Set("name", args.info.Name).
						Set("description", args.info.Description).
						Set("list_position", args.info.ListPosition).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.ListPosition,
							args.info.ID,
						).
						WillReturnError(apperrors.ErrListNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrListNotUpdated,
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

			s := NewListStorage(db)

			if err := s.Update(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Update() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresListStorage_GetTasksWithID(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.ListIDs
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
				info: &dto.ListIDs{
					Values: []uint64{1, 2},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allTaskAggFields...).
						From("public.task").
						Where(sq.Eq{"public.task.id_list": args.info.Values}).
						LeftJoin("public.task_user ON public.task_user.id_task = public.task.id").
						GroupBy("public.task.id", "public.task.id_list").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							1, 2,
						).
						WillReturnRows(sqlmock.NewRows(allTaskAggFields).
							AddRow(1, 1, time.Now(), "ame", "dsd", 1, time.Now(), time.Now(), pq.StringArray([]string{})))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query failed",
			args: args{
				info: &dto.ListIDs{
					Values: []uint64{1, 2},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allTaskAggFields...).
						From("public.task").
						Where(sq.Eq{"public.task.id_list": args.info.Values}).
						LeftJoin("public.task_user ON public.task_user.id_task = public.task.id").
						GroupBy("public.task.id", "public.task.id_list").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							1, 2,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewListStorage(db)

			if _, err := s.GetTasksWithID(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStorage.Update() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
