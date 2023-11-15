package storage

import (
	"server/internal/storage/postgresql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storages struct {
	Auth      IAuthStorage
	User      IUserStorage
	Board     IBoardStorage
	CSRF      ICSRFStorage
	List      IListStorage
	Task      ITaskStorage
	Workspace IWorkspaceStorage
}

func NewPostgresStorages(dbConnection *pgxpool.Pool) *Storages {
	return &Storages{
		Auth:      postgresql.NewAuthStorage(dbConnection),
		User:      postgresql.NewUserStorage(dbConnection),
		Board:     postgresql.NewBoardStorage(dbConnection),
		CSRF:      postgresql.NewCSRFStorage(dbConnection),
		List:      postgresql.NewListStorage(dbConnection),
		Task:      postgresql.NewTaskStorage(dbConnection),
		Workspace: postgresql.NewWorkspaceStorage(dbConnection),
	}
}
