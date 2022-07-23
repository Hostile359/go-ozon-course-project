package memoryuserstore

import (
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

var _ userapp.Storage = &Storage{}

type Storage struct {
	data map[user.UserId]user.User
}

func New() userapp.Storage {
	return &Storage{
		data: make(map[user.UserId]user.User),
	}
}

func (s *Storage) Add(u user.User) error {
	id := u.GetId()
	if _, ok := s.data[id]; ok {
		return errors.Wrapf(userapp.ErrUserExists, "user-id: [%d]", id)
	}
	s.data[id] = u

	return nil
}

func (s *Storage) Update(u user.User) error {
	id := u.GetId()
	if _, ok := s.data[id]; !ok {
		return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}
	s.data[id] = u

	return nil
}

func (s *Storage) Get(id user.UserId) (*user.User, error){
	u, ok := s.data[id]
	if !ok {
		return nil, errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}

	return  &u, nil
}

func (s *Storage) List() []user.User {
	res := make([]user.User , 0, len(s.data))
	for _, v := range s.data {
		res = append(res, v)
	}
	return res
}

func (s *Storage) Delete(id user.UserId) error {
	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return nil
	}
	return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
}
