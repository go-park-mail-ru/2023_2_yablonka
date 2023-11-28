package embedded

import (
	"context"
	"reflect"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	"testing"
	"time"
)

func TestAuthService_AuthUser(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthService{
				sessionDuration: tt.fields.sessionDuration,
				sessionIDLength: tt.fields.sessionIDLength,
				authStorage:     tt.fields.authStorage,
			}
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
		sessionIDLength uint
		authStorage     storage.IAuthStorage
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthService{
				sessionDuration: tt.fields.sessionDuration,
				sessionIDLength: tt.fields.sessionIDLength,
				authStorage:     tt.fields.authStorage,
			}
			if got := a.GetLifetime(context.Background()); got != tt.want {
				t.Errorf("GetLifetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_LogOut(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
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

func TestNewAuthService(t *testing.T) {
	type args struct {
		config      config.SessionConfig
		authStorage storage.IAuthStorage
	}
	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.config, tt.args.authStorage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateString(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateString(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
