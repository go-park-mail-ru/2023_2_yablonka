package postgresql

import (
	"context"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresWorkspaceStorage
// Хранилище данных в PostgreSQL
type PostgresWorkspaceStorage struct {
	db *pgxpool.Pool
}

// NewWorkspaceStorage
// возвращает PostgreSQL хранилище рабочих пространств
func NewWorkspaceStorage(db *pgxpool.Pool) *PostgresWorkspaceStorage {
	return &PostgresWorkspaceStorage{
		db: db,
	}
}

// Create
// создает новоt рабочее пространство в БД по данным
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error) {
	return nil, nil
}

// GetUserWorkspaces
// находит пользователя в БД по почте
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetUserWorkspaces(ctx context.Context, userID uint64) (*[]entities.Workspace, error) {
	return nil, nil
}

// GetByID
// находит рабочее пространство в БД по его id
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetByID(ctx context.Context, workspaceID uint64) (*entities.Workspace, error) {
	return nil, nil
}

// Update
// обновляет рабочее пространство в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Update(ctx context.Context, info dto.UpdatedWorkspaceInfo) (*entities.Workspace, error) {
	return nil, nil
}

// Delete
// удаляет данногt рабочее пространство в БД по id
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Delete(ctx context.Context, id uint64) error {
	return nil
}
