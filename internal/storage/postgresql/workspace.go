package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// PostgresWorkspaceStorage
// Хранилище данных в PostgreSQL
type PostgresWorkspaceStorage struct {
	db *sql.DB
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
	funcName := "PostgresWorkspaceStorage.GetUserOwnedWorkspaces"
	errorMessage := "Creating workspace failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetUserOwnedWorkspaces FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetUserOwnedWorkspaces <<<<<<<<<<<<<<<<<<<")

	workspaceQuery, args, err := sq.
		Select("id", "name").
		From("public.workspace").
		Where(sq.Eq{"id_creator": userID.Value}).
		OrderBy("public.workspace.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built owned workspace query\n\t"+workspaceQuery+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err := s.db.Query(workspaceQuery, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotGetWorkspace
	}
	defer rows.Close()
	logger.DebugFmt("Workspaces got", requestID.String(), funcName, nodeName)

	workspaces := []dto.UserOwnedWorkspaceInfo{}
	var ownedIDs []uint64
	for rows.Next() {
		var workspace dto.UserOwnedWorkspaceInfo

		err = rows.Scan(
			&workspace.ID,
			&workspace.Name,
		)
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotGetWorkspace
		}
		workspaces = append(workspaces, workspace)
		ownedIDs = append(ownedIDs, workspace.ID)
	}

	boardQuery, args, err := sq.
		Select("id_workspace", "id", "name", "description", "thumbnail_url").
		From("public.board").
		Where(sq.Eq{"id_workspace": ownedIDs}).
		OrderBy("public.board.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+boardQuery+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err = s.db.Query(boardQuery, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer rows.Close()
	logger.DebugFmt("Boards got", requestID.String(), funcName, nodeName)

	for rows.Next() {
		var board dto.WorkspaceBoardInfo
		var workspaceID uint64

		err = rows.Scan(
			&workspaceID,
			&board.ID,
			&board.Name,
			&board.Description,
			&board.ThumbnailURL,
		)
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotGetBoard
		}

		for index, workspace := range workspaces {
			if workspace.ID == workspaceID {
				workspaces[index].Boards = append(workspaces[index].Boards, board)
				break
			}
		}
	}
	logger.DebugFmt("Boards collected", requestID.String(), funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetUserOwnedWorkspaces SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &workspaces, nil
}

// GetUserGuestWorkspaces
// находит рабочие пространства, связанные с пользователем в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetUserGuestWorkspaces(ctx context.Context, userID dto.UserID) (*[]dto.UserGuestWorkspaceInfo, error) {
	funcName := "PostgresWorkspaceStorage.GetUserGuestWorkspaces"
	errorMessage := "Creating workspace failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetUserGuestWorkspaces FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetUserGuestWorkspaces <<<<<<<<<<<<<<<<<<<")

	workspaceQuery, args, err := sq.
		Select(userGuestWorkspaceFields...).
		From("public.workspace").
		LeftJoin("public.user_workspace ON public.user_workspace.id_workspace = public.workspace.id").
		LeftJoin("public.user ON public.user.id = public.user_workspace.id_user").
		Where(sq.And{
			sq.NotEq{"public.workspace.id_creator": userID.Value},
			sq.Eq{"public.user_workspace.id_user": userID.Value},
		}).
		OrderBy("public.workspace.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built guest workspace query\n\t"+workspaceQuery+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err := s.db.Query(workspaceQuery, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotGetWorkspace
	}
	defer rows.Close()
	logger.DebugFmt("Guest workspaces got", requestID.String(), funcName, nodeName)

	workspaceRows := []dto.UserGuestWorkspaceInfo{}
	guestWorkspaceIDs := []uint64{}
	for rows.Next() {
		var workspace dto.UserGuestWorkspaceInfo

		err = rows.Scan(
			&workspace.ID,
			&workspace.Name,
			&workspace.DateCreated,
			&workspace.Owner.ID,
			&workspace.Owner.Email,
			&workspace.Owner.Name,
			&workspace.Owner.Surname,
		)
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotGetWorkspace
		}

		workspaceRows = append(workspaceRows, workspace)
		guestWorkspaceIDs = append(guestWorkspaceIDs, workspace.ID)
	}

	boardQuery, args, err := sq.
		Select(guestBoardFields...).
		From("public.board").
		LeftJoin("public.board_user ON public.board_user.id_board = public.board.id").
		Where(sq.And{
			sq.Eq{"public.board.id_workspace": guestWorkspaceIDs},
			sq.Eq{"public.board_user.id_user": userID.Value},
		}).
		OrderBy("public.board.date_created").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built boards query\n\t"+boardQuery+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err = s.db.Query(boardQuery, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer rows.Close()
	logger.DebugFmt("Boards got", requestID.String(), funcName, nodeName)

	for rows.Next() {
		var board dto.WorkspaceBoardInfo
		var workspaceID uint64

		err = rows.Scan(
			&workspaceID,
			&board.ID,
			&board.Name,
			&board.Description,
			&board.ThumbnailURL,
		)
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotGetBoard
		}

		for index, workspace := range workspaceRows {
			if workspace.ID == workspaceID {
				workspaceRows[index].Boards = append(workspaceRows[index].Boards, board)
				break
			}
		}
	}
	logger.DebugFmt("Boards collected", requestID.String(), funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetUserGuestWorkspaces SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &workspaceRows, nil
}

