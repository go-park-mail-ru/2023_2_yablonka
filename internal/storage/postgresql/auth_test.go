package postgresql

import (
	"context"
	"server/internal/apperrors"
	"server/internal/config"
	logging "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func getLogger() logging.ILogger {
	logger, _ := logging.NewLogrusLogger(&config.LoggingConfig{
		Level:                  "debug",
		DisableTimestamp:       false,
		FullTimestamp:          true,
		LevelBasedReport:       true,
		DisableLevelTruncation: true,
		ReportCaller:           true,
	})
	return &logger
}

func TestPostgresAuthStorage_CreateSession(t *testing.T) {
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
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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

/*
func TestPostgresAuthStorage_DeleteSession(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		token dto.SessionToken
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
			s := PostgresAuthStorage{
				db: tt.fields.db,
			}
			if err := s.DeleteSession(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresAuthStorage_GetSession(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		token dto.SessionToken
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresAuthStorage{
				db: tt.fields.db,
			}
			got, err := s.GetSession(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

*/
