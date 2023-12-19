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
	"github.com/lib/pq"
)

func TestPostgresChecklistStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewChecklistInfo
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
				info: dto.NewChecklistInfo{
					TaskID:       1,
					Name:         "dfdfdfdf",
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.checklist").
						Columns("name", "list_position", "id_task").
						Values(args.info.Name, args.info.ListPosition, args.info.TaskID).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.ListPosition,
							args.info.TaskID,
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
				info: dto.NewChecklistInfo{
					TaskID:       1,
					Name:         "dfdfdfdf",
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.checklist").
						Columns("name", "list_position", "id_task").
						Values(args.info.Name, args.info.ListPosition, args.info.TaskID).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.ListPosition,
							args.info.TaskID,
						).
						WillReturnError(apperrors.ErrChecklistNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrChecklistNotCreated,
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

			s := NewChecklistStorage(db)

			if _, err := s.Create(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresChecklistStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.ChecklistID
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
				info: dto.ChecklistID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.checklist").
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
				info: dto.ChecklistID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.checklist").
						Where(sq.Eq{"id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrChecklistNotDeleted)
				},
			},
			wantErr: true,
			err:     apperrors.ErrChecklistNotDeleted,
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

			s := NewChecklistStorage(db)

			if err := s.Delete(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresChecklistStorage_ReadMany(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.ChecklistIDs
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
				info: dto.ChecklistIDs{
					Values: []string{"1", "2", "3"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allChecklistFields...).
						From("public.checklist").
						LeftJoin("public.checklist_item ON public.checklist.id = public.checklist_item.id_checklist").
						Where(sq.Eq{"public.checklist.id": args.info.Values}).
						GroupBy("public.checklist.id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Values[0],
							args.info.Values[1],
							args.info.Values[2],
						).
						WillReturnRows(sqlmock.NewRows(allChecklistFields).
							AddRow(1, 1, "content", 1, pq.StringArray{"1", "2", "3"}),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.ChecklistIDs{
					Values: []string{"1", "2", "3"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allChecklistFields...).
						From("public.checklist").
						LeftJoin("public.checklist_item ON public.checklist.id = public.checklist_item.id_checklist").
						Where(sq.Eq{"public.checklist.id": args.info.Values}).
						GroupBy("public.checklist.id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Values[0],
							args.info.Values[1],
							args.info.Values[2],
						).
						WillReturnError(apperrors.ErrCouldNotGetChecklist)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetChecklist,
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

			s := NewChecklistStorage(db)

			if _, err := s.ReadMany(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistStorage.ReadMany() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresChecklistStorage_Update(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.UpdatedChecklistInfo
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
				info: dto.UpdatedChecklistInfo{
					ID:           1,
					Name:         "ddfdf",
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.checklist").
						Set("name", args.info.Name).
						Set("list_position", args.info.ListPosition).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
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
			name: "Query fail",
			args: args{
				info: dto.UpdatedChecklistInfo{
					ID:           1,
					Name:         "ddfdf",
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.checklist").
						Set("name", args.info.Name).
						Set("list_position", args.info.ListPosition).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.ListPosition,
							args.info.ID,
						).
						WillReturnError(apperrors.ErrChecklistNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrChecklistNotUpdated,
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

			s := NewChecklistStorage(db)

			if err := s.Update(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistStorage.Update() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
