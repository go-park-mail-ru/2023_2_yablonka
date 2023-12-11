package microservice

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"testing"
)

func TestCommentService_Create(t *testing.T) {
	type fields struct {
		storage storage.ICommentStorage
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
			cs := CommentService{
				storage: tt.fields.storage,
			}
			got, err := cs.Create(tt.args.ctx, tt.args.info)
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

func TestCommentService_GetFromTask(t *testing.T) {
	type fields struct {
		storage storage.ICommentStorage
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
			cs := CommentService{
				storage: tt.fields.storage,
			}
			got, err := cs.GetFromTask(tt.args.ctx, tt.args.id)
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

func TestNewCommentService(t *testing.T) {
	type args struct {
		storage storage.ICommentStorage
	}
	tests := []struct {
		name string
		args args
		want *CommentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentService(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentService() = %v, want %v", got, tt.want)
			}
		})
	}
}
