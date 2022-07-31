package memoryuserstore

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

const (
	poolSize = 10
	storageTimeout = 10*time.Second
)

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
	ch := make(chan error, 1)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer func() {
		s.mu.Unlock()
		<- s.poolCh
		cancel()
	}()

	go func() {
		id := u.GetId()
		if _, ok := s.data[id]; ok {
			ch <- errors.Wrapf(userapp.ErrUserExists, "user-id: [%d]", id)
		}
		s.data[id] = u
		ch <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
 	}
}

func (s *Storage) Update(u user.User) error {
	s.poolCh <- struct{}{}
	s.mu.Lock()
	ch := make(chan error, 1)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer func() {
		s.mu.Unlock()
		<- s.poolCh
		cancel()
	}()
	
	go func() {
		id := u.GetId()
		if _, ok := s.data[id]; !ok {
			ch <- errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
		}
		s.data[id] = u
		ch <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
 	}
}

func (s *Storage) Get(id user.UserId) (*user.User, error){
	s.poolCh <- struct{}{}
	s.mu.RLock()
	type ChMessage struct {
		u *user.User
		err error
	}
	ch := make(chan ChMessage, 1)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer func() {
		s.mu.RUnlock()
		<- s.poolCh
		cancel()
	}()

	go func() {
		u, ok := s.data[id]
		if !ok {
			ch <- ChMessage{
				u: nil,
				err: errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id),
			}
		}

		ch <- ChMessage{
			u: &u,
			err: nil,
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-ch:
		return res.u, res.err
 	}
}

func (s *Storage) List() ([]user.User, error) {
	s.poolCh <- struct{}{}
	s.mu.RLock()
	ch := make(chan []user.User, 1)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer func() {
		s.mu.RUnlock()
		<- s.poolCh
		cancel()
	}()
	
	go func() {
		res := make([]user.User , 0, len(s.data))
		for _, v := range s.data {
			res = append(res, v)
		}
		ch <- res
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-ch:
		return res, nil
 	}
}

func (s *Storage) Delete(id user.UserId) error {
	s.poolCh <- struct{}{}
	s.mu.Lock()
	ch := make(chan error, 1)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, storageTimeout)
	defer func() {
		s.mu.Unlock()
		<- s.poolCh
		cancel()
	}()
	
	go func() {
		if _, ok := s.data[id]; ok {
			delete(s.data, id)
			ch <- nil
		}
		ch <- errors.Wrapf(userapp.ErrUserNotExists, "user-id: [%d]", id)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
 	}
}
