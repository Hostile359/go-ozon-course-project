package user

import (
	"fmt"
)

type UserId uint

type User struct {
	id       UserId
	name     string
	password string
}

func NewUser(id UserId, name, password string) User {
	u := User{
		id: id,
		name: name,
		password: password,
	}

	return u
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetPassword(pwd string) {
	u.password = pwd
}

func (u User) String() string {
	return fmt.Sprintf("%d: %s / %s", u.id, u.name, u.password)
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetPassword() string {
	return u.password
}

func (u User) GetId() UserId {
	return u.id
}
