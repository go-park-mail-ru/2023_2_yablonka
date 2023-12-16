package postgresql

import (
	"context"
	"regexp"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func TestPostgresCSRFStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *entities.CSRF
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
				info: &entities.CSRF{
					Token:          "fgfgfgdfg",
					UserID:         1,
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.csrf").
						Columns("id_user", "expiration_date", "token").
						Values(args.info.UserID, args.info.ExpirationDate, args.info.Token).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.UserID,
							args.info.ExpirationDate,
							args.info.Token,
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
				info: &entities.CSRF{
					Token:          "fgfgfgdfg",
					UserID:         1,
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.csrf").
						Columns("id_user", "expiration_date", "token").
						Values(args.info.UserID, args.info.ExpirationDate, args.info.Token).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.UserID,
							args.info.ExpirationDate,
							args.info.Token,
						).
						WillReturnError(apperrors.ErrSessionNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrSessionNotCreated,
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

			s := NewCSRFStorage(db)

			if err := s.Create(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSRFStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSRFStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.CSRFToken
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
				info: &dto.CSRFToken{
					Value: "fgfgfgdfg",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.csrf").
						Where(sq.Eq{"token": args.info.Value}).
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
			name: "Query fail",
			args: args{
				info: &dto.CSRFToken{
					Value: "fgfgfgdfg",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.csrf").
						Where(sq.Eq{"token": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrSessionNotFound)
				},
			},
			wantErr: true,
			err:     apperrors.ErrSessionNotFound,
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

			s := NewCSRFStorage(db)

			if err := s.Delete(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSRFStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSRFStorage_Get(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.CSRFToken
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
				info: &dto.CSRFToken{
					Value: "fgfgfgdfg",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("id_user", "expiration_date").
						From("public.csrf").
						Where(sq.Eq{"token": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id_user", "expiration_date"}).
							AddRow(1, time.Now()))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: &dto.CSRFToken{
					Value: "fgfgfgdfg",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("id_user", "expiration_date").
						From("public.csrf").
						Where(sq.Eq{"token": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrCSRFNotFound)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCSRFNotFound,
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

			s := NewCSRFStorage(db)

			if _, err := s.Get(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSRFStorage.Get() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
