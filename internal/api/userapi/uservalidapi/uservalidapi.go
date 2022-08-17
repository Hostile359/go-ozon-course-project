package uservalidapi

import (
	"context"

	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrValidationArgs = errors.New("Bad argument")
)

type implementation struct {
	pb.UnimplementedUserServer
	client pb.UserClient
}

func New(client pb.UserClient) pb.UserServer {
	return &implementation{
		client: client,
	}
}

func (i *implementation) UserAdd(ctx context.Context, in *pb.UserAddRequest) (*pb.UserAddResponse, error) {
	if err := checkName(in.GetName()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := checkPassword(in.GetPassword()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	req := pb.UserAddRequest{
		Name: in.GetName(),
		Password: in.GetPassword(),
	}

	return i.client.UserAdd(ctx, &req)
}

func (i *implementation) UserGet(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	req := pb.UserGetRequest{
		Id: in.GetId(),
	}

	return i.client.UserGet(ctx, &req)
}

func (i *implementation) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	req := pb.UserListRequest{
		Page: in.GetPage(),
		PerPage: in.GetPerPage(),
	}
	
	return i.client.UserList(ctx, &req)
}

func (i *implementation) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	if err := checkName(in.GetName()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := checkPassword(in.GetPassword()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	req := pb.UserUpdateRequest{
		Id: in.GetId(),
		Name: in.GetName(),
		Password: in.GetPassword(),
	}

	return i.client.UserUpdate(ctx, &req)
}

func (i *implementation) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	req := pb.UserDeleteRequest{
		Id: in.GetId(),
	}
	return i.client.UserDelete(ctx, &req)
}

func checkName(name string) error {
	if len(name) == 0 || len(name) > 10 {
		return errors.Wrapf(ErrValidationArgs, "<%v>, len should be from 1 to 10", name)
	}
	return nil
}

func checkPassword(password string) error {
	if len(password) < 6 || len(password) > 10 {
		return errors.Wrapf(ErrValidationArgs, "<%v>, len should be from 6 to 10", password)
	}	
	return nil
}
