package microservice

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"testing"
)

func TestListService_Create(t *testing.T) {
	type fields struct {
		storage storage.IListStorage
	}
	type args struct {
		ctx  context.Context
		info dto.NewListInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.List
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := ListService{
				storage: tt.fields.storage,
			}
			got, err := ls.Create(tt.args.ctx, tt.args.info)
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

func TestListService_Delete(t *testing.T) {
	type fields struct {
		storage storage.IListStorage
	}
	type args struct {
		ctx context.Context
		id  dto.ListID
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
			ls := ListService{
				storage: tt.fields.storage,
			}
			if err := ls.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListService_Update(t *testing.T) {
	type fields struct {
		storage storage.IListStorage
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedListInfo
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
			ls := ListService{
				storage: tt.fields.storage,
			}
			if err := ls.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewListService(t *testing.T) {
	type args struct {
		storage    storage.IListStorage
		connection *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want *ListService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListService(tt.args.storage, tt.args.connection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListService() = %v, want %v", got, tt.want)
			}
		})
	}
}
