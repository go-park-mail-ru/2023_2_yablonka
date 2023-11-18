package service

import (
	"server/internal/config"
	auth "server/internal/service/auth"
	board "server/internal/service/board"
	csrf "server/internal/service/csrf"
	list "server/internal/service/list"
	task "server/internal/service/task"
	user "server/internal/service/user"
	workspace "server/internal/service/workspace"
	"server/internal/storage"
)

type Services struct {
	Auth      IAuthService
	User      IUserService
	Board     IBoardService
	CSRF      ICSRFService
	List      IListService
	Task      ITaskService
	Workspace IWorkspaceService
}

func NewServices(storages *storage.Storages, config config.SessionConfig) *Services {
	return &Services{
		Auth:      auth.NewAuthService(config, storages.Auth),
		User:      user.NewUserService(storages.User),
		Board:     board.NewBoardService(storages.Board),
		CSRF:      csrf.NewCSRFService(config, storages.CSRF),
		List:      list.NewListService(storages.List),
		Task:      task.NewTaskService(storages.Task),
		Workspace: workspace.NewWorkspaceService(storages.Workspace),
	}
}