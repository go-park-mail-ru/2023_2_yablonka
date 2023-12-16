package microservice

import (
	"context"
	"google.golang.org/grpc"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/storage"
	csat_microservice "server/microservices/csat/csat_question"
	"testing"
)

func TestCSATQuestionService_CheckRating(t *testing.T) {
	type fields struct {
		storage storage.ICSATQuestionStorage
		client  csat_microservice.CSATQuestionServiceClient
	}
	type args struct {
		ctx  context.Context
		info dto.NewCSATAnswerInfo
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
			cs := CSATQuestionService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			if err := cs.CheckRating(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("CheckRating() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSATQuestionService_Create(t *testing.T) {
	type fields struct {
		storage storage.ICSATQuestionStorage
		client  csat_microservice.CSATQuestionServiceClient
	}
	type args struct {
		ctx  context.Context
		info dto.NewCSATQuestionInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.CSATQuestionFull
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CSATQuestionService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			got, err := cs.Create(tt.args.ctx, tt.args.info)
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

func TestCSATQuestionService_Delete(t *testing.T) {
	type fields struct {
		storage storage.ICSATQuestionStorage
		client  csat_microservice.CSATQuestionServiceClient
	}
	type args struct {
		ctx context.Context
		id  dto.CSATQuestionID
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
			cs := CSATQuestionService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			if err := cs.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCSATQuestionService_GetAll(t *testing.T) {
	type fields struct {
		storage storage.ICSATQuestionStorage
		client  csat_microservice.CSATQuestionServiceClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.CSATQuestionFull
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CSATQuestionService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			got, err := cs.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSATQuestionService_GetStats(t *testing.T) {
	type fields struct {
		storage storage.ICSATQuestionStorage
		client  csat_microservice.CSATQuestionServiceClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]dto.QuestionWithStats
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CSATQuestionService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			got, err := cs.GetStats(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCSATQuestionService_Update(t *testing.T) {
	type fields struct {
		storage storage.ICSATQuestionStorage
		client  csat_microservice.CSATQuestionServiceClient
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedCSATQuestionInfo
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
			cs := CSATQuestionService{
				storage: tt.fields.storage,
				client:  tt.fields.client,
			}
			if err := cs.Update(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCSATQuestionService(t *testing.T) {
	type args struct {
		storage    storage.ICSATQuestionStorage
		connection *grpc.ClientConn
	}
	tests := []struct {
		name string
		args args
		want *CSATQuestionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCSATQuestionService(tt.args.storage, tt.args.connection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCSATQuestionService() = %v, want %v", got, tt.want)
			}
		})
	}
}
