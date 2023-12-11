package postgresql

import (
	"context"
	"database/sql"
	"reflect"
	"server/internal/pkg/dto"
	"testing"
)

func TestNewCSATAnswerStorage(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *PostgresCSATAnswerStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCSATAnswerStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCSATAnswerStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresCSATAnswerStorage_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.NewCSATAnswer
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
			s := PostgresCSATAnswerStorage{
				db: tt.fields.db,
			}
			if err := s.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
