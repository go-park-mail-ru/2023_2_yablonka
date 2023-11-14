package service

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"testing"
)

func TestUserService_RegisterUser(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
	}
	type args struct {
		ctx  context.Context
		info dto.AuthInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := UserService{
				storage: tt.fields.storage,
			}
			got, err := us.RegisterUser(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
