package memoryuserstore

import (
	"sync"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

const poolSize = 10

var _ userapp.Storage = &Storage{}

type Storage struct {
	data map[user.UserId]user.User
	mu sync.RWMutex
	poolCh chan struct{}
}

func New() userapp.Storage {
	return &Storage{
		data: make(map[user.UserId]user.User),
		mu: sync.RWMutex{},
		poolCh: make(chan struct{}, poolSize),
	}
}

func (s *Storage) Add(u user.User) error {
	s.poolCh <- struct{}{}
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		<- s.poolCh
	}()

	id := u.GetId()
	if _, ok := s.data[id]; ok {
		return errors.Wrapf(userapp.ErrUserExists, "user-id: [%d]", id)
	}
	s.data[id] = u

	return nil
}

func (s *Storage) Update(u user.User) error {
	s.poolCh <- struct{}{}
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		<- s.poolCh
	}()

	id := u.GetId()
	if _, ok := s.data[id]; !ok {
		return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}
	s.data[id] = u

	return nil
}

func (s *Storage) Get(id user.UserId) (*user.User, error){
	s.poolCh <- struct{}{}
	s.mu.RLock()
	defer func() {
		s.mu.RUnlock()
		<- s.poolCh
	}()

	u, ok := s.data[id]
	if !ok {
		return nil, errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}

	return  &u, nil
}

func (s *Storage) List() []user.User {
	s.poolCh <- struct{}{}
	s.mu.RLock()
	defer func() {
		s.mu.RUnlock()
		<- s.poolCh
	}()

	res := make([]user.User , 0, len(s.data))
	for _, v := range s.data {
		res = append(res, v)
	}
	return res
}

func (s *Storage) Delete(id user.UserId) error {
	s.poolCh <- struct{}{}
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		<- s.poolCh
	}()

	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return nil
	}
	return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
}
