package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/internal/apperrors"
	logger "server/internal/logging"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

const nodeName string = "storage"

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

// TODO Ограничить количество всего за раз
// GetById
// находит доску и связанные с ней списки и задания по id
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetById(ctx context.Context, id dto.BoardID) (*dto.SingleBoardInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetById"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	boardSql, args, err := sq.Select(allBoardFields...).
		From("public.board").
		Where(sq.Eq{"public.board.id": id.Value}).
		LeftJoin("public.workspace ON public.board.id_workspace = public.workspace.id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+boardSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

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
	logger.Debug(fmt.Sprintf("%+v", board), funcName, nodeName)

	return &board, nil
}

// GetUsers
// находит пользователей, у которых есть доступ к доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetUsers(ctx context.Context, id dto.BoardID) (*[]dto.UserPublicInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetUsers"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	sql, args, err := sq.Select(allPublicUserFields...).
		From("public.user").
		Join("public.board_user ON public.board_user.id_user = public.user.id").
		Where(sq.Eq{"public.board_user.id_board": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+sql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetBoardUsers
	}
	defer rows.Close()
	logger.Debug("Got board users", funcName, nodeName)

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
	logger.Debug("Parsed results", funcName, nodeName)

	return &users, nil
}

// GetLists
// находит списки в доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetLists(ctx context.Context, id dto.BoardID) (*[]dto.SingleListInfo, error) {
	funcName := "PostgreSQLBoardStorage.GetLists"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

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
	logger.Debug("Built query\n\t"+listSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	rows, err := s.db.Query(listSql, args...)
	if err != nil {
		logger.Debug(err.Error(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotGetList
	}
	defer rows.Close()
	logger.Debug("Got list rows", funcName, nodeName)

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
			logger.Debug(err.Error(), funcName, nodeName)
			return nil, apperrors.ErrCouldNotGetBoard
		}
		lists = append(lists, list)
	}
	logger.Debug("Collected list info rows", funcName, nodeName)

	return &lists, nil
}

// CheckAccess
// находит пользователя в задании
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) CheckAccess(ctx context.Context, info dto.CheckBoardAccessInfo) (bool, error) {
	funcName := "PostgreSQLBoardStorage.CheckAccess"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

	listSql, args, err := sq.Select("count(*)").
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
	logger.Debug("Built query\n\t"+listSql+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	row := s.db.QueryRow(listSql, args...)
	logger.Debug("Got user row", funcName, nodeName)

	var count uint64
	if row.Scan(&count) != nil {
		return false, apperrors.ErrCouldNotGetUser
	}
	logger.Debug("checked database", funcName, nodeName)

	return count > 0, nil
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

	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, apperrors.ErrCouldNotStartTransaction
	}
	log.Println("Storage -- Transaction started")

	var boardID int
	row := tx.QueryRow(query1, args...)
	if err := row.Scan(&boardID, &newBoard.DateCreated); err != nil {
		log.Println("Storage -- Board insert failed with error", err.Error())
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
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

	_, err = tx.Exec(query2, args...)
	if err != nil {
		log.Println("Storage -- Board thumbnail insert failed with error", err.Error())
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
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

	_, err = tx.Exec(query3, args...)
	if err != nil {
		log.Println("Storage -- Board user insert failed with error", err.Error())
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return nil, apperrors.ErrBoardNotCreated
	}
	log.Println("Storage -- Creator connection set")

	err = tx.Commit()
	if err != nil {
		log.Println("Storage -- Failed to commit changes")
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
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

	_, err = s.db.Exec(sql, args...)

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

	_, err = s.db.Exec(sql, args...)

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
	logger.Debug("Built query\n\t"+query1+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(query1, args...)
	if err != nil {
		logger.Debug("Insert into board_user failed with error "+err.Error(), funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return apperrors.ErrCouldNotAddBoardUser
	}
	logger.Debug("query executed", funcName, nodeName)

	query2, args, err := sq.
		Insert("public.user_workspace").
		Columns("id_workspace", "id_user").
		Values(info.WorkspaceID, info.UserID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+query2+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = tx.Exec(query2, args...)
	if err != nil {
		logger.Debug("Insert into user_workspace failed with error "+err.Error(), funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return apperrors.ErrCouldNotAddBoardUser
	}
	logger.Debug("query executed", funcName, nodeName)

	err = tx.Commit()
	if err != nil {
		logger.Debug("Failed to commit changes with error "+err.Error(), funcName, nodeName)
		err = tx.Rollback()
		for err != nil {
			err = tx.Rollback()
		}
		return apperrors.ErrCouldNotAddBoardUser
	}
	logger.Debug("Changes committed", funcName, nodeName)

	return nil
}

// AddUser
// добавляет пользователя на доску
func (s *PostgreSQLBoardStorage) RemoveUser(ctx context.Context, info dto.RemoveBoardUserInfo) error {
	funcName := "PostgreSQLBoardStorage.RemoveUser"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)

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
	logger.Debug("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		logger.Debug("Delete failed with error "+err.Error(), funcName, nodeName)
		return apperrors.ErrCouldNotRemoveTaskUser
	}
	logger.Debug("query executed", funcName, nodeName)

	return nil
}
