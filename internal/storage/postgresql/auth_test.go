package postgresql_test

import (
	"context"
	"database/sql"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage/postgresql"
	"testing"
)

func TestPostgresAuthStorage_CreateSession(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx     context.Context
		session *entities.Session
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*
				dbConnection, _ := pgxmock.NewConn()
				s := postgresql.NewAuthStorage(dbConnection)
				if err := s.CreateSession(tt.args.ctx, tt.args.session); (err != nil) != tt.wantErr {
					t.Errorf("CreateSession() error = %v, wantErr %v", err, tt.wantErr)
				}
			*/
		})
	}
}

func TestPostgresAuthStorage_DeleteSession(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		token dto.SessionToken
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
			s := postgresql.NewAuthStorage(tt.fields.db)
			if err := s.DeleteSession(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresAuthStorage_GetSession(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		token dto.SessionToken
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := postgresql.NewAuthStorage(tt.fields.db)
			got, err := s.GetSession(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}
