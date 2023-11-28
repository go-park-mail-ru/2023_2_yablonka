package service

import (
	"server/internal/config"
	"server/internal/service/auth"
	"server/internal/service/board"
	"server/internal/service/comment"
	"server/internal/service/csat"
	"server/internal/service/csrf"
	"server/internal/service/list"
	"server/internal/service/task"
	"server/internal/service/user"
	"server/internal/service/workspace"
	"server/internal/storage"

	"google.golang.org/grpc"
)

type Services struct {
	Auth         auth.IAuthService
	User         user.IUserService
	Board        board.IBoardService
	Comment      comment.ICommentService
	CSRF         csrf.ICSRFService
	List         list.IListService
	Task         task.ITaskService
	Workspace    workspace.IWorkspaceService
	CSATQuestion csat.ICSATQuestionService
	CSATAnswer   csat.ICSATAnswerService
}

func NewEmbeddedServices(storages *storage.Storages, config config.SessionConfig) *Services {
	return &Services{
		Auth:  auth.NewEmbeddedAuthService(storages.Auth, config),
		Board: board.NewEmbeddedBoardService(storages.Board, storages.List, storages.User),
		// Comment:      comment.NewEmbeddedCommentService(storages.Comment),
		CSATAnswer:   csat.NewEmbeddedCSATAnswerService(storages.CSATAnswer),
		CSATQuestion: csat.NewEmbeddedCSATQuestionService(storages.CSATQuestion),
		CSRF:         csrf.NewEmbeddedCSRFService(storages.CSRF, config),
		List:         list.NewEmbeddedListService(storages.List),
		Task:         task.NewEmbeddedTaskService(storages.Task),
		User:         user.NewEmbeddedUserService(storages.User),
		Workspace:    workspace.NewEmbeddedWorkspaceService(storages.Workspace),
	}
}

func NewMicroServices(storages *storage.Storages, config config.SessionConfig, connection *grpc.ClientConn) *Services {
	return &Services{
		Auth:  auth.NewMicroAuthService(storages.Auth, config, connection),
		Board: board.NewMicroBoardService(storages.Board, storages.List, storages.User, connection),
		// Comment:      comment.NewMicroCommentService(storages.Comment, connection),
		CSATAnswer:   csat.NewMicroCSATAnswerService(storages.CSATAnswer, connection),
		CSATQuestion: csat.NewMicroCSATQuestionService(storages.CSATQuestion, connection),
		CSRF:         csrf.NewMicroCSRFService(storages.CSRF, config, connection),
		List:         list.NewMicroListService(storages.List, connection),
		Task:         task.NewMicroTaskService(storages.Task, connection),
		User:         user.NewMicroUserService(storages.User, connection),
		Workspace:    workspace.NewMicroWorkspaceService(storages.Workspace, connection),
	}
}
