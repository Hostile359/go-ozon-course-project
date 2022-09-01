package uservalidapi

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)

		req := pb.UserAddRequest{
			Name: "user1",
			Password: "123456",
		}
		expectedResp := &pb.UserAddResponse{}
		f.producer.ExpectSendMessageAndSucceed()

		// act
		resp, err := f.userValidApi.UserAdd(context.Background(), &req)

		// assert
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("wrong name len 0", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserAddRequest{
				Name: "",
				Password: "123456",
			}
	
			// act
			resp, err := f.userValidApi.UserAdd(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong name len > 10", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserAddRequest{
				Name: "user1234567",
				Password: "123456",
			}
	
			// act
			resp, err := f.userValidApi.UserAdd(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len < 6", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserAddRequest{
				Name: "user1",
				Password: "12345",
			}
	
			// act
			resp, err := f.userValidApi.UserAdd(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len > 10", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserAddRequest{
				Name: "user1",
				Password: "1234567890a",
			}
	
			// act
			resp, err := f.userValidApi.UserAdd(context.Background(), &req)
	
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
		f := setUp(t)

		req := pb.UserGetRequest{
			Id: 1,
		}
		expectedResp:= &pb.UserGetResponse{
			User: &pb.UserGetResponse_User{
				Id: 1,
				Name: "user1",
				Password: "123456",
			},
		}
		f.client.EXPECT().UserGet(gomock.Any(), &req).Return(expectedResp, nil)

		// act
		resp, err := f.userValidApi.UserGet(context.Background(), &req)

		// assert
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})
}

func TestListUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)

		req := pb.UserListRequest{
			Page: 1,
			PerPage: 3,
		}
		expectedResp:= &pb.UserListResponse{
			Users: []*pb.UserListResponse_User{
				{Id: 1, Name: "user1", Password: "123456"},
				{Id: 2, Name: "user2", Password: "678901"},
				{Id: 3, Name: "user3", Password: "Sdasddd"},
			},
		}
		f.client.EXPECT().UserList(gomock.Any(), &req).Return(expectedResp, nil)

		// act
		resp, err := f.userValidApi.UserList(context.Background(), &req)

		// assert
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		// arrange
		f := setUp(t)

		req := pb.UserUpdateRequest{
			Id: 1,
			Name: "user1",
			Password: "qwerty123",
		}
		expectedResp := &pb.UserUpdateResponse{}
		f.producer.ExpectSendMessageAndSucceed()

		// act
		resp, err := f.userValidApi.UserUpdate(context.Background(), &req)

		// assert
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("wrong name len 0", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "",
				Password: "qwerty123",
			}
	
			// act
			resp, err := f.userValidApi.UserUpdate(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong name len > 10", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "user1234567",
				Password: "qwerty123",
			}
	
			// act
			resp, err := f.userValidApi.UserUpdate(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len < 6", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "user1",
				Password: "qwert",
			}
	
			// act
			resp, err := f.userValidApi.UserUpdate(context.Background(), &req)
	
			// assert
			require.Error(t, err)
			assert.Equal(t, codes.InvalidArgument, status.Code(err))
			assert.Nil(t, resp)
		})

		t.Run("wrong password len > 10", func(t *testing.T) {
			// arrange
			f := setUp(t)
	
			req := pb.UserUpdateRequest{
				Id: 1,
				Name: "user1",
				Password: "qwerty12345",
			}
	
			// act
			resp, err := f.userValidApi.UserUpdate(context.Background(), &req)
	
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
		f := setUp(t)

		req := pb.UserDeleteRequest{
			Id: 1,
		}
		expectedResp:= &pb.UserDeleteResponse{}
		f.producer.ExpectSendMessageAndSucceed()

		// act
		resp, err := f.userValidApi.UserDelete(context.Background(), &req)

		// assert
		require.NoError(t, err)
		assert.Equal(t, expectedResp, resp)
	})
}
