package microservice

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/storage"
	"testing"
)

func TestChecklistItemService_Create(t *testing.T) {
	type fields struct {
		storage storage.IChecklistItemStorage
	}
	type args struct {
		ctx  context.Context
		info dto.NewChecklistItemInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.ChecklistItemInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := ChecklistItemService{
				storage: tt.fields.storage,
			}
			got, err := cls.Create(tt.args.ctx, tt.args.info)
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

func TestChecklistItemService_Delete(t *testing.T) {
	type fields struct {
		storage storage.IChecklistItemStorage
	}
	type args struct {
		ctx context.Context
		id  dto.ChecklistItemID
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
			cls := ChecklistItemService{
				storage: tt.fields.storage,
			}
			if err := cls.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecklistItemService_Update(t *testing.T) {
	type fields struct {
		storage storage.IChecklistItemStorage
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedChecklistItemInfo
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
			cls := ChecklistItemService{
				storage: tt.fields.storage,
			}
			if err := cls.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewChecklistItemService(t *testing.T) {
	type args struct {
		storage storage.IChecklistItemStorage
	}
	tests := []struct {
		name string
		args args
		want *ChecklistItemService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChecklistItemService(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChecklistItemService() = %v, want %v", got, tt.want)
			}
		})
	}
}
