package postgresql

import (
	"context"
	"errors"
	"regexp"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func TestBoardStorage_GetUsers(t *testing.T) {
	t.Parallel()
	type args struct {
		id    dto.BoardID
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Succesful run",
			args: args{
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allPublicUserFields...).
						From("public.user").
						Join("public.board_user ON public.board_user.id_user = public.user.id").
						Where(sq.Eq{"public.board_user.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows(allPublicUserFields).
							AddRow(1, "foo@bar.com", "A", "B", "Desc", "avatar.jpg").
							AddRow(2, "foo@baz.com", "C", "D", "", "avatar.jpg"),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				id: dto.BoardID{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (could not exec query)",
			args: args{
				id: dto.BoardID{
					Value: 0,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allPublicUserFields...).
						From("public.user").
						Join("public.board_user ON public.board_user.id_user = public.user.id").
						Where(sq.Eq{"public.board_user.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnError(errors.New("Mock select query fail"))
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetBoardUsers,
		},
		{
			name: "Bad request (could not scan results)",
			args: args{
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allPublicUserFields...).
						From("public.user").
						Join("public.board_user ON public.board_user.id_user = public.user.id").
						Where(sq.Eq{"public.board_user.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows(allPublicUserFields).
							AddRow(1, nil, nil, nil, nil, nil),
						)
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

			tt.args.query(mock, tt.args)

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			_, err = s.GetUsers(ctx, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_GetById(t *testing.T) {
	t.Parallel()
	type args struct {
		id    dto.BoardID
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
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allBoardFields...).
						From("public.board").
						Where(sq.Eq{"public.board.id": args.id.Value}).
						LeftJoin("public.workspace ON public.board.id_workspace = public.workspace.id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows(allBoardFields).
							AddRow(1, 1, 1, "Mock board", time.Now(), "thumbnail.png"),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				id: dto.BoardID{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (could not scan results)",
			args: args{
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allBoardFields...).
						From("public.board").
						Where(sq.Eq{"public.board.id": args.id.Value}).
						LeftJoin("public.workspace ON public.board.id_workspace = public.workspace.id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows(allBoardFields).
							AddRow(0, 0, 0, nil, time.Now(), nil),
						)
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

			tt.args.query(mock, tt.args)

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			_, err = s.GetById(ctx, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_CheckAccess(t *testing.T) {
	t.Parallel()
	type args struct {
		info       dto.CheckBoardAccessInfo
		query      func(mock sqlmock.Sqlmock, args args)
		wantAccess bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Succesful run (has access)",
			args: args{
				info: dto.CheckBoardAccessInfo{
					UserID:  1,
					BoardID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select("count(*)").
						From("public.board_user").
						Where(sq.Eq{
							"id_board": args.info.BoardID,
							"id_user":  args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).
							AddRow(1),
						)
				},
				wantAccess: true,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Succesful run (no access)",
			args: args{
				info: dto.CheckBoardAccessInfo{
					UserID:  1,
					BoardID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select("count(*)").
						From("public.board_user").
						Where(sq.Eq{
							"id_board": args.info.BoardID,
							"id_user":  args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).
							AddRow(0),
						)
				},
				wantAccess: false,
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				info: dto.CheckBoardAccessInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
				wantAccess: false,
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (failed to parse result)",
			args: args{
				info: dto.CheckBoardAccessInfo{
					UserID:  1,
					BoardID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select("count(*)").
						From("public.board_user").
						Where(sq.Eq{
							"id_board": args.info.BoardID,
							"id_user":  args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnError(apperrors.ErrCouldNotGetUser)
				},
				wantAccess: false,
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetUser,
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

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			hasAccess, err := s.CheckAccess(ctx, tt.args.info)

			if hasAccess != tt.args.wantAccess {
				t.Errorf("CheckAccess() hasAccess = %v, wantAccess %v", hasAccess, tt.args.wantAccess)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAccess() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_GetLists(t *testing.T) {
	t.Parallel()
	type args struct {
		id    dto.BoardID
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
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allListTaskAggFields...).
						From("public.list").
						LeftJoin("public.task ON public.task.id_list = public.list.id").
						Where(sq.Eq{"public.list.id_board": args.id.Value}).
						GroupBy("public.list.id").
						OrderBy("public.list.list_position").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows(allListTaskAggFields).
							AddRow(1, 1, "Mock board", 0, pq.StringArray{}),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				id: dto.BoardID{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (could not get list)",
			args: args{
				id: dto.BoardID{
					Value: 0,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allListTaskAggFields...).
						From("public.list").
						LeftJoin("public.task ON public.task.id_list = public.list.id").
						Where(sq.Eq{"public.list.id_board": args.id.Value}).
						GroupBy("public.list.id").
						OrderBy("public.list.list_position").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnError(apperrors.ErrCouldNotGetList)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetList,
		},
		{
			name: "Bad request (could not parse result)",
			args: args{
				id: dto.BoardID{
					Value: 0,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allListTaskAggFields...).
						From("public.list").
						LeftJoin("public.task ON public.task.id_list = public.list.id").
						Where(sq.Eq{"public.list.id_board": args.id.Value}).
						GroupBy("public.list.id").
						OrderBy("public.list.list_position").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows(allListTaskAggFields).
							AddRow(nil, nil, nil, nil, nil),
						)
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

			tt.args.query(mock, tt.args)

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			_, err = s.GetLists(ctx, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetLists() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_GetTags(t *testing.T) {
	t.Parallel()
	type args struct {
		id    dto.BoardID
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
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allTagFields...).
						From("public.tag").
						LeftJoin("public.tag_board ON public.tag.id = public.tag_board.id_tag").
						Where(sq.Eq{"public.tag_board.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(sqlmock.NewRows(allTagFields).
							AddRow(1, "Mock tag", "color"),
						)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query failed",
			args: args{
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allTagFields...).
						From("public.tag").
						LeftJoin("public.tag_board ON public.tag.id = public.tag_board.id_tag").
						Where(sq.Eq{"public.tag_board.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Collecting rows failed",
			args: args{
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allTagFields...).
						From("public.tag").
						LeftJoin("public.tag_board ON public.tag.id = public.tag_board.id_tag").
						Where(sq.Eq{"public.tag_board.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(
							sqlmock.NewRows(allTagFields).
								AddRow(1, "Mock tag", "color").
								RowError(0, apperrors.ErrCouldNotCollectRows))
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotCollectRows,
		},
		{
			name: "Scanning rows failed",
			args: args{
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allTagFields...).
						From("public.tag").
						LeftJoin("public.tag_board ON public.tag.id = public.tag_board.id_tag").
						Where(sq.Eq{"public.tag_board.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(
							sqlmock.NewRows(allTagFields).
								AddRow(nil, nil, nil))
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotScanRows,
		},
		{
			name: "Closing rows failed",
			args: args{
				id: dto.BoardID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.Select(allTagFields...).
						From("public.tag").
						LeftJoin("public.tag_board ON public.tag.id = public.tag_board.id_tag").
						Where(sq.Eq{"public.tag_board.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnRows(
							sqlmock.NewRows(allTagFields).
								AddRow(1, "Mock tag", "color").
								CloseError(apperrors.ErrCouldNotCloseQuery))
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotCloseQuery,
		},
		{
			name: "Error building query",
			args: args{
				id:    dto.BoardID{},
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

			tt.args.query(mock, tt.args)

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			if _, err := s.GetTags(ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetLists() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_UpdateData(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.UpdatedBoardInfo
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
				info: dto.UpdatedBoardInfo{
					ID:   1,
					Name: "New mock name",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Update("public.board").
						Set("name", args.info.Name).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Name, args.info.ID).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				info: dto.UpdatedBoardInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (could not update board)",
			args: args{
				info: dto.UpdatedBoardInfo{
					ID:   0,
					Name: "Bad mock name",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Update("public.board").
						Set("name", args.info.Name).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Name, args.info.ID).
						WillReturnError(apperrors.ErrBoardNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrBoardNotUpdated,
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

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			err = s.UpdateData(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateData() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_UpdateThumbnailUrl(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.BoardImageUrlInfo
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
				info: dto.BoardImageUrlInfo{
					ID:  1,
					Url: "mock.png",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Update("public.board").
						Set("thumbnail_url", args.info.Url).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Url, args.info.ID).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				info: dto.BoardImageUrlInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (could not execute query)",
			args: args{
				info: dto.BoardImageUrlInfo{
					ID:  0,
					Url: "Bad url",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Update("public.board").
						Set("thumbnail_url", args.info.Url).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Url, args.info.ID).
						WillReturnError(apperrors.ErrBoardNotUpdated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrBoardNotUpdated,
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

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			err = s.UpdateThumbnailUrl(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateThumbnailUrl() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_Create(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.NewBoardInfo
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
			name: "Succesful run (provided url)",
			args: args{
				info: dto.NewBoardInfo{
					Name:        "New mock board",
					WorkspaceID: 1,
					OwnerID:     1,
					ThumbnailURL: func() *string {
						url := "img/board_thumbnails/mock_theme.jpg"
						return &url
					}(),
				},
				user: &entities.User{
					ID:    1,
					Email: "mock@mail.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardID := 1

					mock.ExpectBegin()

					insertQuery, _, _ := sq.
						Insert("public.board").
						Columns("id_workspace", "name").
						Values(args.info.WorkspaceID, args.info.Name).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(insertQuery)).
						WithArgs(args.info.WorkspaceID, args.info.Name).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
							AddRow(boardID, time.Now()),
						)

					avatarQuery, _, _ := sq.
						Update("public.board").
						Set("thumbnail_url", args.info.ThumbnailURL).
						Where(sq.And{
							sq.Eq{"id_workspace": args.info.WorkspaceID},
							sq.Eq{"id": boardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(avatarQuery)).
						WithArgs(args.info.ThumbnailURL, args.info.WorkspaceID, boardID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					boardUserQuery, _, _ := sq.
						Insert("public.board_user").
						Columns("id_board", "id_user").
						Values(boardID, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
						WithArgs(boardID, args.info.OwnerID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit().WillReturnError(nil)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Succesful run (no url provided)",
			args: args{
				info: dto.NewBoardInfo{
					Name:        "New mock board",
					WorkspaceID: 1,
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "mock@mail.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardID := 1
					url := "img/board_thumbnails/1.png"

					insertQuery, _, _ := sq.
						Insert("public.board").
						Columns("id_workspace", "name").
						Values(args.info.WorkspaceID, args.info.Name).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					avatarQuery, _, _ := sq.
						Update("public.board").
						Set("thumbnail_url", url).
						Where(sq.And{
							sq.Eq{"id_workspace": args.info.WorkspaceID},
							sq.Eq{"id": boardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					boardUserQuery, _, _ := sq.
						Insert("public.board_user").
						Columns("id_board", "id_user").
						Values(boardID, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectBegin()

					mock.ExpectQuery(regexp.QuoteMeta(insertQuery)).
						WithArgs(args.info.WorkspaceID, args.info.Name).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
							AddRow(boardID, time.Now()),
						)

					mock.ExpectExec(regexp.QuoteMeta(avatarQuery)).
						WithArgs(url, args.info.WorkspaceID, boardID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
						WithArgs(boardID, args.info.OwnerID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit().WillReturnError(nil)
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (no user object in context)",
			args: args{
				info: dto.NewBoardInfo{
					Name:        "New mock board",
					WorkspaceID: 1,
					OwnerID:     1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {

				},
			},
			wantErr: true,
			err:     apperrors.ErrUserNotFound,
		},
		{
			name: "Bad request (couldn't build insert query)",
			args: args{
				info: dto.NewBoardInfo{},
				user: &entities.User{
					ID:    1,
					Email: "mock@mail.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (could not start transaction)",
			args: args{
				info: dto.NewBoardInfo{
					Name:        "New mock board",
					WorkspaceID: 1,
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "mock@mail.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin().WillReturnError(errors.New("Mock tx error"))
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBeginTransaction,
		},
		{
			name: "Bad request (could not insert board)",
			args: args{
				info: dto.NewBoardInfo{
					Name:        "Bad return data",
					WorkspaceID: 1,
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "mock@mail.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					insertQuery, _, _ := sq.
						Insert("public.board").
						Columns("id_workspace", "name").
						Values(args.info.WorkspaceID, args.info.Name).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					mock.ExpectBegin()

					mock.ExpectQuery(regexp.QuoteMeta(insertQuery)).
						WithArgs(args.info.WorkspaceID, args.info.Name).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
							AddRow(nil, nil),
						)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrBoardNotCreated,
		},
		{
			name: "Bad request (couldn't update board with thumbnail URL)",
			args: args{
				info: dto.NewBoardInfo{
					Name:        "New mock board",
					WorkspaceID: 1,
					OwnerID:     1,
				},
				user: &entities.User{
					ID:    1,
					Email: "mock@mail.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardID := 1
					url := "img/board_thumbnails/1.png"

					insertQuery, _, _ := sq.
						Insert("public.board").
						Columns("id_workspace", "name").
						Values(args.info.WorkspaceID, args.info.Name).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					avatarQuery, _, _ := sq.
						Update("public.board").
						Set("thumbnail_url", url).
						Where(sq.And{
							sq.Eq{"id_workspace": args.info.WorkspaceID},
							sq.Eq{"id": 0},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectBegin()

					mock.ExpectQuery(regexp.QuoteMeta(insertQuery)).
						WithArgs(args.info.WorkspaceID, args.info.Name).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
							AddRow(boardID, time.Now()),
						)

					mock.ExpectExec(regexp.QuoteMeta(avatarQuery)).
						WithArgs(url, args.info.WorkspaceID, boardID).
						WillReturnError(apperrors.ErrBoardNotCreated)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrBoardNotCreated,
		},
		{
			name: "Bad request (could not update BoardUser table)",
			args: args{
				info: dto.NewBoardInfo{
					Name:        "New mock board",
					WorkspaceID: 1,
				},
				user: &entities.User{
					ID:    1,
					Email: "mock@mail.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardID := 1
					url := "img/board_thumbnails/1.png"

					insertQuery, _, _ := sq.
						Insert("public.board").
						Columns("id_workspace", "name").
						Values(args.info.WorkspaceID, args.info.Name).
						PlaceholderFormat(sq.Dollar).
						Suffix("RETURNING id, date_created").
						ToSql()

					avatarQuery, _, _ := sq.
						Update("public.board").
						Set("thumbnail_url", url).
						Where(sq.And{
							sq.Eq{"id_workspace": args.info.WorkspaceID},
							sq.Eq{"id": boardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					boardUserQuery, _, _ := sq.
						Insert("public.board_user").
						Columns("id_board", "id_user").
						Values(1, args.info.OwnerID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectBegin()

					mock.ExpectQuery(regexp.QuoteMeta(insertQuery)).
						WithArgs(args.info.WorkspaceID, args.info.Name).
						WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
							AddRow(boardID, time.Now()),
						)

					mock.ExpectExec(regexp.QuoteMeta(avatarQuery)).
						WithArgs(url, args.info.WorkspaceID, boardID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
						WithArgs(boardID, args.info.OwnerID).
						WillReturnError(apperrors.ErrBoardNotCreated)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrBoardNotCreated,
		},
		// {
		// 	name: "Bad request (could not commit transactions)",
		// 	args: args{
		// 		info: dto.NewBoardInfo{
		// 			Name:        "New mock board",
		// 			WorkspaceID: 1,
		// 			OwnerID:     1,
		// 		},
		// 		user: &entities.User{
		// 			ID:    1,
		// 			Email: "mock@mail.com",
		// 		},
		// 		query: func(mock sqlmock.Sqlmock, args args) {
		// 			boardID := 1
		// 			url := "img/board_thumbnails/1.png"

		// 			insertQuery, _, _ := sq.
		// 				Insert("public.board").
		// 				Columns("id_workspace", "name").
		// 				Values(args.info.WorkspaceID, args.info.Name).
		// 				PlaceholderFormat(sq.Dollar).
		// 				Suffix("RETURNING id, date_created").
		// 				ToSql()

		// 			avatarQuery, _, _ := sq.
		// 				Update("public.board").
		// 				Set("thumbnail_url", url).
		// 				Where(sq.And{
		// 					sq.Eq{"id_workspace": args.info.WorkspaceID},
		// 					sq.Eq{"id": boardID},
		// 				}).
		// 				PlaceholderFormat(sq.Dollar).
		// 				ToSql()

		// 			boardUserQuery, _, _ := sq.
		// 				Insert("public.board_user").
		// 				Columns("id_board", "id_user").
		// 				Values(1, args.info.OwnerID).
		// 				PlaceholderFormat(sq.Dollar).
		// 				ToSql()

		// 			mock.ExpectBegin()

		// 			mock.ExpectQuery(regexp.QuoteMeta(insertQuery)).
		// 				WithArgs(args.info.WorkspaceID, args.info.Name).
		// 				WillReturnRows(sqlmock.NewRows([]string{"id", "date_created"}).
		// 					AddRow(boardID, time.Now()),
		// 				)

		// 			mock.ExpectExec(regexp.QuoteMeta(avatarQuery)).
		// 				WithArgs(url, args.info.WorkspaceID, boardID).
		// 				WillReturnResult(sqlmock.NewResult(1, 1))

		// 			mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
		// 				WithArgs(boardID, args.info.OwnerID).
		// 				WillReturnResult(sqlmock.NewResult(1, 1))

		// 			mock.ExpectCommit().WillReturnError(errors.New("Mock commit error"))
		// 			mock.ExpectRollback()
		// 		},
		// 	},
		// 	wantErr: true,
		// 	err:     apperrors.ErrBoardNotCreated,
		// },
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

			ctx := context.WithValue(
				context.WithValue(
					context.WithValue(context.Background(), dto.UserObjKey, tt.args.user),
					dto.RequestIDKey, uuid.New(),
				), dto.LoggerKey, getLogger(),
			)
			s := NewBoardStorage(db)

			_, err = s.Create(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.BoardDeleteRequest
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
				info: dto.BoardDeleteRequest{
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Delete("public.board").
						Where(sq.Eq{"id": args.info.BoardID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(args.info.BoardID).
						WillReturnResult(sqlmock.NewResult(1, 1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (could not build query)",
			args: args{
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (couldn't execute query)",
			args: args{
				info: dto.BoardDeleteRequest{
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					sql, _, _ := sq.
						Delete("public.board").
						Where(sq.Eq{"id": args.info.BoardID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(args.info.BoardID).
						WillReturnError(errors.New("Mock delete fail"))
				},
			},
			wantErr: true,
			err:     apperrors.ErrBoardNotDeleted,
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

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			err = s.Delete(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_AddUser(t *testing.T) {
	t.Parallel()
	type args struct {
		info      dto.AddBoardUserInfo
		addedUser dto.UserPublicInfo
		query     func(mock sqlmock.Sqlmock, args args)
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
				info: dto.AddBoardUserInfo{
					UserID:      1,
					WorkspaceID: 1,
					BoardID:     1,
				},
				addedUser: dto.UserPublicInfo{
					ID:    1,
					Email: "mock@user.com",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardUserQuery, _, _ := sq.
						Insert("public.board_user").
						Columns("id_board", "id_user").
						Values(args.info.BoardID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					userWorkSpaceQuery, _, _ := sq.
						Insert("public.user_workspace").
						Columns("id_workspace", "id_user").
						Values(args.info.WorkspaceID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					userQuery, _, _ := sq.
						Select(allPublicUserFields...).
						From("public.user").
						Where(sq.Eq{"id": args.info.UserID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectBegin()

					mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta(userWorkSpaceQuery)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()

					mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
						WithArgs(args.info.UserID).
						WillReturnRows(sqlmock.NewRows(allPublicUserFields).
							AddRow(args.addedUser.ID, args.addedUser.Email, "", "", "", ""),
						)

				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Bad request (couldn't build BoardUser query)",
			args: args{
				info: dto.AddBoardUserInfo{
					UserID:      1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (couldn't execute BoardUser query)",
			args: args{
				info: dto.AddBoardUserInfo{
					UserID:      1,
					WorkspaceID: 1,
					BoardID:     1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardUserQuery, _, _ := sq.
						Insert("public.board_user").
						Columns("id_board", "id_user").
						Values(args.info.BoardID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectBegin()

					mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnError(errors.New("Mock insert query fail"))

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotAddBoardUser,
		},
		{
			name: "Bad request (could not build UserWorkspace query)",
			args: args{
				info: dto.AddBoardUserInfo{
					UserID:      1,
					WorkspaceID: 1,
					BoardID:     1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardUserQuery, _, _ := sq.
						Insert("public.board_user").
						Columns("id_board", "id_user").
						Values(args.info.BoardID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectBegin()

					mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Bad request (could not execute UserWorkspace query)",
			args: args{
				info: dto.AddBoardUserInfo{
					UserID:      1,
					WorkspaceID: 1,
					BoardID:     1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					boardUserQuery, _, _ := sq.
						Insert("public.board_user").
						Columns("id_board", "id_user").
						Values(args.info.BoardID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					userWorkSpaceQuery, _, _ := sq.
						Insert("public.user_workspace").
						Columns("id_workspace", "id_user").
						Values(args.info.WorkspaceID, args.info.UserID).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectBegin()

					mock.ExpectExec(regexp.QuoteMeta(boardUserQuery)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectExec(regexp.QuoteMeta(userWorkSpaceQuery)).
						WithArgs(args.info.BoardID, args.info.UserID).
						WillReturnError(errors.New("Mock insert query error"))

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotAddBoardUser,
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

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			_, err = s.AddUser(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestBoardStorage_RemoveUser(t *testing.T) {
	t.Parallel()
	type args struct {
		info  dto.RemoveBoardUserInfo
		query func(mock sqlmock.Sqlmock, args args)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path (no boards left)",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query2, _, _ := sq.Select("count(*)").
						From("public.board_user").
						LeftJoin("public.board ON public.board_user.id_board = public.board.id").
						Where(sq.Eq{
							"public.board.id_workspace": args.info.WorkspaceID,
							"public.board_user.id_user": args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.WorkspaceID,
							args.info.UserID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).
							AddRow(0),
						)

					query3, _, _ := sq.
						Delete("public.user_workspace").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_workspace": args.info.WorkspaceID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query3)).
						WithArgs(
							args.info.UserID,
							args.info.WorkspaceID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit()
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Happy path (boards left)",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query2, _, _ := sq.Select("count(*)").
						From("public.board_user").
						LeftJoin("public.board ON public.board_user.id_board = public.board.id").
						Where(sq.Eq{
							"public.board.id_workspace": args.info.WorkspaceID,
							"public.board_user.id_user": args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.UserID,
							args.info.WorkspaceID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).
							AddRow(1),
						)

					mock.ExpectCommit()
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Building query 1 failed",
			args: args{
				info:  dto.RemoveBoardUserInfo{},
				query: func(mock sqlmock.Sqlmock, args args) {},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBuildQuery,
		},
		{
			name: "Beginning transaction failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin().WillReturnError(apperrors.ErrCouldNotBeginTransaction)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotBeginTransaction,
		},
		{
			name: "Query 1 execution failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
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
			name: "Executing query 2 failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query2, _, _ := sq.Select("count(*)").
						From("public.board_user").
						LeftJoin("public.board ON public.board_user.id_board = public.board.id").
						Where(sq.Eq{
							"public.board.id_workspace": args.info.WorkspaceID,
							"public.board_user.id_user": args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.WorkspaceID,
							args.info.UserID,
						).
						WillReturnError(apperrors.ErrCouldNotScanRows)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotScanRows,
		},
		{
			name: "Executing query 3 failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query2, _, _ := sq.Select("count(*)").
						From("public.board_user").
						LeftJoin("public.board ON public.board_user.id_board = public.board.id").
						Where(sq.Eq{
							"public.board.id_workspace": args.info.WorkspaceID,
							"public.board_user.id_user": args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.UserID,
							args.info.WorkspaceID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback()
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotExecuteQuery,
		},
		{
			name: "Commiting failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query2, _, _ := sq.Select("count(*)").
						From("public.board_user").
						LeftJoin("public.board ON public.board_user.id_board = public.board.id").
						Where(sq.Eq{
							"public.board.id_workspace": args.info.WorkspaceID,
							"public.board_user.id_user": args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.WorkspaceID,
							args.info.UserID,
						).
						WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).
							AddRow(0),
						)

					query3, _, _ := sq.
						Delete("public.user_workspace").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_workspace": args.info.WorkspaceID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query3)).
						WithArgs(
							args.info.UserID,
							args.info.WorkspaceID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					mock.ExpectCommit().WillReturnError(apperrors.ErrCouldNotCommit)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotCommit,
		},
		{
			name: "Query 1 execution rollback failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
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
			name: "Executing query 2 and rollback failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query2, _, _ := sq.Select("count(*)").
						From("public.board_user").
						LeftJoin("public.board ON public.board_user.id_board = public.board.id").
						Where(sq.Eq{
							"public.board.id_workspace": args.info.WorkspaceID,
							"public.board_user.id_user": args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.WorkspaceID,
							args.info.UserID,
						).
						WillReturnError(apperrors.ErrCouldNotScanRows)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
		},
		{
			name: "Executing query 3 and rollback failed",
			args: args{
				info: dto.RemoveBoardUserInfo{
					UserID:      1,
					BoardID:     1,
					WorkspaceID: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					mock.ExpectBegin()

					query1, _, _ := sq.
						Delete("public.board_user").
						Where(sq.And{
							sq.Eq{"id_user": args.info.UserID},
							sq.Eq{"id_board": args.info.BoardID},
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(query1)).
						WithArgs(
							args.info.UserID,
							args.info.BoardID,
						).
						WillReturnResult(sqlmock.NewResult(1, 1))

					query2, _, _ := sq.Select("count(*)").
						From("public.board_user").
						LeftJoin("public.board ON public.board_user.id_board = public.board.id").
						Where(sq.Eq{
							"public.board.id_workspace": args.info.WorkspaceID,
							"public.board_user.id_user": args.info.UserID,
						}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query2)).
						WithArgs(
							args.info.UserID,
							args.info.WorkspaceID,
						).
						WillReturnError(apperrors.ErrCouldNotExecuteQuery)

					mock.ExpectRollback().WillReturnError(apperrors.ErrCouldNotRollback)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotRollback,
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

			ctx := context.WithValue(
				context.WithValue(context.Background(), dto.LoggerKey, getLogger()),
				dto.RequestIDKey, uuid.New(),
			)

			s := NewBoardStorage(db)

			err = s.RemoveUser(ctx, tt.args.info)

			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v, err = %v", err != nil, tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
