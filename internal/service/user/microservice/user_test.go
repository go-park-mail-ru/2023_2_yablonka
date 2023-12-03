package microservice

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	user_microservice "server/microservices/user/user"
	"testing"

	"google.golang.org/grpc"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		storage storage.IUserStorage
		conn    *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.storage, tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CheckPassword(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
		client  user_microservice.UserServiceClient
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
				client:  tt.fields.client,
			}
			got, err := us.CheckPassword(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckPassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
		client  user_microservice.UserServiceClient
	}
	type args struct {
		ctx context.Context
		id  dto.UserID
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
			us := UserService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			if err := us.DeleteUser(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_GetWithID(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
		client  user_microservice.UserServiceClient
	}
	type args struct {
		ctx context.Context
		id  dto.UserID
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
				client:  tt.fields.client,
			}
			got, err := us.GetWithID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWithID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWithID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_RegisterUser(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
		client  user_microservice.UserServiceClient
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
				client:  tt.fields.client,
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

func TestUserService_UpdateAvatar(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
		client  user_microservice.UserServiceClient
	}
	type args struct {
		ctx  context.Context
		info dto.AvatarChangeInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.UrlObj
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := UserService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			got, err := us.UpdateAvatar(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAvatar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAvatar() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdatePassword(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
		client  user_microservice.UserServiceClient
	}
	type args struct {
		ctx  context.Context
		info dto.PasswordChangeInfo
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
			us := UserService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			if err := us.UpdatePassword(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_UpdateProfile(t *testing.T) {
	type fields struct {
		storage storage.IUserStorage
		client  user_microservice.UserServiceClient
	}
	type args struct {
		ctx  context.Context
		info dto.UserProfileInfo
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
			us := UserService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			if err := us.UpdateProfile(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
