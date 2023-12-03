package microservice

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	auth_microservice "server/microservices/auth/auth"
	"server/mocks/mock_storage"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestAuthService_AuthUser(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		client          auth_microservice.AuthServiceClient
		sessionIDLength uint
		authStorage     storage.IAuthStorage
	}
	type args struct {
		ctx context.Context
		id  dto.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.SessionToken
		wantErr bool
	}{
		{
			name:    "Normal query",
			fields:  fields{},
			args:    args{},
			want:    dto.SessionToken{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			storage := mock_storage.NewMockIAuthStorage(ctrl)
			grcpConn, err := grpc.Dial(
				fmt.Sprintf("%v:%v", "localhost", "8083"),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				log.Println("Failed to connect to the GRPC server as client")
			}
			defer grcpConn.Close()

			cfg, _ := config.NewSessionConfig()

			a := NewAuthService(*cfg, storage, grcpConn)

			got, err := a.AuthUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_GetLifetime(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		client          auth_microservice.AuthServiceClient
		sessionIDLength uint
		authStorage     storage.IAuthStorage
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthService{
				sessionDuration: tt.fields.sessionDuration,
				client:          tt.fields.client,
				sessionIDLength: tt.fields.sessionIDLength,
				authStorage:     tt.fields.authStorage,
			}
			if got := a.GetLifetime(tt.args.ctx); got != tt.want {
				t.Errorf("GetLifetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_LogOut(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		client          auth_microservice.AuthServiceClient
		sessionIDLength uint
		authStorage     storage.IAuthStorage
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
			a := &AuthService{
				sessionDuration: tt.fields.sessionDuration,
				client:          tt.fields.client,
				sessionIDLength: tt.fields.sessionIDLength,
				authStorage:     tt.fields.authStorage,
			}
			if err := a.LogOut(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("LogOut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthService_VerifyAuth(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		client          auth_microservice.AuthServiceClient
		sessionIDLength uint
		authStorage     storage.IAuthStorage
	}
	type args struct {
		ctx   context.Context
		token dto.SessionToken
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.UserID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthService{
				sessionDuration: tt.fields.sessionDuration,
				client:          tt.fields.client,
				sessionIDLength: tt.fields.sessionIDLength,
				authStorage:     tt.fields.authStorage,
			}
			got, err := a.VerifyAuth(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyAuth() got = %v, want %v", got, tt.want)
			}
		})
	}
}
