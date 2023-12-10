package postgresql

import (
	"context"
	"database/sql"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"testing"
)

func TestNewTaskStorage(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *PostgresTaskStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskStorage(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresTaskStorage_AddUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.AddTaskUserInfo
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
			s := PostgresTaskStorage{
				db: tt.fields.db,
			}
			if err := s.AddUser(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresTaskStorage_CheckAccess(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.CheckTaskAccessInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresTaskStorage{
				db: tt.fields.db,
			}
			got, err := s.CheckAccess(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckAccess() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresTaskStorage_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.NewTaskInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresTaskStorage{
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

func TestPostgresTaskStorage_Delete(t *testing.T) {
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostgresTaskStorage{
				db: tt.fields.db,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresTaskStorage_Read(t *testing.T) {
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
		want    *dto.SingleTaskInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresTaskStorage{
				db: tt.fields.db,
			}
			got, err := s.Read(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresTaskStorage_ReadMany(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx context.Context
		id  dto.TaskIDs
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.SingleTaskInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PostgresTaskStorage{
				db: tt.fields.db,
			}
			got, err := s.ReadMany(tt.args.ctx, tt.args.id)
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

func TestPostgresTaskStorage_RemoveUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.RemoveTaskUserInfo
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
			s := PostgresTaskStorage{
				db: tt.fields.db,
			}
			if err := s.RemoveUser(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgresTaskStorage_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedTaskInfo
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
			s := PostgresTaskStorage{
				db: tt.fields.db,
			}
			if err := s.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
