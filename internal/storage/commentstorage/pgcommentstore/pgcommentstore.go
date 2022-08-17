package pgcommentstore

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/commentapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/comment"
)

const (
	
)

var _ commentapp.Storage = &Storage{}

type Storage struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) commentapp.Storage {
	return &Storage{
		pool: pool,
	}
}

func (s *Storage) Add(ctx context.Context, c comment.Comment) error {
	query, args, err := squirrel.Insert("comments").
		Columns("comment, user_id").
		Values(c.GetComment(), c.GetUserId()).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Errorf("Storage.Add: to sql: %v", err)
	}
	if _, err = s.pool.Exec(ctx, query, args...); err != nil {
		return errors.Errorf("Storage.Add: insert: %v", err)
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, c comment.Comment) error {
	query, args, err := squirrel.Update("comments").
		Set("comment", c.GetComment()).
		Where(squirrel.Eq{
			"id": c.GetId(),
		}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Errorf("Storage.Update: to sql: %v", err)
	}
	result, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Errorf("Storage.Update: Update: %v", err)
	}
	if result.RowsAffected() == 0 {
		return errors.Wrapf(commentapp.ErrCommentNotExists, "comment-id: [%d]", c.GetId())
	}
	return nil
}

func (s *Storage) Get(ctx context.Context, id comment.CommentId) (*comment.Comment, error) {
	query, args, err := squirrel.Select("id, comment, user_id").
		From("comments").
		Where(squirrel.Eq{
			"id": id,
		}).PlaceholderFormat(squirrel.Dollar).ToSql()
	
	if err != nil {
		return nil, errors.Errorf("Storage.Get: to sql: %v", err)
	}
	var comments []comment.Comment
	if err := pgxscan.Select(ctx, s.pool, &comments, query, args...); err != nil {
		return nil, errors.Errorf("Storage.Get: select: %v", err)
	}
	
	if len(comments) == 0 {
		return nil, errors.Wrapf(commentapp.ErrCommentNotExists, "comment-id: [%d]", id)
	}
	return &comments[0], nil
}

func (s *Storage) List(ctx context.Context, offset, limit uint64) ([]comment.Comment, error) {
	query, args, err := squirrel.Select("id, comment, user_id").
		From("comments").
		Offset(offset).
		Limit(limit).
		OrderBy("id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	
	if err != nil {
		return nil, errors.Errorf("Storage.List: to sql: %v", err)
	}
	var comments []comment.Comment
	if err := pgxscan.Select(ctx, s.pool, &comments, query, args...); err != nil {
		return nil, errors.Errorf("Storage.List: select: %v", err)
	}

	return comments, nil
}

func (s *Storage) Delete(ctx context.Context, id comment.CommentId) error {
	query, args, err := squirrel.Delete("comments").
		Where(squirrel.Eq{
			"id": id,
		}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Errorf("Storage.Delete: to sql: %v", err)
	}
	result, err := s.pool.Exec(ctx, query, args...)
	if err != nil {
		return errors.Errorf("Storage.Delete: Delete: %v", err)
	}
	if result.RowsAffected() == 0 {
		return errors.Wrapf(commentapp.ErrCommentNotExists, "comment-id: [%d]", id)
	}

	return nil
}

