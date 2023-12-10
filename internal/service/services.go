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
	Auth          IAuthService
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

func NewMicroServices(storages *storage.Storages, config config.SessionConfig, conn *grpc.ClientConn) *Services {
	return &Services{
		Auth:          auth.NewMicroAuthService(storages.Auth, config, conn),
		Board:         board.NewMicroBoardService(storages.Board, storages.Task, storages.User, storages.Comment, storages.Checklist, storages.ChecklistItem, conn),
		Comment:       comment.NewMicroCommentService(storages.Comment, conn),
		Checklist:     checklist.NewMicroChecklistService(storages.Checklist, conn),
		ChecklistItem: checklist_item.NewMicroChecklistItemService(storages.ChecklistItem, conn),
		CSATAnswer:    csat.NewMicroCSATAnswerService(storages.CSATAnswer, conn),
		CSATQuestion:  csat.NewMicroCSATQuestionService(storages.CSATQuestion, conn),
		CSRF:          csrf.NewMicroCSRFService(storages.CSRF, config, conn),
		List:          list.NewMicroListService(storages.List, conn),
		Task:          task.NewMicroTaskService(storages.Task, storages.User, conn),
		User:          user.NewMicroUserService(storages.User, conn),
		Workspace:     workspace.NewMicroWorkspaceService(storages.Workspace, conn),
	}
}
