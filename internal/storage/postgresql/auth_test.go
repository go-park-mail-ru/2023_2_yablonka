package postgresql

import (
	"context"
	"log"
	"regexp"
	"server/internal/apperrors"
	"server/internal/config"
	logging "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/DATA-DOG/go-sqlmock"
)

func getLogger() logging.ILogger {
	logger, _ := logging.NewLogrusLogger(&config.LoggingConfig{
		Level:                  "info",
		DisableTimestamp:       false,
		FullTimestamp:          true,
		LevelBasedReport:       true,
		DisableLevelTruncation: true,
		ReportCaller:           true,
	})
	return &logger
}

func TestPostgresAuthStorage_CreateSession(t *testing.T) {
	t.Parallel()
	type args struct {
		session *entities.Session
		query   func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Normal session",
			args: args{
				session: &entities.Session{
					SessionID:  "sdfgsdfgsdfgsdfgsdfgsdf",
					UserID:     1,
					ExpiryDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectExec("INSERT INTO public.session").
						WithArgs(
							args.session.UserID,
							args.session.ExpiryDate,
							args.session.SessionID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad query (Could not build)",
			args: args{
				session: &entities.Session{
					SessionID:  ".",
					UserID:     0,
					ExpiryDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectExec("INSERT INTO public.session").
						WithArgs(
							args.session.UserID,
							args.session.ExpiryDate,
							args.session.SessionID,
						).
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad query (Could not exec)",
			args: args{
				session: &entities.Session{
					SessionID:  ".",
					UserID:     0,
					ExpiryDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectExec("INSERT INTO public.session").
						WithArgs(
							args.session.UserID,
							args.session.ExpiryDate,
							args.session.SessionID,
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

			s := NewAuthStorage(db)

			if err := s.CreateSession(ctx, tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("CreateSession() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresAuthStorage_DeleteSession(t *testing.T) {
	t.Parallel()
	type args struct {
		token dto.SessionToken
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
				token: dto.SessionToken{
					ID:             "sdfgsdfgsdfgsdfgsdfgsdf",
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectExec("DELETE FROM public.session").
						WithArgs(
							args.token.ID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (Could not build query)",
			args: args{
				token: dto.SessionToken{
					ID:             "sdfgsdfgsdfgsdfgsdfgsdf",
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectExec("DELETE FROM public.session").
						WithArgs(
							args.token.ID,
						).
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (Session not found)",
			args: args{
				token: dto.SessionToken{
					ID:             "sdfgsdfgsdfgsdfgsdfgsdf",
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectExec("DELETE FROM public.session").
						WithArgs(
							args.token.ID,
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

			s := NewAuthStorage(db)

			if err := s.DeleteSession(ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSession() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresAuthStorage_GetSession(t *testing.T) {
	t.Parallel()
	type args struct {
		token dto.SessionToken
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
				token: dto.SessionToken{
					ID:             "sdfgsdfgsdfgsdfgsdfgsdf",
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allSessionFields...).
						From("public.session").
						Where(sq.Eq{"id_session": args.token.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.token.ID,
						).
						WillReturnRows(sqlmock.NewRows(allSessionFields).AddRow(1, args.token.ExpirationDate))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				token: dto.SessionToken{
					ID:             "sdfgsdfgsdfgsdfgsdfgsdf",
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allSessionFields...).
						From("public.session").
						Where(sq.Eq{"id_session": args.token.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.token.ID,
						).
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (session not found)",
			args: args{
				token: dto.SessionToken{
					ID:             "sdfgsdfgsdfgsdfgsdfgsdf",
					ExpirationDate: time.Now(),
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allSessionFields...).
						From("public.session").
						Where(sq.Eq{"id_session": args.token.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.token.ID,
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

			tt.args.query(mock, tt.args)

			s := NewAuthStorage(db)

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)
			_, err = s.GetSession(ctx, tt.args.token)

			log.Println(err)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
