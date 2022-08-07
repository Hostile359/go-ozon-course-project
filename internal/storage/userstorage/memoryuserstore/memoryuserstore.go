package memoryuserstore

import (
	"context"
	"fmt"
	"sync"
	
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

const (
	poolSize = 10
)

var _ userapp.Storage = &Storage{}

type Storage struct {
	data   map[user.UserId]user.User
	mu     sync.RWMutex
	poolCh chan struct{}
	lastId user.UserId
}

func New() userapp.Storage {
	return &Storage{
		data:   make(map[user.UserId]user.User),
		mu:     sync.RWMutex{},
		poolCh: make(chan struct{}, poolSize),
		lastId: 1,
	}
}

func (s *Storage) Add(ctx context.Context, u user.User) error {
	select {
	case s.poolCh <- struct{}{}:
		defer func() { <-s.poolCh }()
	case <-ctx.Done():
		return ctx.Err()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.lastId
	u.Id = s.lastId
	if _, ok := s.data[id]; ok {
		return errors.Wrapf(userapp.ErrUserExists, "user-id: [%d]", id)
	}
	s.data[id] = u
	s.lastId += 1

	return nil
}

func (s *Storage) Update(ctx context.Context, u user.User) error {
	select {
	case s.poolCh <- struct{}{}:
		defer func() { <-s.poolCh }()
	case <-ctx.Done():
		return ctx.Err()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	id := u.GetId()
	if _, ok := s.data[id]; !ok {
		return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}
	s.data[id] = u

	return nil
}

func (s *Storage) Get(ctx context.Context, id user.UserId) (*user.User, error) {
	select {
	case s.poolCh <- struct{}{}:
		defer func() { <-s.poolCh }()
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.data[id]
	if !ok {
		return nil, errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}

	return &u, nil
}

func (s *Storage) List(ctx context.Context, offset, limit uint64) ([]user.User, error) {
	select {
	case s.poolCh <- struct{}{}:
		defer func() { <-s.poolCh }()
	case <-ctx.Done():
		fmt.Println("Tuta timeput")
		return nil, ctx.Err()
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	res := make([]user.User, 0, len(s.data))
	for _, v := range s.data {
		res = append(res, v)
	}

	return res, nil
}

func (s *Storage) Delete(ctx context.Context, id user.UserId) error {
	select {
	case s.poolCh <- struct{}{}:
		defer func() { <-s.poolCh }()
	case <-ctx.Done():
		return ctx.Err()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[id]; ok {
		delete(s.data, id)
		return nil
	}
	return errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
}