// Create
// создает новоt рабочее пространство в БД по данным
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Create(ctx context.Context, info dto.NewWorkspaceInfo) (*entities.Workspace, error) {
	funcName := "PostgresWorkspaceStorage.Create"
	errorMessage := "Creating workspace failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Create FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Create <<<<<<<<<<<<<<<<<<<")

	query1, args, err := sq.
		Insert("public.workspace").
		Columns("name", "description", "id_creator").
		Values(info.Name, info.Description, info.OwnerID).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query1+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	user := ctx.Value(dto.UserObjKey).(*entities.User)
	workspace := entities.Workspace{
		Name:        info.Name,
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
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotStartTransaction
	}
	logger.DebugFmt("Began transaction", requestID.String(), funcName, nodeName)

	row := tx.QueryRow(query1, args...)
	if err := row.Scan(&workspace.ID, &workspace.DateCreated); err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotRollback
		}
		logger.Debug(failBorder)
		return nil, apperrors.ErrWorkspaceNotCreated
	}
	logger.DebugFmt("Workspace created", requestID.String(), funcName, nodeName)

	query2, args, err := sq.
		Insert("public.user_workspace").
		Columns("id_workspace", "id_user").
		Values(workspace.ID, info.OwnerID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotRollback
		}
		logger.Debug(failBorder)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query2+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = tx.Exec(query2, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotRollback
		}
		logger.Debug(failBorder)
		return nil, apperrors.ErrWorkspaceNotCreated
	}
	logger.DebugFmt("Workspace - user connection created", requestID.String(), funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, apperrors.ErrCouldNotRollback
		}
		logger.Debug(failBorder)
		return nil, apperrors.ErrWorkspaceNotCreated
	}
	logger.DebugFmt("Transaction commited", requestID.String(), funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Create SUCCESS <<<<<<<<<<<<<<<<<<<")

	return &workspace, nil
}

// UpdateData
// обновляет рабочее пространство в БД
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) UpdateData(ctx context.Context, info dto.UpdatedWorkspaceInfo) error {
	funcName := "PostgresWorkspaceStorage.UpdateData"
	errorMessage := "Updating workspace data failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Create FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.UpdateData <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Update("public.workspace").
		Set("name", info.Name).
		Set("description", info.Description).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrUserNotUpdated
	}
	logger.DebugFmt("Workspace updated", requestID.String(), funcName, nodeName)

	return nil
}

// Delete
// удаляет данногt рабочее пространство в БД по id
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) Delete(ctx context.Context, id dto.WorkspaceID) error {
	funcName := "PostgresWorkspaceStorage.Delete"
	errorMessage := "Deleting workspace failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Delete FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Delete <<<<<<<<<<<<<<<<<<<")

	query, args, err := sq.
		Delete("public.workspace").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return apperrors.ErrWorkspaceNotDeleted
	}
	logger.DebugFmt("Workspace deleted", requestID.String(), funcName, nodeName)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.Delete SUCCESS <<<<<<<<<<<<<<<<<<<")

	return nil
}

/*

// GetWorkspace
// находит рабочее пространство и связанные доски в БД по его id
// или возвращает ошибки ...
func (s PostgresWorkspaceStorage) GetWorkspace(ctx context.Context, id dto.WorkspaceID) (*entities.Workspace, error) {
	funcName := "PostgresWorkspaceStorage.GetWorkspace"
	errorMessage := "Creating workspace failed with error: "
	failBorder := ">>>>>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetWorkspace FAIL <<<<<<<<<<<<<<<<<<<<<<<"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	logger.Debug(">>>>>>>>>>>>>>>> PostgresWorkspaceStorage.GetWorkspace <<<<<<<<<<<<<<<<<<<")

	sql, args, err := sq.
		Select(allWorkspaceAndBoardFields...).
		From("public.workspace").
		Join("public.board ON public.workspace.id = public.board.workspace_id").
		Where(sq.Eq{"public.workspace.id": id.Value}).
		ToSql()
	if err != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
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
			logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
			logger.Debug(failBorder)
			return nil, err
		}

		workspace.Boards = append(workspace.Boards, board)
	}
	if rows.Err() != nil {
		logger.DebugFmt(errorMessage+err.Error(), requestID.String(), funcName, nodeName)
		logger.Debug(failBorder)
		return nil, err
	}

	return &workspace, nil
}

*/
