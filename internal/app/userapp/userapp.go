package userapp

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

var (
	ErrValidationArgs = errors.New("Bad argument")
	ErrUserNotExists = errors.New("user does not exist")
	ErrUserExists    = errors.New("user exists")
)

type Storage interface {
	Add(user.User) error
	Update(user.User) error
	Get(user.UserId) (*user.User, error)
	List() ([]user.User, error)
	Delete(user.UserId) error
}

type App struct {
	userStorage Storage
}

func New(userStorage Storage) *App {
	return &App{
		userStorage: userStorage,
	}
}

func (a *App) Add(u user.User) error {
	if err := checkName(u.GetName()); err != nil {
		return err
	}
	if err := checkPassword(u.GetPassword()); err != nil {
		return err
	}
	
	return a.userStorage.Add(u)
}

func (a *App) Update(u user.User) error {
	if err := checkName(u.GetName()); err != nil {
		return err
	}
	if err := checkPassword(u.GetPassword()); err != nil {
		return err
	}

	return a.userStorage.Update(u)
}

func (a App) Get(id user.UserId)  (*user.User, error) {
	return a.userStorage.Get(id)
}

func (a App) List() ([]user.User, error) {
	return a.userStorage.List()
}

func (a *App)Delete(id user.UserId) error {
	return a.userStorage.Delete(id)
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
