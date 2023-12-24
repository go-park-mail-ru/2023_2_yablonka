package postgresql

import (
	"context"
	"database/sql/driver"
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

func TestPostgresWorkspaceStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewWorkspaceInfo
		user  *entities.User
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
				info: dto.NewWorkspaceInfo{
					Name:        "Name",
					Description: "Description",
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "kdanil01@mail.ru",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Insert("public.workspace").
						Columns("name", "description", "id_creator").
						Values(args.info.Name, args.info.Description, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.OwnerID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).AddRow(1, time.Now()))

					query2, _, _ := sq.
						Insert("public.user_workspace").
						Columns("id_workspace", "id_user").
						Values(1, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.OwnerID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Couldn't begin transaction",
			args: args{
				info: dto.NewWorkspaceInfo{
					Name:        "Name",
					Description: "Description",
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "kdanil01@mail.ru",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin().WillReturnError(apperrors.ErrCouldNotBeginTransaction)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBeginTransaction,
		},
		{
			name: "Bad workspace query",
			args: args{
				info: dto.NewWorkspaceInfo{
					Name:        "Name",
					Description: "Description",
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "kdanil01@mail.ru",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Insert("public.workspace").
						Columns("name", "description", "id_creator").
						Values(args.info.Name, args.info.Description, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.OwnerID,
						).
						WillReturnError(apperrors.ErrWorkspaceNotCreated)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrWorkspaceNotCreated,
		},
		{
			name: "Bad user_workspace query",
			args: args{
				info: dto.NewWorkspaceInfo{
					Name:        "Name",
					Description: "Description",
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "kdanil01@mail.ru",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Insert("public.workspace").
						Columns("name", "description", "id_creator").
						Values(args.info.Name, args.info.Description, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.OwnerID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).AddRow(1, time.Now()))

					query2, _, _ := sq.
						Insert("public.user_workspace").
						Columns("id_workspace", "id_user").
						Values(1, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.OwnerID,
						).
						WillReturnError(apperrors.ErrWorkspaceNotCreated)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrWorkspaceNotCreated,
		},
		{
			name: "Could not commit",
			args: args{
				info: dto.NewWorkspaceInfo{
					Name:        "Name",
					Description: "Description",
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "kdanil01@mail.ru",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Insert("public.workspace").
						Columns("name", "description", "id_creator").
						Values(args.info.Name, args.info.Description, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.OwnerID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).AddRow(1, time.Now()))

					query2, _, _ := sq.
						Insert("public.user_workspace").
						Columns("id_workspace", "id_user").
						Values(1, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.OwnerID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit().
						WillReturnError(apperrors.ErrWorkspaceNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrWorkspaceNotCreated,
		},
		{
			name: "Could not rollback first query",
			args: args{
				info: dto.NewWorkspaceInfo{
					Name:        "Name",
					Description: "Description",
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "kdanil01@mail.ru",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Insert("public.workspace").
						Columns("name", "description", "id_creator").
						Values(args.info.Name, args.info.Description, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.OwnerID,
						).
						WillReturnError(apperrors.ErrWorkspaceNotCreated)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
		},
		{
			name: "Could not roll back second query",
			args: args{
				info: dto.NewWorkspaceInfo{
					Name:        "Name",
					Description: "Description",
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "kdanil01@mail.ru",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Insert("public.workspace").
						Columns("name", "description", "id_creator").
						Values(args.info.Name, args.info.Description, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					mock.ExpectBegin()
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.OwnerID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).AddRow(1, time.Now()))

					query2, _, _ := sq.
						Insert("public.user_workspace").
						Columns("id_workspace", "id_user").
						Values(1, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query2)).
						WithArgs(
							1,
							args.info.OwnerID,
						).
						WillReturnError(apperrors.ErrWorkspaceNotCreated)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
		},
		{
			name: "Building query failed",
			args: args{
				info:  dto.NewWorkspaceInfo{},
				user:  &entities.User{},
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
			ctx = context.WithValue(ctx, dto.UserObjKey, tt.args.user)

			tt.args.query(mock, tt.args)

			s := NewWorkspaceStorage(db)

			if _, err := s.Create(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresWorkspaceStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		id    dto.WorkspaceID
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
				id: dto.WorkspaceID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.workspace").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Execution error",
			args: args{
				id: dto.WorkspaceID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.workspace").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnError(apperrors.ErrWorkspaceNotDeleted)

				},
			},
			wantErr: true,
			err:     apperrors.ErrWorkspaceNotDeleted,
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

			s := NewWorkspaceStorage(db)

			if err := s.Delete(ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresWorkspaceStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_UpdateData(t *testing.T) {
	t.Parallel()

	description := "New description"
	type args struct {
		info  dto.UpdatedWorkspaceInfo
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
				info: dto.UpdatedWorkspaceInfo{
					ID:          1,
					Name:        "New name",
					Description: &description,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.workspace").
						Set("name", args.info.Name).
						Set("description", args.info.Description).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.ID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query execution failed",
			args: args{
				info: dto.UpdatedWorkspaceInfo{
					ID:          1,
					Name:        "New name",
					Description: &description,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.workspace").
						Set("name", args.info.Name).
						Set("description", args.info.Description).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Description,
							args.info.ID,
						).
						WillReturnError(apperrors.ErrUserNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrUserNotUpdated,
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

			s := NewWorkspaceStorage(db)

			if err := s.UpdateData(ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresWorkspaceStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_GetUserGuestWorkspaces(t *testing.T) {
	t.Parallel()
	type args struct {
		userID     dto.UserID
		workspaces []dto.UserGuestWorkspaceInfo
		boards     []dto.BoardReturn
		query      func(mock sqlmock.Sqlmock, args args)
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
				userID: dto.UserID{
					Value: 1,
				},
				workspaces: []dto.UserGuestWorkspaceInfo{
					{
						ID: 1,
						Owner: dto.UserOwnerInfo{
							ID: 2,
						},
					},
					{
						ID: 2,
						Owner: dto.UserOwnerInfo{
							ID: 3,
						},
					},
				},
				boards: []dto.BoardReturn{
					{
						WorkspaceID: 1,
					},
					{
						WorkspaceID: 2,
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select(userGuestWorkspaceFields...).
						From("public.workspace").
						LeftJoin("public.user_workspace ON public.user_workspace.id_workspace = public.workspace.id").
						LeftJoin("public.user ON public.user.id = public.user_workspace.id_user").
						Where(sq.And{
							sq.NotEq{"public.workspace.id_creator": args.userID.Value},
							sq.Eq{"public.user_workspace.id_user": args.userID.Value},
						}).
						OrderBy("public.workspace.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					workspaceRows := sqlmock.NewRows(userGuestWorkspaceFields)
					query2Args := []driver.Value{}
					for _, workspace := range args.workspaces {
						workspaceRows.AddRow(
							workspace.ID,
							workspace.Name,
							workspace.DateCreated,
							workspace.Owner.ID,
							workspace.Owner.Email,
							workspace.Owner.Name,
							workspace.Owner.Surname,
						)
						query2Args = append(query2Args, workspace.ID)
					}
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.userID.Value,
							args.userID.Value,
						).
						WillReturnRows(workspaceRows)

					query2, _, _ := sq.
						Select(guestBoardFields...).
						From("public.board").
						LeftJoin("public.board_user ON public.board_user.id_board = public.board.id").
						Where(sq.And{
							sq.Eq{"public.board.id_workspace": query2Args},
							sq.Eq{"public.board_user.id_user": args.userID.Value},
						}).
						OrderBy("public.board.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					boardRows := sqlmock.NewRows(guestBoardFields)
					for _, board := range args.boards {
						boardRows.AddRow(
							board.WorkspaceID,
							board.ID,
							board.Name,
							board.Description,
							board.ThumbnailURL,
						)
					}

					query2Args = append(query2Args, args.userID.Value)
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(query2Args...).
						WillReturnRows(boardRows)

				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "First query fail",
			args: args{
				userID: dto.UserID{
					Value: 1,
				},
				workspaces: []dto.UserGuestWorkspaceInfo{
					{
						ID: 1,
						Owner: dto.UserOwnerInfo{
							ID: 2,
						},
					},
					{
						ID: 2,
						Owner: dto.UserOwnerInfo{
							ID: 3,
						},
					},
				},
				boards: []dto.BoardReturn{
					{
						WorkspaceID: 1,
					},
					{
						WorkspaceID: 2,
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select(userGuestWorkspaceFields...).
						From("public.workspace").
						LeftJoin("public.user_workspace ON public.user_workspace.id_workspace = public.workspace.id").
						LeftJoin("public.user ON public.user.id = public.user_workspace.id_user").
						Where(sq.And{
							sq.NotEq{"public.workspace.id_creator": args.userID.Value},
							sq.Eq{"public.user_workspace.id_user": args.userID.Value},
						}).
						OrderBy("public.workspace.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.userID.Value,
							args.userID.Value,
						).
						WillReturnError(apperrors.ErrCouldNotGetWorkspace)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetWorkspace,
		},
		{
			name: "Second query fail",
			args: args{
				userID: dto.UserID{
					Value: 1,
				},
				workspaces: []dto.UserGuestWorkspaceInfo{
					{
						ID: 1,
						Owner: dto.UserOwnerInfo{
							ID: 2,
						},
					},
					{
						ID: 2,
						Owner: dto.UserOwnerInfo{
							ID: 3,
						},
					},
				},
				boards: []dto.BoardReturn{
					{
						WorkspaceID: 1,
					},
					{
						WorkspaceID: 2,
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select(userGuestWorkspaceFields...).
						From("public.workspace").
						LeftJoin("public.user_workspace ON public.user_workspace.id_workspace = public.workspace.id").
						LeftJoin("public.user ON public.user.id = public.user_workspace.id_user").
						Where(sq.And{
							sq.NotEq{"public.workspace.id_creator": args.userID.Value},
							sq.Eq{"public.user_workspace.id_user": args.userID.Value},
						}).
						OrderBy("public.workspace.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					workspaceRows := sqlmock.NewRows(userGuestWorkspaceFields)
					query2Args := []driver.Value{}
					for _, workspace := range args.workspaces {
						workspaceRows.AddRow(
							workspace.ID,
							workspace.Name,
							workspace.DateCreated,
							workspace.Owner.ID,
							workspace.Owner.Email,
							workspace.Owner.Name,
							workspace.Owner.Surname,
						)
						query2Args = append(query2Args, workspace.ID)
					}
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.userID.Value,
							args.userID.Value,
						).
						WillReturnRows(workspaceRows)

					query2, _, _ := sq.
						Select(guestBoardFields...).
						From("public.board").
						LeftJoin("public.board_user ON public.board_user.id_board = public.board.id").
						Where(sq.And{
							sq.Eq{"public.board.id_workspace": query2Args},
							sq.Eq{"public.board_user.id_user": args.userID.Value},
						}).
						OrderBy("public.board.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					boardRows := sqlmock.NewRows(guestBoardFields)
					for _, board := range args.boards {
						boardRows.AddRow(
							board.WorkspaceID,
							board.ID,
							board.Name,
							board.Description,
							board.ThumbnailURL,
						)
					}

					query2Args = append(query2Args, args.userID.Value)
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(query2Args...).
						WillReturnError(apperrors.ErrCouldNotGetBoard)

				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetBoard,
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

			s := NewWorkspaceStorage(db)

			if _, err := s.GetUserGuestWorkspaces(ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("PostgresWorkspaceStorage.GetUserGuestWorkspaces() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_GetUserOwnedWorkspaces(t *testing.T) {
	t.Parallel()
	type args struct {
		userID     dto.UserID
		workspaces []dto.UserOwnedWorkspaceInfo
		boards     []dto.BoardReturn
		query      func(mock sqlmock.Sqlmock, args args)
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
				userID: dto.UserID{
					Value: 1,
				},
				workspaces: []dto.UserOwnedWorkspaceInfo{
					{
						ID:   1,
						Name: "name1",
					},
					{
						ID:   2,
						Name: "name2",
					},
				},
				boards: []dto.BoardReturn{
					{
						WorkspaceID: 1,
					},
					{
						WorkspaceID: 2,
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select("id", "name").
						From("public.workspace").
						Where(sq.Eq{"id_creator": args.userID.Value}).
						OrderBy("public.workspace.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					workspaceRows := sqlmock.NewRows([]string{"id", "name"})
					ownedIDs := []driver.Value{}
					for _, workspace := range args.workspaces {
						workspaceRows.AddRow(
							workspace.ID,
							workspace.Name,
						)
						ownedIDs = append(ownedIDs, workspace.ID)
					}
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.userID.Value,
						).
						WillReturnRows(workspaceRows)

					query2, _, _ := sq.
						Select("id_workspace", "id", "name", "description", "thumbnail_url").
						From("public.board").
						Where(sq.Eq{"id_workspace": ownedIDs}).
						OrderBy("public.board.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					boardRows := sqlmock.NewRows(guestBoardFields)
					for _, board := range args.boards {
						boardRows.AddRow(
							board.WorkspaceID,
							board.ID,
							board.Name,
							board.Description,
							board.ThumbnailURL,
						)
					}

					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(ownedIDs...).
						WillReturnRows(boardRows)

				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "First query fail",
			args: args{
				userID: dto.UserID{
					Value: 1,
				},
				workspaces: []dto.UserOwnedWorkspaceInfo{
					{
						ID:   1,
						Name: "name1",
					},
					{
						ID:   2,
						Name: "name2",
					},
				},
				boards: []dto.BoardReturn{
					{
						WorkspaceID: 1,
					},
					{
						WorkspaceID: 2,
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select("id", "name").
						From("public.workspace").
						Where(sq.Eq{"id_creator": args.userID.Value}).
						OrderBy("public.workspace.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					workspaceRows := sqlmock.NewRows([]string{"id", "name"})
					for _, workspace := range args.workspaces {
						workspaceRows.AddRow(
							workspace.ID,
							workspace.Name,
						)
					}
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.userID.Value,
						).
						WillReturnError(apperrors.ErrCouldNotGetWorkspace)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetWorkspace,
		},
		{
			name: "Second query failed",
			args: args{
				userID: dto.UserID{
					Value: 1,
				},
				workspaces: []dto.UserOwnedWorkspaceInfo{
					{
						ID:   1,
						Name: "name1",
					},
					{
						ID:   2,
						Name: "name2",
					},
				},
				boards: []dto.BoardReturn{
					{
						WorkspaceID: 1,
					},
					{
						WorkspaceID: 2,
					},
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query1, _, _ := sq.
						Select("id", "name").
						From("public.workspace").
						Where(sq.Eq{"id_creator": args.userID.Value}).
						OrderBy("public.workspace.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					workspaceRows := sqlmock.NewRows([]string{"id", "name"})
					ownedIDs := []driver.Value{}
					for _, workspace := range args.workspaces {
						workspaceRows.AddRow(
							workspace.ID,
							workspace.Name,
						)
						ownedIDs = append(ownedIDs, workspace.ID)
					}
					mock.ExpectQuery(regexp.QuoteMeta(query1)).
						WithArgs(
							args.userID.Value,
						).
						WillReturnRows(workspaceRows)

					query2, _, _ := sq.
						Select("id_workspace", "id", "name", "description", "thumbnail_url").
						From("public.board").
						Where(sq.Eq{"id_workspace": ownedIDs}).
						OrderBy("public.board.date_created").
						PlaceholderFormat(sq.Dollar).
						ToSql()

					boardRows := sqlmock.NewRows(guestBoardFields)
					for _, board := range args.boards {
						boardRows.AddRow(
							board.WorkspaceID,
							board.ID,
							board.Name,
							board.Description,
							board.ThumbnailURL,
						)
					}

					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(ownedIDs...).
						WillReturnError(apperrors.ErrCouldNotGetBoard)

				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetBoard,
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

			s := NewWorkspaceStorage(db)

			if _, err := s.GetUserOwnedWorkspaces(ctx, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("PostgresWorkspaceStorage.GetUserOwnedWorkspaces() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
