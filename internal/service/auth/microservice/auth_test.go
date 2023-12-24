package microservice_test

import (
	"context"
	"reflect"
	"server/internal/apperrors"
	"server/internal/config"
	logging "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/service/auth/microservice"
	auth_microservice "server/microservices/auth/auth"
	"server/mocks/mock_grcp"
	"server/mocks/mock_storage"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getLogger() logging.ILogger {
	logger, _ := logging.NewLogrusLogger(&config.LoggingConfig{
		Level:                  "debug",
		DisableTimestamp:       false,
		FullTimestamp:          true,
		LevelBasedReport:       true,
		DisableLevelTruncation: true,
		ReportCaller:           true,
	})
	return &logger
}

func TestAuthService_AuthUser(t *testing.T) {
	type args struct {
		id    dto.UserID
		ctx   context.Context
		query func(ctx context.Context, storage mock_storage.MockIAuthStorage, client mock_grcp.MockAuthServiceClient, args args) dto.SessionToken
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Happy path",
			args: args{
				ctx: context.WithValue(
					context.WithValue(
						context.Background(), dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(ctx context.Context, storage mock_storage.MockIAuthStorage, client mock_grcp.MockAuthServiceClient, args args) dto.SessionToken {
					grpcRequest := &auth_microservice.AuthUserRequest{
						RequestID: ctx.Value(dto.RequestIDKey).(uuid.UUID).String(),
						Value:     &auth_microservice.UserID{Value: args.id.Value},
					}
					response := auth_microservice.SessionToken{
						ID:             "Session ID",
						ExpirationDate: timestamppb.Now(),
					}

					client.EXPECT().AuthUser(ctx, grpcRequest).Return(
						&auth_microservice.AuthUserResponse{
							Code:     auth_microservice.ErrorCode_OK,
							Response: &response,
						},
						nil,
					)
					return dto.SessionToken{
						ID:             response.ID,
						ExpirationDate: response.ExpirationDate.AsTime(),
					}
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "No request ID",
			args: args{
				ctx: context.WithValue(
					context.Background(), dto.LoggerKey, getLogger()),
				query: func(ctx context.Context, storage mock_storage.MockIAuthStorage, client mock_grcp.MockAuthServiceClient, args args) dto.SessionToken {
					return dto.SessionToken{}
				},
			},
			wantErr: true,
			err:     apperrors.ErrNoRequestIDFound,
		},
		{
			name: "GRPC request failed",
			args: args{
				ctx: context.WithValue(
					context.WithValue(
						context.Background(), dto.LoggerKey, getLogger()),
					dto.RequestIDKey, uuid.New(),
				),
				query: func(ctx context.Context, storage mock_storage.MockIAuthStorage, client mock_grcp.MockAuthServiceClient, args args) dto.SessionToken {
					grpcRequest := &auth_microservice.AuthUserRequest{
						RequestID: ctx.Value(dto.RequestIDKey).(uuid.UUID).String(),
						Value:     &auth_microservice.UserID{Value: args.id.Value},
					}
					response := auth_microservice.SessionToken{}

					client.EXPECT().AuthUser(ctx, grpcRequest).Return(
						&auth_microservice.AuthUserResponse{
							Code:     auth_microservice.ErrorCode_COULD_NOT_BUILD_QUERY,
							Response: &response,
						},
						nil,
					)
					return dto.SessionToken{}
				},
			},
			wantErr: true,
			err:     microservice.AuthServiceErrors[auth_microservice.ErrorCode_COULD_NOT_BUILD_QUERY],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			storage := mock_storage.NewMockIAuthStorage(ctrl)
			client := mock_grcp.NewMockAuthServiceClient(ctrl)

			cfg, _ := config.NewSessionConfig()

			want := tt.args.query(tt.args.ctx, *storage, *client, tt.args)

			a := microservice.NewAuthService(*cfg, storage, client)

			got, err := a.AuthUser(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthService.AuthUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("AuthUser() got = %v, want %v", got, want)
			}
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("AuthUser() got = %v, want %v", got, want)
			}
		})
	}
}

/*
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
*/
