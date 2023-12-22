package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// PostgreSQLBoardStorage
// Хранилище досок в PostgreSQL
type PostgreSQLBoardStorage struct {
	db *sql.DB
}

type BoardListReturn struct {
	ID               uint64
	WorkspaceID      uint64
	WorkspaceOwnerID uint64
	Name             string
	DateCreated      time.Time
	ThumbnailURL     *string
	Lists            []uint64
}

type ListTaskReturn struct {
	ID           uint64
	BoardID      uint64
	Name         string
	ListPosition uint64
	Tasks        []uint64
}

func NewBoardStorage(db *sql.DB) *PostgreSQLBoardStorage {
	return &PostgreSQLBoardStorage{
		db: db,
	}
}

// GetById
// находит доску и связанные с ней списки и задания по id
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetById(ctx context.Context, id dto.BoardID) (*dto.SingleBoardInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetById"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	boardSql, args, err := sq.Select(allBoardFields...).
		From("public.board").
		Where(sq.Eq{"public.board.id": id.Value}).
		LeftJoin("public.workspace ON public.board.id_workspace = public.workspace.id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+boardSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	row := s.db.QueryRow(boardSql, args...)

	var board dto.SingleBoardInfo
	err = row.Scan(
		&board.ID,
		&board.WorkspaceID,
		&board.WorkspaceOwnerID,
		&board.Name,
		&board.DateCreated,
		&board.ThumbnailURL,
	)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetBoard
	}
	logger.DebugFmt(fmt.Sprintf("%+v", board), requestID.String(), funcName, nodeName)

	return &board, nil
}

// GetUsers
// находит пользователей, у которых есть доступ к доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetUsers(ctx context.Context, id dto.BoardID) (*[]dto.UserPublicInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetUsers"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	sql, args, err := sq.Select(allPublicUserFields...).
		From("public.user").
		Join("public.board_user ON public.board_user.id_user = public.user.id").
		Where(sq.Eq{"public.board_user.id_board": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetBoardUsers
	}
	defer rows.Close()
	logger.DebugFmt("Got board users", requestID.String(), funcName, nodeName)

	users := []dto.UserPublicInfo{}
	for rows.Next() {
		var user dto.UserPublicInfo

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Description,
			&user.Surname,
			&user.AvatarURL,
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}
		users = append(users, user)
	}
	logger.DebugFmt("Parsed results", requestID.String(), funcName, nodeName)

	return &users, nil
}

// GetLists
// находит списки в доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetLists(ctx context.Context, id dto.BoardID) (*[]dto.SingleListInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetLists"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	listSql, args, err := sq.Select(allListTaskAggFields...).
		From("public.list").
		LeftJoin("public.task ON public.task.id_list = public.list.id").
		Where(sq.Eq{"public.list.id_board": id.Value}).
		GroupBy("public.list.id").
		OrderBy("public.list.list_position").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+listSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err := s.db.Query(listSql, args...)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotGetList
	}
	defer rows.Close()
	logger.DebugFmt("Got list rows", requestID.String(), funcName, nodeName)

	lists := []dto.SingleListInfo{}
	for rows.Next() {
		var list dto.SingleListInfo

		err = rows.Scan(
			&list.ID,
			&list.BoardID,
			&list.Name,
			&list.ListPosition,
			(*pq.StringArray)(&list.TaskIDs),
		)
		if err != nil {
			logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotGetBoard
		}
		lists = append(lists, list)
	}
	logger.DebugFmt("Collected list info rows", requestID.String(), funcName, nodeName)

	return &lists, nil
}

// GetTags
// находит тэги в доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetTags(ctx context.Context, id dto.BoardID) (*[]dto.TagInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetTags"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	query, args, err := sq.Select(allTagFields...).
		From("public.tag").
		LeftJoin("public.tag_board ON public.tag.id = public.tag_board.id_tag").
		Where(sq.Eq{"public.tag_board.id_board": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotExecuteQuery
	}
	logger.DebugFmt("Created query", requestID.String(), funcName, nodeName)

	tags := []dto.TagInfo{}
	for rows.Next() {
		var tag dto.TagInfo

		err = rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Color,
		)
		if err != nil {
			logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotScanRows
		}

		tags = append(tags, tag)
	}

	err = rows.Err()
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotCollectRows
	}
	logger.DebugFmt("Collected tags", requestID.String(), funcName, nodeName)

	err = rows.Close()
	if err != nil {
		logger.DebugFmt(err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotCloseQuery
	}
	logger.DebugFmt("Closed rows", requestID.String(), funcName, nodeName)

	return &tags, nil
}

