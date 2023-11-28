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

	sq "github.com/Masterminds/squirrel"
)

// PostgresCommentStorage
// Хранилище данных в PostgreSQL
type PostgresCommentStorage struct {
	db *sql.DB
}

// NewCommentStorage
// возвращает PostgreSQL хранилище комментариев
func NewCommentStorage(db *sql.DB) *PostgresCommentStorage {
	return &PostgresCommentStorage{
		db: db,
	}
}

// Create
// создает новый комментарий в бд
// или возвращает ошибки ...
func (s PostgresCommentStorage) Create(ctx context.Context, info dto.NewCommentInfo) (*entities.Comment, error) {
	sql, args, err := sq.
		Insert("public.comment").
		Columns(newCommentFields...).
		Values(info.TaskID, info.UserID, info.Text).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, date_created").
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Formed query\n\t", sql, "\nwith args\n\t", args)

	comment := entities.Comment{
		UserID: info.UserID,
		TaskID: info.TaskID,
		Text:   info.Text,
	}

	query := s.db.QueryRow(sql, args...)
	err = query.Scan(&comment.ID, &comment.DateCreated)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetTaskComments
	}
	log.Println("Created comment")

	return &comment, nil
}

// GetFromTask
// возвращает все комментарии у задания в БД
// или возвращает ошибки ...
func (s *PostgresCommentStorage) GetFromTask(ctx context.Context, id dto.TaskID) (*[]dto.CommentInfo, error) {
	sql, args, err := sq.
		Select(allCommentFields...).
		From("public.comment").
		Where(sq.Eq{"id_task": id.Value}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	log.Println("Formed query\n\t", sql, "\nwith args\n\t", args)

	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetTaskComments
	}
	defer rows.Close()

	comments := []dto.CommentInfo{}
	for rows.Next() {
		var comment dto.CommentInfo

		err = rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.Text,
			&comment.DateCreated,
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}
		comments = append(comments, comment)
	}

	if err != nil {
		return nil, apperrors.ErrCouldNotCollectRows
	}
	log.Println("Collected rows")

	return &comments, nil
}

// ReadMany
// находит задание в БД по их id
// или возвращает ошибки ...
func (s *PostgresCommentStorage) ReadMany(ctx context.Context, ids dto.CommentIDs) (*[]dto.CommentInfo, error) {
	funcName := "PostgresCommentStorage.ReadMany"
	logger := ctx.Value(dto.LoggerKey).(logger.ILogger)
	logger.Debug("got logger", funcName, nodeName)

	query, args, err := sq.
		Select(allCommentFields...).
		From("public.comment").
		Where(sq.Eq{"id": ids.Values}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		logger.Debug(err.Error(), funcName, nodeName)
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	logger.Debug("Built query\n\t"+query+"\nwith args\n\t"+fmt.Sprintf("%+v", args), funcName, nodeName)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, apperrors.ErrCouldNotGetBoardUsers
	}
	defer rows.Close()
	logger.Debug("Got comment rows", funcName, nodeName)

	comments := []dto.CommentInfo{}
	for rows.Next() {
		var comment dto.CommentInfo

		err = rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.Text,
			&comment.DateCreated,
		)
		if err != nil {
			return nil, apperrors.ErrCouldNotGetBoard
		}
		comments = append(comments, comment)
	}
	logger.Debug("Got task from DB", funcName, nodeName)

	return &comments, nil
}
