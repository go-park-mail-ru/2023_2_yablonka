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
	"github.com/sirupsen/logrus"
)

// PostgreSQLBoardStorage
// Хранилище досок в PostgreSQL
type PostgreSQLBoardStorage struct {
	db *pgxpool.Pool
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

func NewBoardStorage(db *pgxpool.Pool) *PostgreSQLBoardStorage {
	return &PostgreSQLBoardStorage{
		db: db,
	}
}

// TODO Ограничить количество всего за раз
// GetById
// находит доску и связанные с ней списки и задания по id
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetById(ctx context.Context, id dto.BoardID) (*dto.FullBoardResult, error) {
	funcName := "GetById"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	// user, ok := ctx.Value(dto.UserObjKey).(*entities.User)
	// if !ok {
	// 	log.Println("Storage -- Failed to get user")
	// 	return nil, apperrors.ErrCouldNotBuildQuery
	// }

	boardSql, args, err := sq.Select(allBoardListAggFields...).
		From("public.board").
		Where(sq.Eq{"public.board.id": id.Value}).
		LeftJoin("public.list ON public.list.id_board = public.board.id").
		LeftJoin("public.workspace ON public.board.id_workspace = public.workspace.id").
		GroupBy(allBoardListFields...).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	storageDebugLog(logger, funcName, "Built query\n\t"+boardSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args))

	rows, err := s.db.Query(context.Background(), boardSql, args...)
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to get board with error "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer rows.Close()
	storageDebugLog(logger, funcName, "Got board info rows")

	boardRows, err := pgx.CollectRows(rows, pgx.RowToStructByPos[BoardListReturn])
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to collect rows with error "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	storageDebugLog(logger, funcName, "Collected "+fmt.Sprintf("%d", len(boardRows))+" board info rows")
	storageDebugLog(logger, funcName, fmt.Sprintf("%+v", boardRows))

	result := dto.FullBoardResult{
		Board: dto.SingleBoardInfo{
			ID:           boardRows[0].ID,
			Name:         boardRows[0].Name,
			WorkspaceID:  boardRows[0].WorkspaceID,
			OwnerID:      boardRows[0].WorkspaceOwnerID,
			ThumbnailURL: boardRows[0].ThumbnailURL,
			DateCreated:  boardRows[0].DateCreated,
			Lists:        boardRows[0].Lists,
		},
	}
	listIDs := result.Board.Lists

	listSql, args, err := sq.Select(allListTaskAggFields...).
		From("public.list").
		Where(sq.Eq{"public.list.id": listIDs}).
		LeftJoin("public.task ON public.task.id_list = public.list.id").
		GroupBy(allListTaskFields...).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	storageDebugLog(logger, funcName, "Built query\n\t"+listSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args))

	rows, err = s.db.Query(context.Background(), listSql, args...)
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to get lists with error "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer rows.Close()
	storageDebugLog(logger, funcName, "Got list info rows")

	lists, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.SingleListInfo])
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to collect rows with error "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	storageDebugLog(logger, funcName, "Collected list info rows")

	result.Lists = lists
	taskIDs := []uint64{}

	for _, list := range lists {
		taskIDs = append(taskIDs, list.Tasks...)
	}

	taskSql, args, err := sq.Select(allTaskAggFields...).
		From("public.task").
		Where(sq.Eq{"public.task.id": taskIDs}).
		LeftJoin("public.task_user ON public.task_user.id_task = public.task.id").
		LeftJoin("public.user ON public.task_user.id_user = public.user.id").
		GroupBy("public.task.id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	storageDebugLog(logger, funcName, "Built query\n\t"+taskSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args))

	rows, err = s.db.Query(context.Background(), taskSql, args...)
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to get tasks with error "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer rows.Close()
	storageDebugLog(logger, funcName, "Got task info rows")

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.SingleTaskInfo])
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to collect rows with error "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	storageDebugLog(logger, funcName, "Collected task info rows")

	result.Tasks = tasks

	return &result, nil
}

// GetUsers
// находит пользователей, у которых есть доступ к доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetUsers(ctx context.Context, id dto.BoardID) (*[]dto.UserPublicInfo, error) {
	funcName := "GetUsers"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	sql, args, err := sq.Select(allPublicUserFields...).
		From("public.user").
		Join("public.board_user ON public.board_user.id_user = public.user.id").
		Where(sq.Eq{"public.board_user.id_board": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	storageDebugLog(logger, funcName, "Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args))

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to get board users with error "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoardUsers
	}
	defer rows.Close()
	storageDebugLog(logger, funcName, "Got board users")

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.UserPublicInfo])
	if err != nil {
		storageDebugLog(logger, funcName, "Failed to parse results into struct with "+err.Error())
		return nil, apperrors.ErrCouldNotGetBoard
	}
	storageDebugLog(logger, funcName, "Parsed results")

	return &users, nil
}