// CheckAccess
// находит пользователя в доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) CheckAccess(ctx context.Context, info dto.CheckBoardAccessInfo) (bool, error) {
	funcName := "PostgreSQLBoardStorage.CheckAccess"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	userSql, args, err := sq.Select("count(*)").
		From("public.board_user").
		Where(sq.Eq{
			"id_board": info.BoardID,
			"id_user":  info.UserID,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+userSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	row := s.db.QueryRow(userSql, args...)
	logger.DebugFmt("Got user row", requestID.String(), funcName, nodeName)

	var count uint64
	if row.Scan(&count) != nil {
		return false, apperrors.ErrCouldNotGetUser
	}
	logger.DebugFmt("checked database", requestID.String(), funcName, nodeName)

	return count > 0, nil
}

// Create
// создает доску
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) Create(ctx context.Context, info dto.NewBoardInfo) (*entities.Board, error) {
	funcName := "PostgreSQLBoardStorage.Create"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	user, ok := ctx.Value(dto.UserObjKey).(*entities.User)
	if !ok || user == nil {
		logger.DebugFmt("No user object in context", requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrNoBoardAccess
	}

	newBoard := &entities.Board{
		Name: info.Name,
		Owner: dto.UserID{
			Value: info.OwnerID,
		},
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

	query1, args, err := sq.
		Insert("public.board").
		Columns("id_workspace", "name").
		Values(info.WorkspaceID, info.Name).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query1+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		logger.DebugFmt("Failed to start transaction with error "+err.Error(), requestID.String(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotStartTransaction
	}
	logger.DebugFmt("Transaction started", requestID.String(), funcName, nodeName)

	var boardID int
	row := tx.QueryRow(query1, args...)
	if err := row.Scan(&boardID, &newBoard.DateCreated); err != nil {
		logger.DebugFmt("Board insert failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	logger.DebugFmt("Board created", requestID.String(), funcName, nodeName)

	var url string
	if info.ThumbnailURL != nil {
		url = *info.ThumbnailURL
	} else {
		url = "img/board_thumbnails/" + strconv.Itoa(boardID) + ".png"
	}

	newBoard.ThumbnailURL = &url
	newBoard.ID = uint64(boardID)

	query2, args, err := sq.
		Update("public.board").
		Set("thumbnail_url", url).
		Where(sq.And{
			sq.Eq{"id_workspace": info.WorkspaceID},
			sq.Eq{"id": boardID},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query2+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = tx.Exec(query2, args...)
	if err != nil {
		logger.DebugFmt("Failed to execute query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	logger.DebugFmt("Board thumbnail URL set", requestID.String(), funcName, nodeName)

	query3, args, err := sq.
		Insert("public.board_user").
		Columns("id_board", "id_user").
		Values(boardID, info.OwnerID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query3+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = tx.Exec(query3, args...)
	if err != nil {
		logger.DebugFmt("Failed to execute query with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	logger.DebugFmt("Board linked to creator", requestID.String(), funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.DebugFmt("Failed to commit changes", requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		if err != nil {
			logger.DebugFmt("Transaction rollback failed with error "+err.Error(), requestID.String(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	logger.DebugFmt("Changes commited", requestID.String(), funcName, nodeName)

	return newBoard, nil
}

// UpdateData
// обновляет данные доски
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	funcName := "PostgreSQLBoardStorage.Create"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	sql, args, err := sq.
		Update("public.board").
		Set("name", info.Name).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		logger.DebugFmt("Failed to build query with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotUpdated
	}

	return nil
}

// UpdateThumbnailUrl
// обновляет картинку доски
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) UpdateThumbnailUrl(ctx context.Context, info dto.BoardImageUrlInfo) error {
	sql, args, err := sq.
		Update("public.board").
		Set("thumbnail_url", info.Url).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotUpdated
	}

	return nil
}

// Delete
// удаляет доску
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) Delete(ctx context.Context, id dto.BoardID) error {
	sql, args, err := sq.
		Delete("public.board").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotDeleted
	}

	return nil
}

// AddUser
// добавляет пользователя на доску
func (s *PostgreSQLBoardStorage) AddUser(ctx context.Context, info dto.AddBoardUserInfo) error {
	funcName := "PostgreSQLBoardStorage.AddUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return apperrors.ErrCouldNotStartTransaction
	}

	query1, args, err := sq.
		Insert("public.board_user").
		Columns("id_board", "id_user").
		Values(info.BoardID, info.UserID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query1+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query1, args...)
	if err != nil {
		logger.DebugFmt("Insert into board_user failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return apperrors.ErrCouldNotAddBoardUser
	}
	logger.DebugFmt("query executed", requestID.String(), funcName, nodeName)

	query2, args, err := sq.
		Insert("public.user_workspace").
		Columns("id_workspace", "id_user").
		Values(info.WorkspaceID, info.UserID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query2+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = tx.Exec(query2, args...)
	if err != nil {
		logger.DebugFmt("Insert into user_workspace failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return apperrors.ErrCouldNotAddBoardUser
	}
	logger.DebugFmt("query executed", requestID.String(), funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.DebugFmt("Failed to commit changes with error "+err.Error(), requestID.String(), funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return apperrors.ErrCouldNotAddBoardUser
	}
	logger.DebugFmt("Changes committed", requestID.String(), funcName, nodeName)

	return nil
}

// AddUser
// добавляет пользователя на доску
func (s *PostgreSQLBoardStorage) RemoveUser(ctx context.Context, info dto.RemoveBoardUserInfo) error {
	funcName := "PostgreSQLBoardStorage.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	requestID := ctx.Value(dto.RequestIDKey).(uuid.UUID)

	query, args, err := sq.
		Delete("public.board_user").
		Where(sq.And{
			sq.Eq{"id_user": info.UserID},
			sq.Eq{"id_board": info.BoardID},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.DebugFmt("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), requestID.String(), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.DebugFmt("Delete failed with error "+err.Error(), requestID.String(), funcName, nodeName)
		return apperrors.ErrCouldNotRemoveTaskUser
	}
	logger.DebugFmt("query executed", requestID.String(), funcName, nodeName)

	// TODO: Remove user from workspace if that was the last board

	return nil
}
