package postgresql

import (
	"context"
	"database/sql"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
)

func TestNewCommentStorage(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *PostgresCommentStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresCommentStorage_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.NewCommentInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresCommentStorage{
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

func TestPostgresCommentStorage_GetFromTask(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  dto.TaskID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.CommentInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresCommentStorage{
				db: tt.fields.db,
			}
			got, err := s.GetFromTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFromTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFromTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresCommentStorage_ReadMany(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		ids dto.CommentIDs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.CommentInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresCommentStorage{
				db: tt.fields.db,
			}
			got, err := s.ReadMany(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadMany() got = %v, want %v", got, tt.want)
			}
		})
	}
}
