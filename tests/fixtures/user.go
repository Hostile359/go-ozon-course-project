// +build integration

package fixtures

import "gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"

type UserBuilder struct {
	instance *user.User
}

func User() *UserBuilder {
	return &UserBuilder{
		instance: &user.User{},
	}
}

func (b *UserBuilder) Id(v user.UserId) *UserBuilder {
	b.instance.Id = v
	return b
}

func (b *UserBuilder) Name(v string) *UserBuilder {
	b.instance.Name = v
	return b
}

func (b *UserBuilder) Password(v string) *UserBuilder {
	b.instance.Password = v
	return b
}

func (b *UserBuilder) P() *user.User {
	return b.instance
}

func (b *UserBuilder) V() user.User {
	return *b.instance
}