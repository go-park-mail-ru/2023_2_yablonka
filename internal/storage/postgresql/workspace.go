package postgresql

import (
	"context"
	"fmt"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresWorkspaceStorage
// Хранилище данных в PostgreSQL
type PostgresWorkspaceStorage struct {
	db *pgxpool.Pool
}

type GuestWorkspaceReturn struct {
	WorkspaceID          uint64
	WorkspaceName        string
	WorkspaceDateCreated time.Time
	dto.UserOwnerInfo
}

type BoardReturn struct {
	WorkspaceID uint64
	dto.WorkspaceBoardInfo
}

// NewWorkspaceStorage
// возвращает PostgreSQL хранилище рабочих пространств
func NewWorkspaceStorage(db *pgxpool.Pool) *PostgresWorkspaceStorage {
	return &PostgresWorkspaceStorage{
		db: db,
	}
}

// GetUserOwnedWorkspaces
// находит рабочие пространства, связанные с пользователем в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetUserOwnedWorkspaces(ctx context.Context, userID dto.UserID) (*[]dto.UserOwnedWorkspaceInfo, error) {
	workspaceQuery, args, err := sq.
		Select("id", "name").
		From("public.workspace").
		Where(sq.Eq{"id_creator": userID.Value}).
		OrderBy("public.workspace.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built user owned workspaces query\n\t", workspaceQuery, "\nwith args\n\t", args)

	rows, err := s.db.Query(context.Background(), workspaceQuery, args...)
	if err != nil {
		log.Println("Storage -- DB workspaces query failed with error", err.Error())
		return nil, err
	}
	log.Println("Workspaces got")
	defer rows.Close()

	// workspaceQuery, args, err := sq.
	// 	Select(allWorkspaceFields...).
	// 	From("public.workspace").
	// 	Join("public.user_workspace ON public.user_workspace.id_workspace = public.workspace.id").
	// 	Join("public.user ON public.user_workspace.id_user = public.user.id").
	// 	Join("public.board ON public.board.id_workspace = public.workspace.id").
	// 	Join("public.board_user ON public.board_user.id_board = public.board.id").
	// 	Where(sq.Eq{"public.user_workspace.id_user": userID.Value}).
	// 	PlaceholderFormat(sq.Dollar).
	// 	ToSql()

	// var workspaces []dto.UserOwnedWorkspaceInfo
	workspaces := map[uint64]dto.UserOwnedWorkspaceInfo{}
	var ownedID []uint64
	for rows.Next() {
		var workspace dto.UserOwnedWorkspaceInfo

		err = rows.Scan(
			&workspace.ID,
			&workspace.Name,
		)
		if err != nil {
			fmt.Println("Scanning failed due to error", err.Error())
			return nil, err
		}
		workspaces[workspace.ID] = workspace
		ownedID = append(ownedID, workspace.ID)
	}

	boardQuery, args, err := sq.
		Select("id_workspace", "id", "name", "description").
		From("public.board").
		Where(sq.Eq{"id_workspace": ownedID}).
		OrderBy("public.board.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built boards query\n\t", boardQuery, "\nwith args\n\t", args)

	rows, err = s.db.Query(context.Background(), boardQuery, args...)
	if err != nil {
		log.Println("Storage -- DB boards query failed with error", err.Error())
		return nil, err
	}
	log.Println("Boards got")
	defer rows.Close()

	boardRows, err := pgx.CollectRows(rows, pgx.RowToStructByPos[BoardReturn])
	if err != nil {
		fmt.Println("Failed collecting boards with error:", err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	log.Println("Boards collected")

	for _, boardRow := range boardRows {
		board := dto.WorkspaceBoardInfo{
			ID:          boardRow.ID,
			Name:        boardRow.Name,
			Description: boardRow.Description,
		}
		ws := workspaces[boardRow.WorkspaceID]
		ws.Boards = append(ws.Boards, board)
		workspaces[boardRow.WorkspaceID] = ws
	}
	log.Println("Boards appended to workspaces")

	var result []dto.UserOwnedWorkspaceInfo
	for _, value := range workspaces {
		result = append(result, value)
	}

	return &result, nil
}

// GetUserGuestWorkspaces
// находит рабочие пространства, связанные с пользователем в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetUserGuestWorkspaces(ctx context.Context, userID dto.UserID) (*[]dto.UserGuestWorkspaceInfo, error) {
	workspaceQuery, args, err := sq.
		Select(userGuestWorkspaceFields...).
		Distinct().
		From("public.workspace").
		Join("public.user_workspace ON public.user_workspace.id_user = " + strconv.FormatUint(uint64(userID.Value), 10)).
		Join("public.user ON public.user.id = public.workspace.id_creator").
		Where(sq.NotEq{"public.workspace.id_creator": userID.Value}).
		OrderBy("public.workspace.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built user guest workspaces query\n\t", workspaceQuery, "\nwith args\n\t", args)

	rows, err := s.db.Query(context.Background(), workspaceQuery, args...)
	if err != nil {
		log.Println("Storage -- DB workspaces query failed with error", err.Error())
		return nil, err
	}
	log.Println("Workspaces got")
	defer rows.Close()

	workspaceRows, err := pgx.CollectRows(rows, pgx.RowToStructByPos[GuestWorkspaceReturn])
	if err != nil {
		fmt.Println("Failed collecting boards with error:", err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}

	workspaces := map[uint64]dto.UserGuestWorkspaceInfo{}
	var guestID []uint64
	for _, row := range workspaceRows {
		workspaces[row.WorkspaceID] = dto.UserGuestWorkspaceInfo{
			ID:   row.WorkspaceID,
			Name: row.WorkspaceName,
			Owner: dto.UserOwnerInfo{
				ID:      row.ID,
				Email:   row.Email,
				Name:    row.Name,
				Surname: row.Surname,
			},
		}
		guestID = append(guestID, row.WorkspaceID)
	}

	boardQuery, args, err := sq.
		Select("id_workspace", "id", "name", "description").
		From("public.board").
		Where(sq.Eq{"id_workspace": guestID}).
		OrderBy("public.board.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built boards query\n\t", boardQuery, "\nwith args\n\t", args)

	rows, err = s.db.Query(context.Background(), boardQuery, args...)
	if err != nil {
		log.Println("Storage -- DB boards query failed with error", err.Error())
		return nil, err
	}
	log.Println("Boards got")
	defer rows.Close()

	boardRows, err := pgx.CollectRows(rows, pgx.RowToStructByPos[BoardReturn])
	if err != nil {
		fmt.Println("Failed collecting boards with error:", err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	log.Println("Boards collected")

	for _, boardRow := range boardRows {
		board := dto.WorkspaceBoardInfo{
			ID:          boardRow.ID,
			Name:        boardRow.Name,
			Description: boardRow.Description,
		}
		ws := workspaces[boardRow.WorkspaceID]
		ws.Boards = append(ws.Boards, board)
		workspaces[boardRow.WorkspaceID] = ws
	}
	log.Println("Boards appended to workspaces")

	var result []dto.UserGuestWorkspaceInfo
	for _, workspace := range workspaces {
		result = append(result, workspace)
	}

	return &result, nil
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
		Columns("name", "description", "id_creator").
		Values(info.Name, info.Description, info.OwnerID).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()

	if err != nil {
		fmt.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built workspace query\n\t", query1, "\nwith args\n\t", args)

	workspace := entities.Workspace{
		Name:        *info.Name,
		Description: info.Description,
		Boards:      []entities.Board{},
	}

	row := tx.QueryRow(ctx, query1, args...)
	log.Println("Storage -- DB queried")
	if err := row.Scan(&workspace.ID, &workspace.DateCreated); err != nil {
		log.Println("Storage -- Workspace insert failed with error", err.Error())
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	log.Println("Storage -- Workspace created")

	query2, args, err := sq.
		Insert("user_workspace").
		Columns("id_workspace", "id_user", "id_role").
		Values(workspace.ID, info.OwnerID, 1).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built workspace user query\n\t", query2, "\nwith args\n\t", args)

	_, err = tx.Exec(ctx, query2, args...)

	if err != nil {
		log.Println("Storage -- Workspace user insert failed with error", err.Error())
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	log.Println("Storage -- Committing changes")
	err = tx.Commit(ctx)
	if err != nil {
		log.Println("Storage -- Failed to commit changes")
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	log.Println("Storage -- Committed changes")

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
		log.Println("Storage -- Couldn't build query")
		return apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built workspace DELETE query\n\t", sql, "\nwith args\n\t", args)

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		log.Println("Failed to delete workspace with error", err.Error())
		return apperrors.ErrWorkspaceNotDeleted
	}

	return nil
}
