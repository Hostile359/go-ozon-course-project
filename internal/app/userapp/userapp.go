package userapp

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

const (
	storageTimeout = 10*time.Second
	pageDefault = 1
	perPageDefault = 5
)

var (
	ErrValidationArgs = errors.New("Bad argument")
	ErrUserNotExists = errors.New("user does not exist")
	ErrUserExists    = errors.New("user exists")
)

type Storage interface {
	Add(context.Context, user.User) error
	Update(context.Context, user.User) error
	Get(context.Context, user.UserId) (*user.User, error)
	List(context.Context, uint64, uint64) ([]user.User, error)
	Delete(context.Context, user.UserId) error
}

type App struct {
	userStorage Storage
}

func New(userStorage Storage) *App {
	return &App{
		userStorage: userStorage,
	}
}

func (a *App) Add(ctx context.Context, u user.User) error {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.userStorage.Add(ctx, u)
}

func (a *App) Update(ctx context.Context, u user.User) error {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.userStorage.Update(ctx, u)
}

func (a App) Get(ctx context.Context, id user.UserId)  (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.userStorage.Get(ctx, id)
}

func (a App) List(ctx context.Context, page, perPage uint64) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	if page == 0 {
		page = pageDefault
	}
	if perPage == 0 {
		perPage = perPageDefault
	}

	offset := (page - 1) * perPage
	limit := perPage
	return a.userStorage.List(ctx, offset, limit)
}

func (a *App)Delete(ctx context.Context, id user.UserId) error {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()
	
	return a.userStorage.Delete(ctx, id)
}
