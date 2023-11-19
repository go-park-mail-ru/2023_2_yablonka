package postgresql

import (
	"context"
	"log"
	"server/internal/apperrors"
	"server/internal/pkg/dto"
	"server/internal/pkg/entities"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgreSQLBoardStorage
// Хранилище досок в PostgreSQL
type PostgreSQLBoardStorage struct {
	db *pgxpool.Pool
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
func (s *PostgreSQLBoardStorage) GetById(ctx context.Context, id dto.BoardID) (*entities.Board, error) {
	sql, args, err := sq.Select(append(allBoardFields, append(allListFields, allTaskFields...)...)...).
		From("public.board").
		Join("public.list ON public.board.id = public.list.board_id").
		Join("public.task ON public.list.id = public.task.list_id").
		Where(sq.Eq{"public.board.id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	defer rows.Close()

	var board entities.Board
	for rows.Next() {
		var list entities.List

		err = rows.Scan(
			&board.ID,
			&board.Name,
			&list.ID,
			&list.Name,
			&list.Description,
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}

		for rows.Next() {
			var task entities.Task

			err = rows.Scan(
				&board.ID,
				&board.Name,
				&list.ID,
				&list.Name,
				&list.Description,
				&task.ID,
				&task.Name,
				&task.Description,
			)
			if err != nil {
				return nil, apperrors.ErrCouldNotGetBoard
			}
			list.Tasks = append(list.Tasks, task)
		}

		board.Lists = append(board.Lists, list)
	}

	return &board, nil
}

// GetUsers
// находит пользователей, у которых есть доступ к доске
// или возвращает ошибки ...
func (s *PostgreSQLBoardStorage) GetUsers(ctx context.Context, id dto.BoardID) (*[]dto.UserPublicInfo, error) {
	sql, args, err := sq.Select("user.id").
		From("public.user").
		Join("public.board_user ON public.board_user.id_user = public.user.id").
		Where(sq.Eq{"public.board_user.id_board": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetBoard
	}
	defer rows.Close()

	var users []dto.UserPublicInfo
	for rows.Next() {
		var user dto.UserPublicInfo

		err = rows.Scan(&user)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}

		users = append(users, user)
	}

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

	query1, args, err := sq.
		Insert("public.board").
		Columns("id_workspace", "name").
		Values(info.WorkspaceID, info.Name).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id").
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
	if err := row.Scan(&boardID); err != nil {
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

	query2, args, err := sq.
		Update("public.board").
		Set("thumbnail_url", url).
		Where(sq.Eq{"board.id_workspace": info.WorkspaceID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Println("Storage -- Failed to build thumbnail update query")
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Built board query\n\t", query1, "\nwith args\n\t", args)

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

	return &entities.Board{
		ID:   uint64(boardID),
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
		ThumbnailURL: &url,
	}, nil
}

func (s *PostgreSQLBoardStorage) UpdateData(ctx context.Context, info dto.UpdatedBoardInfo) error {
	sql, args, err := sq.
		Update("public.board").
		Set("name", info.Name).
		Set("description", info.Description).
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
