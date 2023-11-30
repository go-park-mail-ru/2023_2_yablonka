package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// PostgresWorkspaceStorage
// Хранилище данных в PostgreSQL
type PostgresWorkspaceStorage struct {
	db *sql.DB
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
func NewWorkspaceStorage(db *sql.DB) *PostgresWorkspaceStorage {
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

	rows, err := s.db.Query(workspaceQuery, args...)
	if err != nil {
		log.Println("Storage -- DB workspaces query failed with error", err.Error())
		return nil, err
	}
	defer rows.Close()
	log.Println("Workspaces got")

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
		Select("id_workspace", "id", "name", "description", "thumbnail_url").
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

	rows, err = s.db.Query(boardQuery, args...)
	if err != nil {
		log.Println("Storage -- DB boards query failed with error", err.Error())
		return nil, err
	}
	log.Println("Boards got")
	defer rows.Close()

	boards := []BoardReturn{}
	for rows.Next() {
		var board BoardReturn

		err = rows.Scan(
			&board.WorkspaceID,
			&board.ID,
			&board.Name,
			&board.Description,
			&board.ThumbnailURL,
		)
		if err != nil {
			fmt.Println("Scanning failed due to error", err.Error())
			return nil, err
		}
		boards = append(boards, board)
	}

	for _, boardRow := range boards {
		board := dto.WorkspaceBoardInfo{
			ID:           boardRow.ID,
			Name:         boardRow.Name,
			Description:  boardRow.Description,
			ThumbnailURL: boardRow.ThumbnailURL,
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
		Join("public.user_workspace ON public.user_workspace.id_workspace = public.workspace.id").
		Join("public.user ON public.user.id = public.workspace.id_creator").
		Where(sq.And{
			sq.NotEq{"public.workspace.id_creator": userID.Value},
			sq.Eq{"public.user_workspace.id_user": userID.Value},
		}).
		OrderBy("public.workspace.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built user guest workspaces query\n\t", workspaceQuery, "\nwith args\n\t", args)

	rows, err := s.db.Query(workspaceQuery, args...)
	if err != nil {
		log.Println("Storage -- DB workspaces query failed with error", err.Error())
		return nil, err
	}
	log.Println("Workspaces got")
	defer rows.Close()

	workspaceRows := []dto.UserGuestWorkspaceInfo{}
	guestWorkspaceID := []uint64{}
	for rows.Next() {
		var workspace dto.UserGuestWorkspaceInfo
		var owner dto.UserOwnerInfo

		err = rows.Scan(
			&workspace.ID,
			&workspace.Name,
			&workspace.DateCreated,
			&owner.ID,
			&owner.Email,
			&owner.Name,
			&owner.Surname,
		)
		if err != nil {
			fmt.Println("Scanning failed due to error", err.Error())
			return nil, err
		}
		workspace.Owner = owner
		workspaceRows = append(workspaceRows, workspace)
		guestWorkspaceID = append(guestWorkspaceID, workspace.ID)
	}

	workspaces := map[uint64]dto.UserGuestWorkspaceInfo{}
	for _, row := range workspaceRows {
		workspaces[row.ID] = row
	}

	boardQuery, args, err := sq.
		Select("public.board.id_workspace", "public.board.id", "public.board.name", "public.board.description", "public.board.thumbnail_url").
		From("public.board").
		LeftJoin("public.board_user ON public.board_user.id_board = public.board.id").
		Where(sq.And{
			sq.Eq{"public.board.id_workspace": guestWorkspaceID},
			sq.Eq{"public.board_user.id_user": userID.Value},
		}).
		OrderBy("public.board.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built boards query\n\t", boardQuery, "\nwith args\n\t", args)

	rows, err = s.db.Query(boardQuery, args...)
	if err != nil {
		log.Println("Storage -- DB boards query failed with error", err.Error())
		return nil, err
	}
	log.Println("Boards got")
	defer rows.Close()

	boardRows := []BoardReturn{}
	for rows.Next() {
		var board BoardReturn
		err = rows.Scan(
			&board.WorkspaceID,
			&board.ID,
			&board.Name,
			&board.Description,
			&board.ThumbnailURL,
		)
		if err != nil {
			fmt.Println("Scanning failed due to error", err.Error())
			return nil, err
		}
		boardRows = append(boardRows, board)
	}
	log.Println("Boards collected")

	for _, boardRow := range boardRows {
		board := dto.WorkspaceBoardInfo{
			ID:           boardRow.ID,
			Name:         boardRow.Name,
			Description:  boardRow.Description,
			ThumbnailURL: boardRow.ThumbnailURL,
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

	rows, err := s.db.Query(sql, args...)
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
	user := ctx.Value(dto.UserObjKey).(*entities.User)

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

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built workspace query\n\t", query1, "\nwith args\n\t", args)

	workspace := entities.Workspace{
		Name:        *info.Name,
		Description: info.Description,
		Users: []dto.UserPublicInfo{
			{
				ID:          user.ID,
				Email:       user.Email,
				Name:        user.Name,
				Surname:     user.Surname,
				Description: user.Description,
				AvatarURL:   user.AvatarURL,
			},
		},
	}

	row := tx.QueryRow(query1, args...)
	log.Println("Storage -- DB queried")
	if err := row.Scan(&workspace.ID, &workspace.DateCreated); err != nil {
		log.Println("Storage -- Workspace insert failed with error", err.Error())
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	log.Println("Storage -- Workspace created")

	query2, args, err := sq.
		Insert("user_workspace").
		Columns("id_workspace", "id_user").
		Values(workspace.ID, info.OwnerID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	log.Println("Built workspace user query\n\t", query2, "\nwith args\n\t", args)

	_, err = tx.Exec(query2, args...)

	if err != nil {
		log.Println("Storage -- Workspace user insert failed with error", err.Error())
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrWorkspaceNotCreated
	}

	log.Println("Storage -- Committing changes")
	err = tx.Commit()
	if err != nil {
		log.Println("Storage -- Failed to commit changes")
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
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

	_, err = s.db.Exec(sql, args...)

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

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		log.Println("Failed to delete workspace with error", err.Error())
		return apperrors.ErrWorkspaceNotDeleted
	}

	return nil
}
