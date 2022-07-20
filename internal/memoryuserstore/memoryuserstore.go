package memoryuserstore

import (
	"strconv"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/userstore"
)

var UserNotExists = errors.New("user does not exist")

type MemoryUserStore struct {
	data map[uint]*userstore.User
	lastId uint
}

func (memoryUserStore *MemoryUserStore) Init() {
	memoryUserStore.data = make(map[uint]*userstore.User)
	memoryUserStore.lastId = 0
}

func (memoryUserStore *MemoryUserStore) AddUser(name, password string) error {
	u := userstore.NewUser(memoryUserStore.lastId, name, password)
	memoryUserStore.data[memoryUserStore.lastId] = u
	memoryUserStore.lastId += 1

	return nil
}

func (memoryUserStore *MemoryUserStore) UpdateUser(id uint, newName, newPassword string) error {
	if _, ok := memoryUserStore.data[id]; !ok {
		return errors.Wrap(UserNotExists, strconv.FormatUint(uint64(id), 10))
	}

	memoryUserStore.data[id].SetName(newName)
	memoryUserStore.data[id].SetPassword(newPassword)

	return nil
}

func (memoryUserStore *MemoryUserStore) GetUser(id uint) (*userstore.User, error){
	u, ok := memoryUserStore.data[id]
	if !ok {
		return nil, errors.Wrap(UserNotExists, strconv.FormatUint(uint64(id), 10))
	}

	return  u, nil
}

func (memoryUserStore *MemoryUserStore) GetAllUsers() []*userstore.User {
	res := make([]*userstore.User, 0, len(memoryUserStore.data))
	for _, v := range memoryUserStore.data {
		res = append(res, v)
	}
	return res
}

func (memoryUserStore *MemoryUserStore) DeleteUser(id uint) error {
	if _, ok := memoryUserStore.data[id]; ok {
		delete(memoryUserStore.data, id)
		return nil
	}
	return errors.Wrap(UserNotExists, strconv.FormatUint(uint64(id), 10))
}
