//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"gitlab.ozon.dev/Hostile359/homework-1/tests/fixtures"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func TestAddUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		req := pb.UserAddRequest{
			Name: "user1",
			Password: "123456",
		}

		//act
		_, err := UserClient.UserAdd(context.Background(), &req)

		//assert
		require.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("wrong name len 0", func(t *testing.T) {
			// arrange
			req := pb.UserAddRequest{
				Name: "",
				Password: "123456",
			}
	
			// act
			resp, err := UserClient.UserAdd(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})
		t.Run("wrong name len > 10", func(t *testing.T) {
			// arrange
	
			req := pb.UserAddRequest{
				Name: "user1234567",
				Password: "123456",
			}
	
			// act
			resp, err := UserClient.UserAdd(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len < 6", func(t *testing.T) {
			// arrange
	
			req := pb.UserAddRequest{
				Name: "user1",
				Password: "12345",
			}
	
			// act
			resp, err := UserClient.UserAdd(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len > 10", func(t *testing.T) {
			// arrange
	
			req := pb.UserAddRequest{
				Name: "user1",
				Password: "1234567890a",
			}
	
			// act
			resp, err := UserClient.UserAdd(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})
	})
}

func TestGetUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		u := fixtures.User().Id(1).Name("user1").Password("123456").V()
		req := pb.UserGetRequest{
			Id: uint64(u.Id),
		}

		// act
		var resp *pb.UserGetResponse
		var err error
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			resp, err = UserClient.UserGet(context.Background(), &req)
			if err == nil {
				break
			}
		}

		// assert
		require.NoError(t, err)
		assert.Equal(t, uint64(u.Id), resp.User.Id)
		assert.Equal(t, u.Name, resp.User.Name)
		assert.Equal(t, u.Password, resp.User.Password)
	})

	t.Run("error", func(t *testing.T) {
		// arrange
		u := fixtures.User().Id(2).Name("user2").Password("123456").V()
		req := pb.UserGetRequest{
			Id: uint64(u.Id),
		}

		// act
		resp, err := UserClient.UserGet(context.Background(), &req)

		// assert
		require.Error(t, err)
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Nil(t, resp)
	})
}

func TestListUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		req := pb.UserListRequest{}

		// act
		var resp *pb.UserListResponse
		var err error
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			resp, err = UserClient.UserList(context.Background(), &req)
			if err == nil {
				break
			}
		}

		// assert
		require.NoError(t, err)
		assert.Equal(t, 1, len(resp.Users))
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//arrange
		req := pb.UserUpdateRequest{
			Id: 1,
			Name: "user1",
			Password: "12345678",
		}

		//act
		_, err := UserClient.UserUpdate(context.Background(), &req)

		//assert
		require.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("wrong name len 0", func(t *testing.T) {
			// arrange
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "",
				Password: "123456",
			}
	
			// act
			resp, err := UserClient.UserUpdate(context.Background(), &req)
			
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})
		t.Run("wrong name len > 10", func(t *testing.T) {
			// arrange
	
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "user1234567",
				Password: "123456",
			}
	
			// act
			resp, err := UserClient.UserUpdate(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len < 6", func(t *testing.T) {
			// arrange
	
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "user1",
				Password: "12345",
			}
	
			// act
			resp, err := UserClient.UserUpdate(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len > 10", func(t *testing.T) {
			// arrange
	
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "user1",
				Password: "1234567890a",
			}
	
			// act
			resp, err := UserClient.UserUpdate(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		req := pb.UserDeleteRequest{
			Id: 1,
		}

		// act
		_, err := UserClient.UserDelete(context.Background(), &req)

		// assert
		require.NoError(t, err)
	})
}
