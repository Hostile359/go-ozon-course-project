package userstore

type UserStore interface{
	Init()
	AddUser(string, string) error
	UpdateUser(uint, string, string) error
	GetUser(uint) (*User, error)
	GetAllUsers() []*User
	DeleteUser(uint) error
}
