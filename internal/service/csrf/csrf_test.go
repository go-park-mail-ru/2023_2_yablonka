package service

import (
	"context"
	"reflect"
	"server/internal/config"
	"server/internal/pkg/dto"
	"server/internal/storage"
	"testing"
	"time"
)

func TestCSRFService_DeleteCSRF(t *testing.T) {
	type fields struct {
		sessionDuration time.Duration
		sessionIDLength uint
		storage         storage.ICSRFStorage
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
			cs := &CSRFService{
				sessionDuration: tt.fields.sessionDuration,
				sessionIDLength: tt.fields.sessionIDLength,
				storage:         tt.fields.storage,
			}
			if got := cs.GetLifetime(); got != tt.want {
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
			if got := NewCSRFService(tt.args.config, tt.args.CSRFStorage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCSRFService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateToken(t *testing.T) {
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
			got, err := generateToken(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
