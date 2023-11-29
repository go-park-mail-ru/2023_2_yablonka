package service

import (
	"server/internal/config"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	checklist "server/internal/service/checklist"
	checklist_item "server/internal/service/checklist_item"
	comment "server/internal/service/comment"
	csat "server/internal/service/csat"
	csrf "server/internal/service/csrf"
	list "server/internal/service/list"
	task "server/internal/service/task"
	user "server/internal/service/user"
	workspace "server/internal/service/workspace"
	"server/internal/storage"

	"google.golang.org/grpc"
)

type Services struct {
	Auth          auth.IAuthService
	User          IUserService
	Board         IBoardService
	CSRF          ICSRFService
	List          IListService
	Task          ITaskService
	Comment       ICommentService
	Checklist     IChecklistService
	ChecklistItem IChecklistItemService
	Workspace     IWorkspaceService
	CSATQuestion  ICSATQuestionService
	CSATAnswer    ICSATSAnswerService
}

func NewEmbeddedServices(storages *storage.Storages, config config.SessionConfig) *Services {
	return &Services{
		Auth:          auth.NewEmbeddedAuthService(storages.Auth, config),
		User:          user.NewUserService(storages.User),
		Board:         board.NewEmbeddedBoardService(storages.Board, storages.Task, storages.User, storages.Comment, storages.Checklist, storages.ChecklistItem),
		CSRF:          csrf.NewCSRFService(config, storages.CSRF),
		List:          list.NewListService(storages.List),
		Task:          task.NewEmbeddedTaskService(storages.Task),
		Comment:       comment.NewCommentService(storages.Comment),
		Checklist:     checklist.NewChecklistService(storages.Checklist),
		ChecklistItem: checklist_item.NewChecklistItemService(storages.ChecklistItem),
		Workspace:     workspace.NewWorkspaceService(storages.Workspace),
		CSATAnswer:    csat.NewCSATAnswerService(storages.CSATAnswer),
		CSATQuestion:  csat.NewCSATQuestionService(storages.CSATQuestion),
	}
}

func NewMicroServices(storages *storage.Storages, config config.SessionConfig, conn *grpc.ClientConn) *Services {
	return &Services{
		Auth:  auth.NewMicroAuthService(storages.Auth, config, conn),
		Board: board.NewMicroBoardService(storages.Board, storages.Task, storages.User, storages.Comment, storages.Checklist, storages.ChecklistItem, conn),
		Task:  task.NewMicroTaskService(storages.Task, conn),
	}
}
