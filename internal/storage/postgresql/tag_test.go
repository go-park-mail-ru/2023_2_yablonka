package postgresql

import (
	"context"
	"regexp"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func TestPostgresTagStorage_Create(t *testing.T) {
	t.Parallel()
	color := "dfbdfbdf"
	type args struct {
		info  dto.NewTagInfo
		query func(mock sqlmock.Sqlmock, args args)
		ctx   context.Context
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
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1),
						)

					query2, _, _ := sq.
						Insert("public.tag_board").
						Columns("id_tag", "id_board").
						Values(1, args.info.BoardID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query3, _, _ := sq.
						Insert("public.tag_task").
						Columns("id_tag", "id_task").
						Values(1, args.info.TaskID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query3)).
						WithArgs(
							1,
							args.info.TaskID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Building query failed",
			args: args{
				info: dto.NewTagInfo{},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Beginning transaction failed",
			args: args{
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
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
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnError(apperrors.ErrCouldNotScanRows)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotScanRows,
		},
		{
			name: "Executing query 1 and rollback failed",
			args: args{
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnError(apperrors.ErrCouldNotScanRows)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
		},
		{
			name: "Executing query 2 failed",
			args: args{
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1),
						)

					query2, _, _ := sq.
						Insert("public.tag_board").
						Columns("id_tag", "id_board").
						Values(1, args.info.BoardID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.BoardID,
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
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1),
						)

					query2, _, _ := sq.
						Insert("public.tag_board").
						Columns("id_tag", "id_board").
						Values(1, args.info.BoardID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.BoardID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
		},
		{
			name: "Executing query 3 failed",
			args: args{
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1),
						)

					query2, _, _ := sq.
						Insert("public.tag_board").
						Columns("id_tag", "id_board").
						Values(1, args.info.BoardID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query3, _, _ := sq.
						Insert("public.tag_task").
						Columns("id_tag", "id_task").
						Values(1, args.info.TaskID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query3)).
						WithArgs(
							1,
							args.info.TaskID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Executing query 3 and rollback failed",
			args: args{
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1),
						)

					query2, _, _ := sq.
						Insert("public.tag_board").
						Columns("id_tag", "id_board").
						Values(1, args.info.BoardID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query3, _, _ := sq.
						Insert("public.tag_task").
						Columns("id_tag", "id_task").
						Values(1, args.info.TaskID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query3)).
						WithArgs(
							1,
							args.info.TaskID,
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
				info: dto.NewTagInfo{
					TaskID:  1,
					BoardID: 1,
					Name:    "gfhfghfh",
					Color:   color,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Insert("public.tag").
						Columns("name", "color").
						Values(args.info.Name, args.info.Color).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Color,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1),
						)

					query2, _, _ := sq.
						Insert("public.tag_board").
						Columns("id_tag", "id_board").
						Values(1, args.info.BoardID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query3, _, _ := sq.
						Insert("public.tag_task").
						Columns("id_tag", "id_task").
						Values(1, args.info.TaskID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query3)).
						WithArgs(
							1,
							args.info.TaskID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit().WillReturnError(apperrors.ErrCouldNotCommit)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotCommit,
		},
		{
			name: "No logger found",
			args: args{
				info:  dto.NewTagInfo{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoLoggerFound,
		},
		{
			name: "No request ID found",
			args: args{
				info: dto.NewTagInfo{},
				ctx: context.WithValue(
					context.Background(),
					dto.LoggerKey, getLogger()),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoRequestIDFound,
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

			tt.args.query(mock, tt.args)

			s := NewTagStorage(db)

			if _, err := s.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTagStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTagStorage_Update(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.UpdatedTagInfo
		query func(mock sqlmock.Sqlmock, args args)
		ctx   context.Context
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
				info: dto.UpdatedTagInfo{
					ID:    0,
					Name:  "Bad mock name",
					Color: "dasdasd",
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Update("public.tag").
						Set("name", args.info.Name).
						Set("color", args.info.Color).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(
							args.info.Name,
							args.info.Color,
							args.info.ID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Building qury failed",
			args: args{
				info: dto.UpdatedTagInfo{},
				ctx:  context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Executing query failed",
			args: args{
				info: dto.UpdatedTagInfo{
					ID:    0,
					Name:  "Bad mock name",
					Color: "dasdasd",
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Update("public.tag").
						Set("name", args.info.Name).
						Set("color", args.info.Color).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(
							args.info.Name,
							args.info.Color,
							args.info.ID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "No logger found",
			args: args{
				info:  dto.UpdatedTagInfo{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoLoggerFound,
		},
		{
			name: "No request ID found",
			args: args{
				info: dto.UpdatedTagInfo{},
				ctx: context.WithValue(
					context.Background(),
					dto.LoggerKey, getLogger()),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoRequestIDFound,
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

			tt.args.query(mock, tt.args)

			s := NewTagStorage(db)

			if err := s.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTagStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTagStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.TagID
		query func(mock sqlmock.Sqlmock, args args)
		ctx   context.Context
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
				info: dto.TagID{
					Value: 0,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Delete("public.tag").
						Where(sq.Eq{"id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
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
			name: "Building query failed",
			args: args{
				info:  dto.TagID{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "No logger found",
			args: args{
				info:  dto.TagID{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoLoggerFound,
		},
		{
			name: "No request ID found",
			args: args{
				info: dto.TagID{},
				ctx: context.WithValue(
					context.Background(),
					dto.LoggerKey, getLogger()),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoRequestIDFound,
		},
		{
			name: "Executing query failed",
			args: args{
				info: dto.TagID{
					Value: 0,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Delete("public.tag").
						Where(sq.Eq{"id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(
							args.info.Value,
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
			t.Parallel()
			db, mock, err := sqlmock.New()

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.args.query(mock, tt.args)

			s := NewTagStorage(db)

			if err := s.Delete(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTagStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTagStorage_AddToTask(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.TagAndTaskIDs
		query func(mock sqlmock.Sqlmock, args args)
		ctx   context.Context
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
				info: dto.TagAndTaskIDs{
					TagID:  0,
					TaskID: 0,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Insert("public.tag_task").
						Columns("id_tag", "id_task").
						Values(args.info.TagID, args.info.TaskID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(
							args.info.TagID,
							args.info.TaskID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.TagAndTaskIDs{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "No logger found",
			args: args{
				info:  dto.TagAndTaskIDs{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoLoggerFound,
		},
		{
			name: "No request ID found",
			args: args{
				info: dto.TagAndTaskIDs{},
				ctx: context.WithValue(
					context.Background(),
					dto.LoggerKey, getLogger()),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoRequestIDFound,
		},
		{
			name: "Executing query failed",
			args: args{
				info: dto.TagAndTaskIDs{
					TagID:  0,
					TaskID: 0,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Insert("public.tag_task").
						Columns("id_tag", "id_task").
						Values(args.info.TagID, args.info.TaskID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(
							args.info.TagID,
							args.info.TaskID,
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
			t.Parallel()
			db, mock, err := sqlmock.New()

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.args.query(mock, tt.args)

			s := NewTagStorage(db)

			if err := s.AddToTask(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTagStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresTagStorage_RemoveFromTask(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.TagAndTaskIDs
		query func(mock sqlmock.Sqlmock, args args)
		ctx   context.Context
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
				info: dto.TagAndTaskIDs{
					TagID:  0,
					TaskID: 0,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Delete("public.tag_task").
						Where(sq.And{
							sq.Eq{"id_task": args.info.TaskID},
							sq.Eq{"id_tag": args.info.TagID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(
							args.info.TagID,
							args.info.TaskID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.TagAndTaskIDs{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "No logger found",
			args: args{
				info:  dto.TagAndTaskIDs{},
				ctx:   context.Background(),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoLoggerFound,
		},
		{
			name: "No request ID found",
			args: args{
				info: dto.TagAndTaskIDs{},
				ctx: context.WithValue(
					context.Background(),
					dto.LoggerKey, getLogger()),
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrNoRequestIDFound,
		},
		{
			name: "Executing query failed",
			args: args{
				info: dto.TagAndTaskIDs{
					TagID:  0,
					TaskID: 0,
				},
				ctx: context.WithValue(
					context.WithValue(
						context.Background(),
						dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Delete("public.tag_task").
						Where(sq.And{
							sq.Eq{"id_task": args.info.TaskID},
							sq.Eq{"id_tag": args.info.TagID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(
							args.info.TagID,
							args.info.TaskID,
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
			t.Parallel()
			db, mock, err := sqlmock.New()

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.args.query(mock, tt.args)

			s := NewTagStorage(db)

			if err := s.RemoveFromTask(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresTagStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
