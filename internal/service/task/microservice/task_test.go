package task

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"testing"
)

func TestNewTaskService(t *testing.T) {
	type args struct {
		ts storage.ITaskStorage
		us storage.IUserStorage
	}
	tests := []struct {
		name string
		args args
		want *TaskService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskService(tt.args.ts, tt.args.us); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_AddUser(t *testing.T) {
	type fields struct {
		taskStorage storage.ITaskStorage
		userStorage storage.IUserStorage
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
			ts := TaskService{
				taskStorage: tt.fields.taskStorage,
				userStorage: tt.fields.userStorage,
			}
			if err := ts.AddUser(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskService_Create(t *testing.T) {
	type fields struct {
		taskStorage storage.ITaskStorage
		userStorage storage.IUserStorage
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
			ts := TaskService{
				taskStorage: tt.fields.taskStorage,
				userStorage: tt.fields.userStorage,
			}
			got, err := ts.Create(tt.args.ctx, tt.args.info)
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

func TestTaskService_Delete(t *testing.T) {
	type fields struct {
		taskStorage storage.ITaskStorage
		userStorage storage.IUserStorage
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
			ts := TaskService{
				taskStorage: tt.fields.taskStorage,
				userStorage: tt.fields.userStorage,
			}
			if err := ts.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskService_Read(t *testing.T) {
	type fields struct {
		taskStorage storage.ITaskStorage
		userStorage storage.IUserStorage
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
			ts := TaskService{
				taskStorage: tt.fields.taskStorage,
				userStorage: tt.fields.userStorage,
			}
			got, err := ts.Read(tt.args.ctx, tt.args.id)
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

func TestTaskService_RemoveUser(t *testing.T) {
	type fields struct {
		taskStorage storage.ITaskStorage
		userStorage storage.IUserStorage
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
			ts := TaskService{
				taskStorage: tt.fields.taskStorage,
				userStorage: tt.fields.userStorage,
			}
			if err := ts.RemoveUser(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	type fields struct {
		taskStorage storage.ITaskStorage
		userStorage storage.IUserStorage
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
			ts := TaskService{
				taskStorage: tt.fields.taskStorage,
				userStorage: tt.fields.userStorage,
			}
			if err := ts.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
