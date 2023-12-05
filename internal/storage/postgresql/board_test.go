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
					sql, _, _ := sq.Select(allPublicUserFields...).
						From("public.user").
						Join("public.board_user ON public.board_user.id_user = public.user.id").
						Where(sq.Eq{"public.board_user.id_board": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
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
						WillReturnError(apperrors.ErrCouldNotGetBoardUsers)
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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
			name: "Successful run",
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
					sql, _, _ := sq.Select(allBoardFields...).
						From("public.board").
						Where(sq.Eq{"public.board.id": args.id.Value}).
						LeftJoin("public.workspace ON public.board.id_workspace = public.workspace.id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(sql)).
						WithArgs(args.id.Value).
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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
			name: "Successful run",
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
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
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
						WillReturnRows(sqlmock.NewRows(allListTaskAggFields).
							AddRow(nil, nil, nil, nil, nil),
						)
				},
			},
			wantErr: true,
			err:     apperrors.ErrCouldNotGetBoard,
		},
		{
			name: "Bad request (could not scan values)",
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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
			name: "Successful run",
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
					sql, _, _ := sq.
						Update("public.board").
						Set("name", args.info.Name).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectExec(regexp.QuoteMeta(sql)).
						WithArgs(args.info.Name, args.info.ID).
						WillReturnError(apperrors.ErrCouldNotBuildQuery)
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

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

// func TestBoardStorage_UpdateThumbnailUrl(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, mock, err := sqlmock.New()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

// 			s := NewBoardStorage(db)

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

// func TestBoardStorage_Create(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, mock, err := sqlmock.New()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

// 			s := NewBoardStorage(db)

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

// func TestBoardStorage_Delete(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, mock, err := sqlmock.New()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

// 			s := NewBoardStorage(db)

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

// func TestBoardStorage_AddUser(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, mock, err := sqlmock.New()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

// 			s := NewBoardStorage(db)

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

// func TestBoardStorage_RemoveUser(t *testing.T) {
// 	t.Parallel()
// 	type args struct {
// 		query func(mock sqlmock.Sqlmock, args args)
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 		err     error
// 	}{}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, mock, err := sqlmock.New()

// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			defer db.Close()

// 			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

// 			s := NewBoardStorage(db)

// 			if err := mock.ExpectationsWereMet(); err != nil {
// 				t.Errorf("there were unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }
