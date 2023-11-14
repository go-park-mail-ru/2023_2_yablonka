package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
)

func TestNewUserStorage(t *testing.T) {
	type args struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *PostgresUserStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresUserStorage_Create(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.SignupInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			got, err := s.Create(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresUserStorage_Delete(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.UserID
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
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresUserStorage_GetLoginInfoWithID(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.LoginInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			got, err := s.GetLoginInfoWithID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLoginInfoWithID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLoginInfoWithID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresUserStorage_GetWithID(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			got, err := s.GetWithID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWithID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWithID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresUserStorage_GetWithLogin(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx   context.Context
		login dto.UserLogin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			got, err := s.GetWithLogin(tt.args.ctx, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWithLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWithLogin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresUserStorage_UpdateAvatarUrl(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.ImageUrlInfo
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
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			if err := s.UpdateAvatarUrl(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAvatarUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresUserStorage_UpdatePassword(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.PasswordHashesInfo
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
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			if err := s.UpdatePassword(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresUserStorage_UpdateProfile(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.UserProfileInfo
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
			s := &PostgresUserStorage{
				db: tt.fields.db,
			}
			if err := s.UpdateProfile(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
