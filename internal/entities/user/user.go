package user

import (
	"fmt"
)

type UserId uint

type User struct {
	Id       UserId `db:"id"        json:"id"`
	Name     string `db:"name"      json:"name"`
	Password string	`db:"password"  json:"password"`
}

func NewUser(id UserId, name, password string) User {
	u := User{
		Id: id,
		Name: name,
		Password: password,
	}

	return u
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetPassword(pwd string) {
	u.Password = pwd
}

func (u User) String() string {
	return fmt.Sprintf("%d: %s / %s", u.Id, u.Name, u.Password)
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetId() UserId {
	return u.Id
}