func (s *PostgreSQLBoardStorage) Create(ctx context.Context, info dto.NewBoardInfo) (*entities.Board, error) {
	user, ok := ctx.Value(dto.UserObjKey).(*entities.User)
	if !ok {
		log.Println("Storage -- Failed to get user")
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	baseURL, ok := ctx.Value(dto.BaseURLKey).(string)
	if !ok {
		log.Println("Storage -- Failed to get base url")
		return nil, apperrors.ErrCouldNotBuildQuery
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
		log.Println("Storage -- Failed to build query1")
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built board query\n\t", query1, "\nwith args\n\t", args)

	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, apperrors.ErrCouldNotStartTransaction
	}
	log.Println("Storage -- Transaction started")

	var boardID int
	row := tx.QueryRow(ctx, query1, args...)
	if err := row.Scan(&boardID, &newBoard.DateCreated); err != nil {
		log.Println("Storage -- Board insert failed with error", err.Error())
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	log.Println("Storage -- Board created")

	var url string
	if info.ThumbnailURL != nil {
		url = *info.ThumbnailURL
	} else {
		url = baseURL + "img/board_thumbnails/" + strconv.Itoa(boardID) + ".png"
	}

	newBoard.ThumbnailURL = &url
	newBoard.ID = uint64(boardID)

	query2, args, err := sq.
		Update("public.board").
		Set("thumbnail_url", url).
		Where(sq.Eq{"id_workspace": info.WorkspaceID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build thumbnail update query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built board query\n\t", query2, "\nwith args\n\t", args)

	_, err = tx.Exec(ctx, query2, args...)
	if err != nil {
		log.Println("Storage -- Board thumbnail insert failed with error", err.Error())
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	log.Println("Storage -- Board thumbnail set")

	query3, args, err := sq.
		Insert("public.board_user").
		Columns("id_board", "id_user").
		Values(boardID, info.OwnerID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built board user query\n\t", query3, "\nwith args\n\t", args)

	_, err = tx.Exec(ctx, query3, args...)
	if err != nil {
		log.Println("Storage -- Board user insert failed with error", err.Error())
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	log.Println("Storage -- Creator connection set")

	err = tx.Commit(ctx)
	if err != nil {
		log.Println("Storage -- Failed to commit changes")
		err = tx.Rollback(ctx)
		for err != nil {
			err = tx.Rollback(ctx)
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	log.Println("Storage -- Committed changes")

	return newBoard, nil
}

func (s *PostgreSQLBoardStorage) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	sql, args, err := sq.
		Update("public.board").
		Set("name", info.Name).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		log.Println("Storage -- Failed to build query")
		return apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built board query\n\t", sql, "\nwith args\n\t", args)

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotUpdated
	}

	return nil
}

func (s *PostgreSQLBoardStorage) UpdateThumbnailUrl(ctx context.Context, info dto.ImageUrlInfo) error {
	sql, args, err := sq.
		Update("public.board").
		Set("thumbnail_url", info.Url).
		Where(sq.Eq{"id": info.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotUpdated
	}

	return nil
}

func (s *PostgreSQLBoardStorage) Delete(ctx context.Context, id dto.BoardID) error {
	sql, args, err := sq.
		Delete("public.board").
		Where(sq.Eq{"id": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}

	_, err = s.db.Exec(ctx, sql, args...)

	if err != nil {
		return apperrors.ErrBoardNotDeleted
	}

	return nil
}

// AddUser
// добавляет пользователя на доску
func (s *PostgreSQLBoardStorage) AddUser(ctx context.Context, info dto.AddBoardUserInfo) error {
	funcName := "AddUser"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	storageDebugLog(logger, funcName, "Adding user to board")
	query, args, err := sq.
		Insert("public.board_user").
		Columns("id_board", "id_user").
		Values(info.BoardID, info.UserID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	storageDebugLog(logger, funcName, "Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args))

	_, err = s.db.Exec(ctx, query, args...)

	if err != nil {
		storageDebugLog(logger, funcName, "Failed to execute query with error "+err.Error())
		return apperrors.ErrCouldNotAddBoardUser
	}

	return nil
}

// AddUser
// добавляет пользователя на доску
func (s *PostgreSQLBoardStorage) RemoveUser(ctx context.Context, info dto.RemoveBoardUserInfo) error {
	funcName := "RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(*logrus.Logger)

	storageDebugLog(logger, funcName, "Removing user from board")
	query, args, err := sq.
		Delete("public.board_user").
		Where(sq.Eq{"id_board": info.BoardID, "id_user": info.UserID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	storageDebugLog(logger, funcName, "Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args))

	_, err = s.db.Exec(ctx, query, args...)

	if err != nil {
		storageDebugLog(logger, funcName, "Failed to execute query with error "+err.Error())
		return apperrors.ErrCouldNotAddBoardUser
	}

	return nil
}
