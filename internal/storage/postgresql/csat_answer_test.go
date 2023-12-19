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

func TestPostgresCSATAnswerStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewCSATAnswer
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
				info: dto.NewCSATAnswer{
					UserID:     1,
					QuestionID: 1,
					Rating:     1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.answer").
						Columns("id_user", "id_question", "score").
						Values(args.info.UserID, args.info.QuestionID, args.info.Rating).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.UserID,
							args.info.QuestionID,
							args.info.Rating,
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
				info: dto.NewCSATAnswer{
					UserID:     1,
					QuestionID: 1,
					Rating:     1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.answer").
						Columns("id_user", "id_question", "score").
						Values(args.info.UserID, args.info.QuestionID, args.info.Rating).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.UserID,
							args.info.QuestionID,
							args.info.Rating,
						).
						WillReturnError(apperrors.ErrCouldNotCreateAnswer)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotCreateAnswer,
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

			s := NewCSATAnswerStorage(db)

			if err := s.Create(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATAnswerStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
