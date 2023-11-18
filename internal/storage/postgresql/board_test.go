package postgresql

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestNewBoardStorage(t *testing.T) {
	type args struct {
		db *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *PostgreSQLBoardStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoardStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoardStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgreSQLBoardStorage_Create(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.NewBoardInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Board
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgreSQLBoardStorage{
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

func TestPostgreSQLBoardStorage_Delete(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.BoardID
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
			s := &PostgreSQLBoardStorage{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgreSQLBoardStorage_GetById(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.BoardID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Board
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgreSQLBoardStorage{
				db: tt.fields.db,
			}
			got, err := s.GetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgreSQLBoardStorage_GetUsers(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  dto.BoardID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.UserPublicInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgreSQLBoardStorage{
				db: tt.fields.db,
			}
			got, err := s.GetUsers(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgreSQLBoardStorage_UpdateData(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedBoardInfo
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
			s := &PostgreSQLBoardStorage{
				db: tt.fields.db,
			}
			if err := s.UpdateData(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgreSQLBoardStorage_UpdateThumbnailUrl(t *testing.T) {
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
			s := &PostgreSQLBoardStorage{
				db: tt.fields.db,
			}
			if err := s.UpdateThumbnailUrl(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateThumbnailUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
