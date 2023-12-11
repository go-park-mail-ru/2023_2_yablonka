package microservice

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/storage"
	"testing"
)

func TestChecklistService_Create(t *testing.T) {
	type fields struct {
		storage storage.IChecklistStorage
	}
	type args struct {
		ctx  context.Context
		info dto.NewChecklistInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.ChecklistInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cls := ChecklistService{
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

func TestChecklistService_Delete(t *testing.T) {
	type fields struct {
		storage storage.IChecklistStorage
	}
	type args struct {
		ctx context.Context
		id  dto.ChecklistID
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
			cls := ChecklistService{
				storage: tt.fields.storage,
			}
			if err := cls.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecklistService_Update(t *testing.T) {
	type fields struct {
		storage storage.IChecklistStorage
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedChecklistInfo
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
			cls := ChecklistService{
				storage: tt.fields.storage,
			}
			if err := cls.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewChecklistService(t *testing.T) {
	type args struct {
		storage storage.IChecklistStorage
	}
	tests := []struct {
		name string
		args args
		want *ChecklistService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChecklistService(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChecklistService() = %v, want %v", got, tt.want)
			}
		})
	}
}
