package postgresql

import (
	"context"
	"regexp"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
)

func TestPostgresUserStorage_Create(t *testing.T) {
	defaultAvatar := "img/user_avatars/avatar.jpg"

	t.Parallel()
	type args struct {
		info  *dto.SignupInfo
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
				info: &dto.SignupInfo{
					Email:        "sdfgsdfgsdfgsdfgsdfgsdf",
					PasswordHash: "sdfgsdfgsdfgsdfgsdfgsdf",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.user").
						Columns("email", "password_hash", "avatar_url").
						Values(args.info.Email, args.info.PasswordHash, defaultAvatar).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Email,
							args.info.PasswordHash,
							defaultAvatar,
						).
						WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Failed query",
			args: args{
				info: &dto.SignupInfo{
					Email:        "sdfgsdfgsdfgsdfgsdfgsdf",
					PasswordHash: "sdfgsdfgsdfgsdfgsdfgsdf",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Insert("public.user").
						Columns("email", "password_hash", "avatar_url").
						Values(args.info.Email, args.info.PasswordHash, defaultAvatar).
						Suffix("RETURNING id").
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Email,
							args.info.PasswordHash,
							defaultAvatar,
						).
						WillReturnError(apperrors.ErrUserNotCreated)
				},
			},
			wantErr: true,
			err:     apperrors.ErrUserNotCreated,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if _, err := s.Create(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.Create() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresUserStorage_Delete(t *testing.T) {
	t.Parallel()
	type args struct {
		id    *dto.UserID
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
				id: &dto.UserID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.user").
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
			name: "Query failed",
			args: args{
				id: &dto.UserID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Delete("public.user").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnError(apperrors.ErrUserNotDeleted)
				},
			},
			wantErr: true,
			err:     apperrors.ErrUserNotDeleted,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if err := s.Delete(ctx, *tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.Delete() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresUserStorage_UpdatePassword(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.PasswordHashesInfo
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
				info: &dto.PasswordHashesInfo{
					UserID:          1,
					NewPasswordHash: "fdfvdfvdfv",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.user").
						Set("password_hash", args.info.NewPasswordHash).
						Where(sq.Eq{"id": args.info.UserID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.NewPasswordHash,
							args.info.UserID,
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
				info: &dto.PasswordHashesInfo{
					UserID:          1,
					NewPasswordHash: "fdfvdfvdfv",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.user").
						Set("password_hash", args.info.NewPasswordHash).
						Where(sq.Eq{"id": args.info.UserID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.NewPasswordHash,
							args.info.UserID,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if err := s.UpdatePassword(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.UpdatePassword() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresUserStorage_UpdateAvatarUrl(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.UserImageUrlInfo
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
				info: &dto.UserImageUrlInfo{
					ID:  1,
					Url: "fdfvdfvdfv",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.user").
						Set("avatar_url", args.info.Url).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Url,
							args.info.ID,
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
				info: &dto.UserImageUrlInfo{
					ID:  1,
					Url: "fdfvdfvdfv",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.user").
						Set("avatar_url", args.info.Url).
						Where(sq.Eq{"id": args.info.ID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Url,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if err := s.UpdateAvatarUrl(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.UpdatePassword() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresUserStorage_UpdateProfile(t *testing.T) {
	t.Parallel()
	type args struct {
		info  *dto.UserProfileInfo
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
				info: &dto.UserProfileInfo{
					UserID:      1,
					Name:        "fdfvdfvdfv",
					Surname:     "fdfvdfvdfv",
					Description: "fdfvdfvdfv",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.user").
						Set("name", args.info.Name).
						Set("surname", args.info.Surname).
						Set("description", args.info.Description).
						Where(sq.Eq{"id": args.info.UserID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Surname,
							args.info.Description,
							args.info.UserID,
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
				info: &dto.UserProfileInfo{
					UserID:      1,
					Name:        "fdfvdfvdfv",
					Surname:     "fdfvdfvdfv",
					Description: "fdfvdfvdfv",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Update("public.user").
						Set("name", args.info.Name).
						Set("surname", args.info.Surname).
						Set("description", args.info.Description).
						Where(sq.Eq{"id": args.info.UserID}).
						PlaceholderFormat(sq.Dollar).
						ToSql()

					mock.ExpectExec(regexp.QuoteMeta(query)).
						WithArgs(
							args.info.Name,
							args.info.Surname,
							args.info.Description,
							args.info.UserID,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if err := s.UpdateProfile(ctx, *tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.UpdatePassword() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresUserStorage_GetLoginInfoWithID(t *testing.T) {
	t.Parallel()
	type args struct {
		id    *dto.UserID
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
				id: &dto.UserID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("email", "password_hash").
						From("public.user").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnRows(
							sqlmock.NewRows([]string{"email", "password_hash"}).
								AddRow("gggg", "sddfsdf"))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query failed",
			args: args{
				id: &dto.UserID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select("email", "password_hash").
						From("public.user").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnError(apperrors.ErrUserNotFound)
				},
			},
			wantErr: true,
			err:     apperrors.ErrUserNotFound,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if _, err := s.GetLoginInfoWithID(ctx, *tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.GetLoginInfoWithID() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresUserStorage_GetWithID(t *testing.T) {
	t.Parallel()
	type args struct {
		id    *dto.UserID
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
				id: &dto.UserID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allUserFields...).
						From("public.user").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnRows(
							sqlmock.NewRows(allUserFields).
								AddRow(1, "email", "password_hash", "name", "surname", "avatar_url", "description"))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query failed",
			args: args{
				id: &dto.UserID{
					Value: 1,
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allUserFields...).
						From("public.user").
						Where(sq.Eq{"id": args.id.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.id.Value,
						).
						WillReturnError(apperrors.ErrUserNotFound)
				},
			},
			wantErr: true,
			err:     apperrors.ErrUserNotFound,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if _, err := s.GetWithID(ctx, *tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.GetWithID() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestPostgresUserStorage_GetWithLogin(t *testing.T) {
	t.Parallel()
	type args struct {
		login *dto.UserLogin
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
				login: &dto.UserLogin{
					Value: "nfgnfgn",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allUserFields...).
						From("public.user").
						Where(sq.Eq{"email": args.login.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.login.Value,
						).
						WillReturnRows(
							sqlmock.NewRows(allUserFields).
								AddRow(1, "email", "password_hash", "name", "surname", "avatar_url", "description"))
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Query fail",
			args: args{
				login: &dto.UserLogin{
					Value: "nfgnfgn",
				},
				query: func(mock sqlmock.Sqlmock, args args) {
					query, _, _ := sq.
						Select(allUserFields...).
						From("public.user").
						Where(sq.Eq{"email": args.login.Value}).
						PlaceholderFormat(sq.Dollar).
						ToSql()
					mock.ExpectQuery(regexp.QuoteMeta(query)).
						WithArgs(
							args.login.Value,
						).
						WillReturnError(apperrors.ErrUserNotFound)
				},
			},
			wantErr: true,
			err:     apperrors.ErrUserNotFound,
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

			ctx := context.WithValue(context.Background(), dto.LoggerKey, getLogger())

			tt.args.query(mock, tt.args)

			s := NewUserStorage(db)

			if _, err := s.GetWithLogin(ctx, *tt.args.login); (err != nil) != tt.wantErr {
				t.Errorf("PostgresUserStorage.GetWithLogin() error = %v, wantErr %v", err != nil, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
