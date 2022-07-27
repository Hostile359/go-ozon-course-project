package api

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	// "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type implementation struct {
	pb.UnimplementedAdminServer
	userApp userapp.App
}

func New(userApp userapp.App) pb.AdminServer {
	return &implementation{
		userApp: userApp,
	}
}

func (i *implementation) UserAdd(ctx context.Context, in *pb.UserAddRequest) (*pb.UserAddResponse, error) {
	u := user.NewUser(user.UserId(in.GetId()), in.GetName(), in.GetPassword())
	if err := i.userApp.Add(u); err != nil {
		if errors.Is(err, userapp.ErrValidationArgs) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, userapp.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserAddResponse{}, nil
}

func (i *implementation) UserGet(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	u, err := i.userApp.Get(user.UserId(in.GetId()))
	if err != nil {
		if errors.Is(err, userapp.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &pb.UserGetResponse{
		User: &pb.UserGetResponse_User{
			Id: uint64(u.GetId()),
			Name: u.GetName(),
			Password: u.GetPassword(),
		},
	}, nil
}

func (i *implementation) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	usersList := i.userApp.List()
	res := make([]*pb.UserListResponse_User, 0, len(usersList))

	for _, u := range usersList {
		res = append(res, &pb.UserListResponse_User{
			Id: uint64(u.GetId()),
			Name: u.GetName(),
			Password: u.GetPassword(),
		})
	}

	return &pb.UserListResponse{
		Users: res,
	}, nil
}

func (i *implementation) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	u := user.NewUser(user.UserId(in.GetId()), in.GetName(), in.GetPassword())
	if err := i.userApp.Update(u); err != nil {
		if errors.Is(err, userapp.ErrValidationArgs) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		} else if errors.Is(err, userapp.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserUpdateResponse{}, nil
}

func (i *implementation) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	if err := i.userApp.Delete(user.UserId(in.GetId())); err != nil {
		if errors.Is(err, userapp.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.UserDeleteResponse{}, nil
}
