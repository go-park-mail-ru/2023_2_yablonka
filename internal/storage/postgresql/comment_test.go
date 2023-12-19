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
)

func TestPostgresCommentStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewCommentInfo
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
				info: dto.NewCommentInfo{
					UserID: 1,
					TaskID: 1,
					Text:   "dfdfdfdf",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.comment").
						Columns(newCommentFields...).
						Values(args.info.TaskID, args.info.UserID, args.info.Text).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.TaskID,
							args.info.UserID,
							args.info.Text,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
							AddRow(1, time.Now()),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.NewCommentInfo{
					UserID: 1,
					TaskID: 1,
					Text:   "dfdfdfdf",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.comment").
						Columns(newCommentFields...).
						Values(args.info.TaskID, args.info.UserID, args.info.Text).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.TaskID,
							args.info.UserID,
							args.info.Text,
						).
						WillReturnError(apperrors.ErrCommentNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCommentNotCreated,
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

			s := NewCommentStorage(db)

			if _, err := s.Create(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCommentStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCommentStorage_GetFromTask(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.TaskID
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
				info: dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allCommentFields...).
						From("public.comment").
						Where(sq.Eq{"id_task": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnRows(sqlmock.NewRows(allCommentFields).
							AddRow(1, 1, "content", time.Now()),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.TaskID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allCommentFields...).
						From("public.comment").
						Where(sq.Eq{"id_task": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrCouldNotGetTaskComments)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetTaskComments,
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

			s := NewCommentStorage(db)

			if _, err := s.GetFromTask(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCommentStorage.GetFromTask() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCommentStorage_ReadMany(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.CommentIDs
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
				info: dto.CommentIDs{
					Values: []string{"1", "2", "3"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allCommentFields...).
						From("public.comment").
						Where(sq.Eq{"id": args.info.Values}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Values[0],
							args.info.Values[1],
							args.info.Values[2],
						).
						WillReturnRows(sqlmock.NewRows(allCommentFields).
							AddRow(1, 1, "content", time.Now()),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.CommentIDs{
					Values: []string{"1", "2", "3"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allCommentFields...).
						From("public.comment").
						Where(sq.Eq{"id": args.info.Values}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Values[0],
							args.info.Values[1],
							args.info.Values[2],
						).
						WillReturnError(apperrors.ErrCouldNotGetComments)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetComments,
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

			s := NewCommentStorage(db)

			if _, err := s.ReadMany(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCommentStorage.ReadMany() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
