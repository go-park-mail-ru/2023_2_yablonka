package service

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"testing"
)

func TestBoardService_Create(t *testing.T) {
	type fields struct {
		us storage.IUserStorage
		bs storage.IBoardStorage
	}
	type args struct {
		ctx   context.Context
		board dto.NewBoardInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Board
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := BoardService{
				boardStorage: tt.fields.bs,
				userStorage:  tt.fields.us,
			}
			got, err := bs.Create(tt.args.ctx, tt.args.board)
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

func TestBoardService_Delete(t *testing.T) {
	type fields struct {
		us storage.IUserStorage
		bs storage.IBoardStorage
	}
	type args struct {
		ctx context.Context
		id  dto.BoardID
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
			bs := BoardService{
				boardStorage: tt.fields.bs,
				userStorage:  tt.fields.us,
			}
			if err := bs.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoardService_GetFullBoard(t *testing.T) {
	type fields struct {
		us storage.IUserStorage
		bs storage.IBoardStorage
	}
	type args struct {
		ctx  context.Context
		info dto.IndividualBoardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Board
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := BoardService{
				boardStorage: tt.fields.bs,
				userStorage:  tt.fields.us,
			}
			got, err := bs.GetFullBoard(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFullBoard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardService_UpdateData(t *testing.T) {
	type fields struct {
		us storage.IUserStorage
		bs storage.IBoardStorage
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedBoardInfo
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
			bs := BoardService{
				boardStorage: tt.fields.bs,
				userStorage:  tt.fields.us,
			}
			if err := bs.UpdateData(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoardService_UpdateThumbnail(t *testing.T) {
	type fields struct {
		us storage.IUserStorage
		bs storage.IBoardStorage
	}
	type args struct {
		ctx  context.Context
		info dto.UpdatedBoardThumbnailInfo
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
			bs := BoardService{
				boardStorage: tt.fields.bs,
				userStorage:  tt.fields.us,
			}
			got, err := bs.UpdateThumbnail(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateThumbnail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateThumbnail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoardService(t *testing.T) {
	type args struct {
		us storage.IUserStorage
		ls storage.IListStorage
		bs storage.IBoardStorage
	}
	tests := []struct {
		name string
		args args
		want *BoardService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoardService(tt.args.bs, tt.args.ls, tt.args.us); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoardService() = %v, want %v", got, tt.want)
			}
		})
	}
}
