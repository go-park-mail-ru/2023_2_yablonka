package service

import (
	"server/internal/config"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	csat "server/internal/service/csat"
	csrf "server/internal/service/csrf"
	list "server/internal/service/list"
	task "server/internal/service/task"
	user "server/internal/service/user"
	workspace "server/internal/service/workspace"
	"server/internal/storage"
)

type Services struct {
	Auth         IAuthService
	User         IUserService
	Board        IBoardService
	CSRF         ICSRFService
	List         IListService
	Task         ITaskService
	Workspace    IWorkspaceService
	CSATQuestion ICSATQuestionService
	CSATAnswer   ICSATSAnswerService
}

func NewServices(storages *storage.Storages, config config.SessionConfig) *Services {
	return &Services{
		Auth:         auth.NewAuthService(config, storages.Auth),
		User:         user.NewUserService(storages.User),
		Board:        board.NewBoardService(storages.Board, storages.User),
		CSRF:         csrf.NewCSRFService(config, storages.CSRF),
		List:         list.NewListService(storages.List),
		Task:         task.NewTaskService(storages.Task),
		Workspace:    workspace.NewWorkspaceService(storages.Workspace),
		CSATAnswer:   csat.NewCSATAnswerService(storages.CSATAnswer),
		CSATQuestion: csat.NewCSATQuestionService(storages.CSATQuestion),
	}
}
