package service

import (
	"context"
	"google.golang.org/grpc"
	"reflect"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	csrf_microservice "server/microservices/csrf/csrf"
	"testing"
	"time"
)

func TestCSRFService_DeleteCSRF(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		sessionIDLength uint
		storage         storage.ICSRFStorage
		client          csrf_microservice.CSRFServiceClient
	}
	type args struct {
		ctx   context.Context
		token dto.CSRFToken
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
			cs := &CSRFService{
				sessionDuration: tt.fields.sessionDuration,
				sessionIDLength: tt.fields.sessionIDLength,
				storage:         tt.fields.storage,
				client:          tt.fields.client,
			}
			if err := cs.DeleteCSRF(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCSRF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSRFService_GetLifetime(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		sessionIDLength uint
		storage         storage.ICSRFStorage
		client          csrf_microservice.CSRFServiceClient
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
			cs := &CSRFService{
				sessionDuration: tt.fields.sessionDuration,
				sessionIDLength: tt.fields.sessionIDLength,
				storage:         tt.fields.storage,
				client:          tt.fields.client,
			}
			if got := cs.GetLifetime(tt.args.ctx); got != tt.want {
				t.Errorf("GetLifetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSRFService_SetupCSRF(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		sessionIDLength uint
		storage         storage.ICSRFStorage
		client          csrf_microservice.CSRFServiceClient
	}
	type args struct {
		ctx context.Context
		id  dto.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.CSRFData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CSRFService{
				sessionDuration: tt.fields.sessionDuration,
				sessionIDLength: tt.fields.sessionIDLength,
				storage:         tt.fields.storage,
				client:          tt.fields.client,
			}
			got, err := cs.SetupCSRF(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetupCSRF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetupCSRF() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSRFService_VerifyCSRF(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		sessionIDLength uint
		storage         storage.ICSRFStorage
		client          csrf_microservice.CSRFServiceClient
	}
	type args struct {
		ctx   context.Context
		token dto.CSRFToken
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
			cs := &CSRFService{
				sessionDuration: tt.fields.sessionDuration,
				sessionIDLength: tt.fields.sessionIDLength,
				storage:         tt.fields.storage,
				client:          tt.fields.client,
			}
			if err := cs.VerifyCSRF(tt.args.ctx, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("VerifyCSRF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCSRFService(t *testing.T) {
	type args struct {
		config      config.SessionConfig
		CSRFStorage storage.ICSRFStorage
		conn        *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want *CSRFService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCSRFService(tt.args.config, tt.args.CSRFStorage, tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCSRFService() = %v, want %v", got, tt.want)
			}
		})
	}
}
