package postgresql

import (
	"context"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
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

// GetUserWorkspaces
// находит рабочие пространства, связанные с пользователем в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetUserWorkspaces(ctx context.Context, userID dto.UserID) (*dto.AllWorkspaces, error) {
	sql, args, err := sq.
		Select(userWorkspaceFields...).
		From("public.workspace").
		Join("public.user_workspace ON public.user_workspace.id_workspace = workspace.id").
		Join("public.user ON public.user_workspace.id_user = user.id").
		Join("public.role ON public.user_workspace.id_role = role.id").
		Where(sq.Eq{"public.user_workspace.id_user": userID.Value}).
		GroupBy("role.id").
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersWithRoles dto.UsersAndRoles
	for rows.Next() {
		var user dto.UserInWorkspace
		var role dto.RoleInWorkspace

		err = rows.Scan(
			&user,
			&role,
		)
		if err != nil {
			return nil, err
		}

		usersWithRoles.Users = append(usersWithRoles.Users, user)
		usersWithRoles.Roles = append(usersWithRoles.Roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

// GetWorkspace
// находит рабочее пространство и связанные доски в БД по его id
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetWorkspace(ctx context.Context, id dto.WorkspaceID) (*entities.Workspace, error) {
	sql, args, err := sq.
		Select(allWorkspaceAndBoardFields...).
		From("public.workspace").
		Join("public.board ON public.workspace.id = public.board.workspace_id").
		Where(sq.Eq{"public.workspace.id": id.Value}).
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
		var board entities.Board

		err = rows.Scan(
			&workspace,
			&board)
		if err != nil {
			return nil, err
		}

		workspace.Boards = append(workspace.Boards, board)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return &workspace, nil
}

// Create
// создает новоt рабочее пространство в БД по данным
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error) {
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
		Insert("workspace_user").
		Columns("id_workspace", "id_user", "id_role").
		Values(workspace.ID, info.OwnerID, 1).
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
}

// UpdateData
// обновляет рабочее пространство в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) UpdateData(ctx context.Context, info dto.UpdatedWorkspaceInfo) error {
	sql, args, err := sq.
		Update("public.workspace").
		Set("name", info.Name).
		Set("description", info.Description).
		Where(sq.Eq{"workspace.id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrUserNotUpdated
	}

	return nil
}

// UpdateUsers
// обновляет людей с доступом в рабочее пространство в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) UpdateUsers(ctx context.Context, info dto.ChangeWorkspaceGuestsInfo) error {
	return nil
}

// Delete
// удаляет данногt рабочее пространство в БД по id
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Delete(ctx context.Context, id dto.WorkspaceID) error {
	sql, args, err := sq.
		Delete("public.workspace").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrUserNotDeleted
	}

	return nil
}
