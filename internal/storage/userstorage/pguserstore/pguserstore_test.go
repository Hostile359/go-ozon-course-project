package pguserstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
)

func TestAddUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		u := user.NewUser(1, "user1", "123456")
		queryStore := "INSERT INTO users (name, password) VALUES ($1,$2)"
		args := []interface{}{u.GetName(), u.GetPassword()}
		f.mockPool.EXPECT().Exec(context.Background(), queryStore, args...).Return(nil, nil)

		// act
		err := f.userStorage.Add(context.Background(), u)

		// assert
		require.NoError(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		u := user.NewUser(1, "user1", "123456")
		queryStore := "UPDATE users SET name = $1, password = $2 WHERE id = $3"
		args := []interface{}{u.GetName(), u.GetPassword(), u.GetId()}
		var expectedResult pgconn.CommandTag
		expectedResult = append(expectedResult, '1')
		f.mockPool.EXPECT().Exec(context.Background(), queryStore, args...).Return(expectedResult, nil)

		// act
		err := f.userStorage.Update(context.Background(), u)

		// assert
		require.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		u := user.NewUser(2, "user1", "123456")
		queryStore := "UPDATE users SET name = $1, password = $2 WHERE id = $3"
		args := []interface{}{u.GetName(), u.GetPassword(), u.GetId()}
		var expectedResult pgconn.CommandTag
		expectedResult = append(expectedResult, '0')
		f.mockPool.EXPECT().Exec(context.Background(), queryStore, args...).Return(expectedResult, nil)

		// act
		err := f.userStorage.Update(context.Background(), u)

		// assert
		require.ErrorIs(t, err, userapp.ErrUserNotExists)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		u := user.NewUser(1, "user1", "123456")
		queryStore := "SELECT id, name, password FROM users WHERE id = $1"
		args := []interface{}{u.GetId()}
		columns := []string{"id", "name", "password"}
		pgxRows := pgxpoolmock.NewRows(columns).AddRow(u.GetId(), u.GetName(), u.GetPassword()).ToPgxRows()
		f.mockPool.EXPECT().Query(context.Background(), queryStore, args...).Return(pgxRows, nil)
		// act
		actualUser, err := f.userStorage.Get(context.Background(), u.GetId())

		// assert
		require.NoError(t, err)
		assert.Equal(t, &u, actualUser)
	})

	t.Run("fail", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		u := user.NewUser(2, "user1", "123456")
		queryStore := "SELECT id, name, password FROM users WHERE id = $1"
		args := []interface{}{u.GetId()}
		columns := []string{"id", "name", "password"}
		pgxRows := pgxpoolmock.NewRows(columns).ToPgxRows()
		f.mockPool.EXPECT().Query(context.Background(), queryStore, args...).Return(pgxRows, nil)
		// act
		actualUser, err := f.userStorage.Get(context.Background(), u.GetId())

		// assert
		assert.ErrorIs(t, err, userapp.ErrUserNotExists)
		assert.Equal(t, (*user.User)(nil), actualUser)
	})
}

func TestListUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		var users []user.User
		users = append(users, user.NewUser(1, "user1", "123456"))
		users = append(users, user.NewUser(2, "user2", "123456"))
		users = append(users, user.NewUser(3, "user3", "123456"))
		var offset, limit uint64 
		offset = 0
		limit = 3
		queryStore := fmt.Sprintf("SELECT id, name, password FROM users ORDER BY id LIMIT %d OFFSET %d", limit, offset)
		args := []interface{}{}
		columns := []string{"id", "name", "password"}
		rows := pgxpoolmock.NewRows(columns)
		for _, u := range users {
			rows.AddRow(u.GetId(), u.GetName(), u.GetPassword())
		}
		pgxRows := rows.ToPgxRows()
		f.mockPool.EXPECT().Query(context.Background(), queryStore, args...).Return(pgxRows, nil)
		
		// act
		actualUsers, err := f.userStorage.List(context.Background(), offset, limit)

		// assert
		require.NoError(t, err)
		assert.Equal(t, users, actualUsers)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		u := user.NewUser(1, "user1", "123456")
		queryStore := "DELETE FROM users WHERE id = $1"
		args := []interface{}{u.GetId()}
		var expectedResult pgconn.CommandTag
		expectedResult = append(expectedResult, '1')
		f.mockPool.EXPECT().Exec(context.Background(), queryStore, args...).Return(expectedResult, nil)

		// act
		err := f.userStorage.Delete(context.Background(), u.GetId())

		// assert
		require.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		u := user.NewUser(2, "user1", "123456")
		queryStore := "DELETE FROM users WHERE id = $1"
		args := []interface{}{u.GetId()}
		var expectedResult pgconn.CommandTag
		expectedResult = append(expectedResult, '0')
		f.mockPool.EXPECT().Exec(context.Background(), queryStore, args...).Return(expectedResult, nil)

		// act
		err := f.userStorage.Delete(context.Background(), u.GetId())

		// assert
		require.ErrorIs(t, err, userapp.ErrUserNotExists)
	})
}
