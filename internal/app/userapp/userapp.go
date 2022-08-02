package userapp

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

const (
	storageTimeout = 10*time.Second
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
	List(context.Context) ([]user.User, error)
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
	if err := checkName(u.GetName()); err != nil {
		return err
	}
	if err := checkPassword(u.GetPassword()); err != nil {
		return err
	}
	
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.userStorage.Add(ctx, u)
}

func (a *App) Update(ctx context.Context, u user.User) error {
	if err := checkName(u.GetName()); err != nil {
		return err
	}
	if err := checkPassword(u.GetPassword()); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.userStorage.Update(ctx, u)
}

func (a App) Get(ctx context.Context, id user.UserId)  (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.userStorage.Get(ctx, id)
}

func (a App) List(ctx context.Context) ([]user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()

	return a.userStorage.List(ctx)
}

func (a *App)Delete(ctx context.Context, id user.UserId) error {
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer cancel()
	
	return a.userStorage.Delete(ctx, id)
}

func checkName(name string) error {
	if len(name) == 0 || len(name) > 10 {
		return errors.Wrapf(ErrValidationArgs, "<%v>, len should be from 1 to 10", name)
	}
	return nil
}

func checkPassword(password string) error {
	if len(password) < 6 || len(password) > 10 {
		return errors.Wrapf(ErrValidationArgs, "<%v>, len should be from 6 to 10", password)
	}	
	return nil
}
