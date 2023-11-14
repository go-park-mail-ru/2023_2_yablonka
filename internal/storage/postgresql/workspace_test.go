package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
)

func TestNewWorkspaceStorage(t *testing.T) {
	type args struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *PostgresWorkspaceStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWorkspaceStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWorkspaceStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_Create(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.NewWorkspaceInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Workspace
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresWorkspaceStorage{
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

func TestPostgresWorkspaceStorage_Delete(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.WorkspaceID
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
			s := PostgresWorkspaceStorage{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_GetUserGuestWorkspaces(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		userID dto.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.UserGuestWorkspaceInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresWorkspaceStorage{
				db: tt.fields.db,
			}
			got, err := s.GetUserGuestWorkspaces(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserGuestWorkspaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserGuestWorkspaces() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_GetUserOwnedWorkspaces(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		userID dto.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.UserOwnedWorkspaceInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresWorkspaceStorage{
				db: tt.fields.db,
			}
			got, err := s.GetUserOwnedWorkspaces(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserOwnedWorkspaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserOwnedWorkspaces() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_GetWorkspace(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.WorkspaceID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Workspace
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresWorkspaceStorage{
				db: tt.fields.db,
			}
			got, err := s.GetWorkspace(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWorkspace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWorkspace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_UpdateData(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedWorkspaceInfo
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
			s := PostgresWorkspaceStorage{
				db: tt.fields.db,
			}
			if err := s.UpdateData(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresWorkspaceStorage_UpdateUsers(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.ChangeWorkspaceGuestsInfo
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
			s := PostgresWorkspaceStorage{
				db: tt.fields.db,
			}
			if err := s.UpdateUsers(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
