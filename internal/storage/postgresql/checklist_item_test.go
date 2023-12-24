package postgresql

import (
	"context"
	"fmt"
	"regexp"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func TestPostgresChecklistItemStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewChecklistItemInfo
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
				info: dto.NewChecklistItemInfo{
					ChecklistID:  1,
					Name:         "dfdfdfdf",
					Done:         true,
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.checklist_item").
						Columns("name", "list_position", "id_checklist", "done").
						Values(args.info.Name, args.info.ListPosition, args.info.ChecklistID, args.info.Done).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.ListPosition,
							args.info.ChecklistID,
							args.info.Done,
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
				info: dto.NewChecklistItemInfo{
					ChecklistID:  1,
					Name:         "dfdfdfdf",
					Done:         true,
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.checklist_item").
						Columns("name", "list_position", "id_checklist", "done").
						Values(args.info.Name, args.info.ListPosition, args.info.ChecklistID, args.info.Done).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.ListPosition,
							args.info.ChecklistID,
							args.info.Done,
						).
						WillReturnError(apperrors.ErrChecklistItemNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrChecklistItemNotCreated,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.NewChecklistItemInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
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

			s := NewChecklistItemStorage(db)

			if _, err := s.Create(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistItemStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresChecklistItemStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.ChecklistItemID
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
				info: dto.ChecklistItemID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.checklist_item").
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
				info: dto.ChecklistItemID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.checklist_item").
						Where(sq.Eq{"id": args.info.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Value,
						).
						WillReturnError(apperrors.ErrChecklistItemNotDeleted)
				},
			},
			wantErr: true,
			err:     apperrors.ErrChecklistItemNotDeleted,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.ChecklistItemID{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
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

			s := NewChecklistItemStorage(db)

			if err := s.Delete(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistItemStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresChecklistItemStorage_ReadMany(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.ChecklistItemStringIDs
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
				info: dto.ChecklistItemStringIDs{
					Values: []string{"1", "2", "3"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("id", "name", "list_position", "id_checklist", "done").
						From("public.checklist_item").
						Where(sq.Eq{"id": args.info.Values}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Values[0],
							args.info.Values[1],
							args.info.Values[2],
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "name", "list_position", "id_checklist", "done"}).
							AddRow(1, "name", 1, 1, true),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				info: dto.ChecklistItemStringIDs{
					Values: []string{"1", "2", "3"},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("id", "name", "list_position", "id_checklist", "done").
						From("public.checklist_item").
						Where(sq.Eq{"id": args.info.Values}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Values[0],
							args.info.Values[1],
							args.info.Values[2],
						).
						WillReturnError(apperrors.ErrCouldNotGetChecklistItem)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetChecklistItem,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.ChecklistItemStringIDs{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
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

			s := NewChecklistItemStorage(db)

			if _, err := s.ReadMany(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistItemStorage.ReadMany() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresChecklistItemStorage_Update(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.UpdatedChecklistItemInfo
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
				info: dto.UpdatedChecklistItemInfo{
					ID:           1,
					Name:         "bfdbdfb",
					Done:         true,
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.checklist_item").
						Set("name", args.info.Name).
						Set("done", args.info.Done).
						Set("list_position", args.info.ListPosition).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Done,
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
				info: dto.UpdatedChecklistItemInfo{
					ID:           1,
					Name:         "bfdbdfb",
					Done:         true,
					ListPosition: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.checklist_item").
						Set("name", args.info.Name).
						Set("done", args.info.Done).
						Set("list_position", args.info.ListPosition).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Done,
							args.info.ListPosition,
							args.info.ID,
						).
						WillReturnError(apperrors.ErrChecklistItemNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrChecklistItemNotUpdated,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.UpdatedChecklistItemInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
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

			s := NewChecklistItemStorage(db)

			if err := s.Update(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistItemStorage.Update() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresChecklistItemStorage_UpdateOrder(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.ChecklistItemIDs
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
				info: dto.ChecklistItemIDs{
					Values: []uint64{1, 2, 3},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					caseBuilder := sq.Case()
					for i, id := range args.info.Values {
						caseBuilder = caseBuilder.When(sq.Eq{"id": fmt.Sprintf("%v", id)}, fmt.Sprintf("%v", i)).Else("list_position")
					}

					query1, _, _ := sq.
						Update("public.checklist_item").
						Set("list_position", caseBuilder).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							"1", "2", "3",
						).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Executing query failed",
			args: args{
				info: dto.ChecklistItemIDs{
					Values: []uint64{1, 2, 3},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					caseBuilder := sq.Case()
					for i, id := range args.info.Values {
						caseBuilder = caseBuilder.When(sq.Eq{"id": fmt.Sprintf("%v", id)}, fmt.Sprintf("%v", i)).Else("list_position")
					}

					query1, _, _ := sq.
						Update("public.checklist_item").
						Set("list_position", caseBuilder).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							"1", "2", "3",
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.ChecklistItemIDs{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
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

			s := NewChecklistItemStorage(db)

			if err := s.UpdateOrder(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresChecklistItemStorage.UpdateOrder() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
