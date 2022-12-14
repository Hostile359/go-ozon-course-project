//go:generate mockgen -source ../../../app/userapp/userapp.go -destination=./mocks/pguserstore.go -package=mock_pguserstore
package pguserstore

import (
	"context"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const (
	
)

var _ userapp.Storage = &Storage{}

type Storage struct {
	pool pgxpoolmock.PgxPool
}

func New(pool pgxpoolmock.PgxPool) userapp.Storage {
	return &Storage{
		pool: pool,
	}
}

func (s *Storage) Add(ctx context.Context, u user.User) error {
	query, args, err := squirrel.Insert("users").
		Columns("name, password").
		Values(u.GetName(), u.GetPassword()).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return errors.Errorf("Storage.Add: to sql: %v", err)
	}
	if _, err = s.pool.Exec(ctx, query, args...); err != nil {
		return errors.Errorf("Storage.Add: insert: %v", err)
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, u user.User) error {
	query, args, err := squirrel.Update("users").
		Set("name", u.GetName()).
		Set("password", u.GetPassword()).
		Where(squirrel.Eq{
			"id": u.GetId(),
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
		return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", u.GetId())
	}
	return nil
}

func (s *Storage) Get(ctx context.Context, id user.UserId) (*user.User, error) {
	newCtx, span := otel.Tracer(userapp.TracerName).Start(ctx, "pgUserStore/UserGet")
	defer span.End()

	span.SetAttributes(attribute.String("id", strconv.FormatUint(uint64(id), 10)))

	query, args, err := squirrel.Select("id, name, password").
		From("users").
		Where(squirrel.Eq{
			"id": id,
		}).PlaceholderFormat(squirrel.Dollar).ToSql()
	
	if err != nil {
		return nil, errors.Errorf("Storage.Get: to sql: %v", err)
	}
	var users []user.User
	if err := pgxscan.Select(newCtx, s.pool, &users, query, args...); err != nil {
		return nil, errors.Errorf("Storage.Get: select: %v", err)
	}
	
	if len(users) == 0 {
		return nil, errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}
	return &users[0], nil
}

func (s *Storage) List(ctx context.Context, offset, limit uint64) ([]user.User, error) {
	newCtx, span := otel.Tracer(userapp.TracerName).Start(ctx, "pgUserStore/UserList")
	defer span.End()

	span.SetAttributes(attribute.String("offset", strconv.FormatUint(offset, 10)))
	span.SetAttributes(attribute.String("limit", strconv.FormatUint(limit, 10)))

	query, args, err := squirrel.Select("id, name, password").
		From("users").
		Offset(offset).
		Limit(limit).
		OrderBy("id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	
	if err != nil {
		return nil, errors.Errorf("Storage.List: to sql: %v", err)
	}
	var users []user.User
	if err := pgxscan.Select(newCtx, s.pool, &users, query, args...); err != nil {
		return nil, errors.Errorf("Storage.List: select: %v", err)
	}

	return users, nil
}

func (s *Storage) Delete(ctx context.Context, id user.UserId) error {
	query, args, err := squirrel.Delete("users").
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
		return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}

	return nil
}

