package storage

import (
	"database/sql"
	"server/internal/storage/postgresql"
)

type Storages struct {
	Auth          IAuthStorage
	User          IUserStorage
	Board         IBoardStorage
	CSRF          ICSRFStorage
	List          IListStorage
	Task          ITaskStorage
	Comment       ICommentStorage
	Checklist     IChecklistStorage
	ChecklistItem IChecklistItemStorage
	Workspace     IWorkspaceStorage
	CSATAnswer    ICSATAnswerStorage
	CSATQuestion  ICSATQuestionStorage
}

func NewPostgresStorages(db *sql.DB) *Storages {
	return &Storages{
		Auth:          postgresql.NewAuthStorage(db),
		User:          postgresql.NewUserStorage(db),
		Board:         postgresql.NewBoardStorage(db),
		CSRF:          postgresql.NewCSRFStorage(db),
		List:          postgresql.NewListStorage(db),
		Task:          postgresql.NewTaskStorage(db),
		Comment:       postgresql.NewCommentStorage(db),
		Checklist:     postgresql.NewChecklistStorage(db),
		ChecklistItem: postgresql.NewChecklistItemStorage(db),
		Workspace:     postgresql.NewWorkspaceStorage(db),
		CSATAnswer:    postgresql.NewCSATAnswerStorage(db),
		CSATQuestion:  postgresql.NewCSATQuestionStorage(db),
	}
}
