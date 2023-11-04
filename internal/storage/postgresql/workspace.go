package postgresql

import (
	"context"
<<<<<<< Updated upstream
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

=======
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
	return nil, nil
=======
	query1, args, err := sq.
		Insert("public.workspace").
		Columns("name", "thumbnail_url", "description").
		Values(info.Name, info.ThumbnailURL, info.Description).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	workspace := entities.Workspace{
		Name:         info.Name,
		Description:  info.Description,
		ThumbnailURL: info.ThumbnailURL,
	}

	row := tx.QueryRow(ctx, query1, args...)
	if row.Scan(&workspace.ID, &workspace.DateCreated) != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	query2, args, err := sq.
		Insert("public.workspace_user").
		Columns("id_workspace", "id_user").
		Values(workspace.ID, info.OwnerID).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	_, err = tx.Exec(ctx, query2, args...)

	if err != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	return &workspace, nil
>>>>>>> Stashed changes
}

// GetUserWorkspaces
// находит пользователя в БД по почте
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetUserWorkspaces(ctx context.Context, userID uint64) (*[]entities.Workspace, error) {
<<<<<<< Updated upstream
	return nil, nil
=======
	sql, args, err := sq.
		Select(allWorkspaceFields...).
		From("public.workspace").
		Join("public.workspace_user").
		Where(sq.Eq{"public.workspace_user.id_user": userID}).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(ctx, sql, args...)

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	defer rows.Close()

	var workspaces []entities.Workspace
	for rows.Next() {
		var workspace entities.Workspace
		err := rows.Scan(
			&workspace.ID,
			&workspace.Name,
			&workspace.Description,
			&workspace.DateCreated,
			&workspace.ThumbnailURL,
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetWorkspace
		}
		workspaces = append(workspaces, workspace)
	}
	if rows.Err() != nil {
		return nil, apperrors.ErrCouldNotGetWorkspace
	}

	return &workspaces, nil
>>>>>>> Stashed changes
}

// GetByID
// находит рабочее пространство в БД по его id
// или возвращает ошибки ...
<<<<<<< Updated upstream
func (s PostgresWorkspaceStorage) GetByID(ctx context.Context, workspaceID uint64) (*entities.Workspace, error) {
	return nil, nil
=======
func (s PostgresWorkspaceStorage) GetByID(ctx context.Context, id uint64) (*entities.Workspace, error) {
	sql, args, err := sq.Select("workspaces.*", "boards.*", "lists.*", "tasks.*", "users.*").
		From("workspaces").
		Join("boards ON workspaces.id = boards.workspace_id").
		Join("lists ON boards.id = lists.board_id").
		Join("tasks ON lists.id = tasks.list_id").
		Join("user_workspace ON workspaces.id = user_workspace.workspace_id").
		Join("users ON user_workspace.user_id = users.id").
		Where(sq.Eq{"workspaces.id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workspace entities.Workspace
	for rows.Next() {
		err = rows.Scan(&workspace.ID, &workspace.Name, &workspace.Description, &workspace.ThumbnailURL, &workspace.DateCreated, &workspace.Boards, &workspace.Users)
		if err != nil {
			return nil, err
		}
	}

	return &workspace, nil
>>>>>>> Stashed changes
}

// Update
// обновляет рабочее пространство в БД
// или возвращает ошибки ...
<<<<<<< Updated upstream
func (s PostgresWorkspaceStorage) Update(ctx context.Context, info dto.UpdatedWorkspaceInfo) (*entities.Workspace, error) {
	return nil, nil
=======
func (s PostgresWorkspaceStorage) Update(ctx context.Context, info dto.UpdatedWorkspaceInfo) error {
	sql, args, err := sq.
		Update("public.workspace").
		Set("name", info.Name).
		Set("description", info.Description).
		Set("thumbnail_url", info.ThumbnailURL).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	query := s.db.QueryRow(ctx, sql, args...)

	if query.Scan() != nil {
		return apperrors.ErrUserNotUpdated
	}

	return nil
>>>>>>> Stashed changes
}

// Delete
// удаляет данногt рабочее пространство в БД по id
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Delete(ctx context.Context, id uint64) error {
<<<<<<< Updated upstream
=======
	sql, args, err := sq.
		Delete("public.workspace").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	query := s.db.QueryRow(ctx, sql, args...)

	if query.Scan() != nil {
		return apperrors.ErrUserNotDeleted
	}

>>>>>>> Stashed changes
	return nil
}
