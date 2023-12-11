package microservice

import (
	"context"
	"google.golang.org/grpc"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/storage"
	csat_microservice "server/microservices/csat/csat_answer"
	"testing"
)

func TestCSATAnswerService_Create(t *testing.T) {
	type fields struct {
		storage storage.ICSATAnswerStorage
		client  csat_microservice.CSATAnswerServiceClient
	}
	type args struct {
		ctx  context.Context
		info dto.NewCSATAnswer
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
			cs := CSATAnswerService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			if err := cs.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCSATAnswerService(t *testing.T) {
	type args struct {
		storage    storage.ICSATAnswerStorage
		connection *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want *CSATAnswerService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCSATAnswerService(tt.args.storage, tt.args.connection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCSATAnswerService() = %v, want %v", got, tt.want)
			}
		})
	}
}
