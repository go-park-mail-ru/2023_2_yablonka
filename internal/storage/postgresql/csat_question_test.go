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

func TestPostgresCSATQuestionStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewCSATQuestion
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
				info: dto.NewCSATQuestion{
					Content: "fgfgfgdfg",
					TypeID:  1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.question").
						Columns("content", "id_type").
						Values(args.info.Content, args.info.TypeID).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Content,
							args.info.TypeID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).
							AddRow(1),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.NewCSATQuestion{
					Content: "fgfgfgdfg",
					TypeID:  1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.question").
						Columns("content", "id_type").
						Values(args.info.Content, args.info.TypeID).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Content,
							args.info.TypeID,
						).
						WillReturnError(apperrors.ErrCouldNotCreateQuestion)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotCreateQuestion,
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

			s := NewCSATQuestionStorage(db)

			if _, err := s.Create(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATQuestionStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSATQuestionStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.CSATQuestionID
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
				info: dto.CSATQuestionID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.question").
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
			name: "Query fail",
			args: args{
				info: dto.CSATQuestionID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.question").
						Where(sq.Eq{"id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrQuestionNotDeleted)
				},
			},
			wantErr: true,
			err:     apperrors.ErrQuestionNotDeleted,
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

			s := NewCSATQuestionStorage(db)

			if err := s.Delete(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATQuestionStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSATQuestionStorage_GetAll(t *testing.T) {
	t.Parallel()
	type args struct {
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
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("public.question.id", "public.question.content", "public.question_type.name").
						From("public.question").
						Join("public.question_type ON public.question.id_type = public.question_type.id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"public.question.id", "public.question.content", "public.question_type.name"}).
							AddRow(1, "bcbcb", "cbcvbcvb"),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("public.question.id", "public.question.content", "public.question_type.name").
						From("public.question").
						Join("public.question_type ON public.question.id_type = public.question_type.id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs().
						WillReturnError(apperrors.ErrCouldNotGetQuestions)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetQuestions,
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

			s := NewCSATQuestionStorage(db)

			if _, err := s.GetAll(ctx); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATQuestionStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSATQuestionStorage_GetQuestionType(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.CSATQuestionID
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
				info: dto.CSATQuestionID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("public.question_type.id", "public.question_type.name", "public.question_type.max_score").
						From("public.question_type").
						Join("public.question ON public.question.id_type = public.question_type.id").
						Where(sq.Eq{"public.question.id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnRows(sqlmock.NewRows([]string{"public.question_type.id", "public.question_type.name", "public.question_type.max_score"}).
							AddRow(1, "dsds", 3),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.CSATQuestionID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("public.question_type.id", "public.question_type.name", "public.question_type.max_score").
						From("public.question_type").
						Join("public.question ON public.question.id_type = public.question_type.id").
						Where(sq.Eq{"public.question.id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrCouldNotGetQuestionType)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetQuestionType,
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

			s := NewCSATQuestionStorage(db)

			if _, err := s.GetQuestionType(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATQuestionStorage.GetQuestionType() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSATQuestionStorage_GetQuestionTypeWithName(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.CSATQuestionTypeName
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
				info: dto.CSATQuestionTypeName{
					Value: "gghjgj",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("id", "name", "max_score").
						From("public.question_type").
						Where(sq.Eq{"name": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "max_score"}).
							AddRow(1, "dsds", 3),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.CSATQuestionTypeName{
					Value: "gghjgj",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("id", "name", "max_score").
						From("public.question_type").
						Where(sq.Eq{"name": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrCouldNotGetQuestionType)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetQuestionType,
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

			s := NewCSATQuestionStorage(db)

			if _, err := s.GetQuestionTypeWithName(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATQuestionStorage.GetQuestionTypeWithName() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSATQuestionStorage_GetStats(t *testing.T) {
	t.Parallel()
	type args struct {
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
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select("public.question.id", "public.question.content", "public.question_type.name").
						From("public.question").
						Join("public.question_type ON public.question_type.id = public.question.id_type").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"public.question.id", "public.question.content", "public.question_type.name"}).
							AddRow(1, "dsds", "hjhhjh"),
						)
					query2, _, _ := sq.
						Select("id_question", "score", "COUNT(score)", "AVG(score)").
						From("public.answer").
						Join("public.question ON public.answer.id_question = public.question.id").
						Where(sq.Eq{"id_question": []uint64{1}}).
						OrderBy("score").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"id_question", "score", "COUNT(score)", "AVG(score)"}).
							AddRow(1, 2, 3, 4),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query 1 fauk",
			args: args{
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select("public.question.id", "public.question.content", "public.question_type.name").
						From("public.question").
						Join("public.question_type ON public.question_type.id = public.question.id_type").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs().
						WillReturnError(apperrors.ErrCouldNotGetQuestions)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetQuestions,
		},
		{
			name: "Query 2 fail",
			args: args{
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select("public.question.id", "public.question.content", "public.question_type.name").
						From("public.question").
						Join("public.question_type ON public.question_type.id = public.question.id_type").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"public.question.id", "public.question.content", "public.question_type.name"}).
							AddRow(1, "dsds", "hjhhjh"),
						)
					query2, _, _ := sq.
						Select("id_question", "score", "COUNT(score)", "AVG(score)").
						From("public.answer").
						Join("public.question ON public.answer.id_question = public.question.id").
						Where(sq.Eq{"id_question": []uint64{1}}).
						OrderBy("score").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs().
						WillReturnError(apperrors.ErrCouldNotGetQuestions)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetQuestions,
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

			s := NewCSATQuestionStorage(db)

			if _, err := s.GetStats(ctx); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATQuestionStorage.GetStats() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresCSATQuestionStorage_Update(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.UpdatedCSATQuestion
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
				info: dto.UpdatedCSATQuestion{
					ID:      1,
					Content: "gjghjg",
					Type:    1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.question").
						Set("name", args.info.Content).
						Set("id_type", args.info.Type).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Content,
							args.info.Type,
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
				info: dto.UpdatedCSATQuestion{
					ID:      1,
					Content: "gjghjg",
					Type:    1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.question").
						Set("name", args.info.Content).
						Set("id_type", args.info.Type).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Content,
							args.info.Type,
							args.info.ID,
						).
						WillReturnError(apperrors.ErrQuestionNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrQuestionNotUpdated,
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

			s := NewCSATQuestionStorage(db)

			if err := s.Update(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresCSATQuestionStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
