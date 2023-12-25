package microservice

import (
	"context"
	"reflect"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"server/internal/storage"
	"testing"

	"google.golang.org/grpc"
)

func TestBoardService_AddUser(t *testing.T) {
	type fields struct {
		boardStorage         storage.IBoardStorage
		userStorage          storage.IUserStorage
		taskStorage          storage.ITaskStorage
		commentStorage       storage.ICommentStorage
		checklistStorage     storage.IChecklistStorage
		checklistItemStorage storage.IChecklistItemStorage
	}
	type args struct {
		ctx     context.Context
		request dto.AddBoardUserRequest
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
				boardStorage:         tt.fields.boardStorage,
				userStorage:          tt.fields.userStorage,
				taskStorage:          tt.fields.taskStorage,
				commentStorage:       tt.fields.commentStorage,
				checklistStorage:     tt.fields.checklistStorage,
				checklistItemStorage: tt.fields.checklistItemStorage,
			}
			if _, err := bs.AddUser(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoardService_Create(t *testing.T) {
	type fields struct {
		boardStorage         storage.IBoardStorage
		userStorage          storage.IUserStorage
		taskStorage          storage.ITaskStorage
		commentStorage       storage.ICommentStorage
		checklistStorage     storage.IChecklistStorage
		checklistItemStorage storage.IChecklistItemStorage
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
				boardStorage:         tt.fields.boardStorage,
				userStorage:          tt.fields.userStorage,
				taskStorage:          tt.fields.taskStorage,
				commentStorage:       tt.fields.commentStorage,
				checklistStorage:     tt.fields.checklistStorage,
				checklistItemStorage: tt.fields.checklistItemStorage,
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
		boardStorage         storage.IBoardStorage
		userStorage          storage.IUserStorage
		taskStorage          storage.ITaskStorage
		commentStorage       storage.ICommentStorage
		checklistStorage     storage.IChecklistStorage
		checklistItemStorage storage.IChecklistItemStorage
	}
	type args struct {
		ctx  context.Context
		info dto.BoardDeleteRequest
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
				boardStorage:         tt.fields.boardStorage,
				userStorage:          tt.fields.userStorage,
				taskStorage:          tt.fields.taskStorage,
				commentStorage:       tt.fields.commentStorage,
				checklistStorage:     tt.fields.checklistStorage,
				checklistItemStorage: tt.fields.checklistItemStorage,
			}
			if err := bs.Delete(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoardService_GetFullBoard(t *testing.T) {
	type fields struct {
		boardStorage         storage.IBoardStorage
		userStorage          storage.IUserStorage
		taskStorage          storage.ITaskStorage
		commentStorage       storage.ICommentStorage
		checklistStorage     storage.IChecklistStorage
		checklistItemStorage storage.IChecklistItemStorage
	}
	type args struct {
		ctx  context.Context
		info dto.IndividualBoardRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.FullBoardResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := BoardService{
				boardStorage:         tt.fields.boardStorage,
				userStorage:          tt.fields.userStorage,
				taskStorage:          tt.fields.taskStorage,
				commentStorage:       tt.fields.commentStorage,
				checklistStorage:     tt.fields.checklistStorage,
				checklistItemStorage: tt.fields.checklistItemStorage,
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

func TestBoardService_RemoveUser(t *testing.T) {
	type fields struct {
		boardStorage         storage.IBoardStorage
		userStorage          storage.IUserStorage
		taskStorage          storage.ITaskStorage
		commentStorage       storage.ICommentStorage
		checklistStorage     storage.IChecklistStorage
		checklistItemStorage storage.IChecklistItemStorage
	}
	type args struct {
		ctx  context.Context
		info dto.RemoveBoardUserInfo
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
				boardStorage:         tt.fields.boardStorage,
				userStorage:          tt.fields.userStorage,
				taskStorage:          tt.fields.taskStorage,
				commentStorage:       tt.fields.commentStorage,
				checklistStorage:     tt.fields.checklistStorage,
				checklistItemStorage: tt.fields.checklistItemStorage,
			}
			if err := bs.RemoveUser(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoardService_UpdateData(t *testing.T) {
	type fields struct {
		boardStorage         storage.IBoardStorage
		userStorage          storage.IUserStorage
		taskStorage          storage.ITaskStorage
		commentStorage       storage.ICommentStorage
		checklistStorage     storage.IChecklistStorage
		checklistItemStorage storage.IChecklistItemStorage
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
				boardStorage:         tt.fields.boardStorage,
				userStorage:          tt.fields.userStorage,
				taskStorage:          tt.fields.taskStorage,
				commentStorage:       tt.fields.commentStorage,
				checklistStorage:     tt.fields.checklistStorage,
				checklistItemStorage: tt.fields.checklistItemStorage,
			}
			if err := bs.UpdateData(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("UpdateData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoardService_UpdateThumbnail(t *testing.T) {
	type fields struct {
		boardStorage         storage.IBoardStorage
		userStorage          storage.IUserStorage
		taskStorage          storage.ITaskStorage
		commentStorage       storage.ICommentStorage
		checklistStorage     storage.IChecklistStorage
		checklistItemStorage storage.IChecklistItemStorage
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
				boardStorage:         tt.fields.boardStorage,
				userStorage:          tt.fields.userStorage,
				taskStorage:          tt.fields.taskStorage,
				commentStorage:       tt.fields.commentStorage,
				checklistStorage:     tt.fields.checklistStorage,
				checklistItemStorage: tt.fields.checklistItemStorage,
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
		bs   storage.IBoardStorage
		ts   storage.ITaskStorage
		us   storage.IUserStorage
		cs   storage.ICommentStorage
		cls  storage.IChecklistStorage
		clis storage.IChecklistItemStorage
		conn *grpc.ClientConn
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
			if got := NewBoardService(tt.args.bs, tt.args.ts, tt.args.us, tt.args.cs, tt.args.cls, tt.args.clis, tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoardService() = %v, want %v", got, tt.want)
			}
		})
	}
}
