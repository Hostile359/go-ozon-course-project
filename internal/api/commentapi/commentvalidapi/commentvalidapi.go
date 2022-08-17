package commentvalidapi

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
	pb.UnimplementedCommentServer
	client pb.CommentClient
}

func New(client pb.CommentClient) pb.CommentServer {
	return &implementation{
		client: client,
	}
}

func (i *implementation) CommentAdd(ctx context.Context, in *pb.CommentAddRequest) (*pb.CommentAddResponse, error) {
	if len(in.GetComment()) == 0 {
		err := errors.Wrap(ErrValidationArgs, "Comment can't be empty")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	req := pb.CommentAddRequest{
		Comment: in.GetComment(),
		UserId: in.GetUserId(),
	}
	
	return i.client.CommentAdd(ctx, &req)
}

func (i *implementation) CommentGet(ctx context.Context, in *pb.CommentGetRequest) (*pb.CommentGetResponse, error) {
	req := pb.CommentGetRequest{
		Id: in.GetId(),
	}

	return i.client.CommentGet(ctx, &req)
}

func (i *implementation) CommentList(ctx context.Context, in *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	req := pb.CommentListRequest{
		Page: in.GetPage(),
		PerPage: in.GetPerPage(),
	}
	
	return i.client.CommentList(ctx, &req)
}

func (i *implementation) CommentUpdate(ctx context.Context, in *pb.CommentUpdateRequest) (*pb.CommentUpdateResponse, error) {
	if len(in.GetComment()) == 0 {
		err := errors.Wrap(ErrValidationArgs, "Comment can't be empty")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	req := pb.CommentUpdateRequest{
		Id: in.GetId(),
		Comment: in.GetComment(),
		UserId: in.GetUserId(),
	}
	
	return i.client.CommentUpdate(ctx, &req)
}

func (i *implementation) CommentDelete(ctx context.Context, in *pb.CommentDeleteRequest) (*pb.CommentDeleteResponse, error) {
	req := pb.CommentDeleteRequest{
		Id: in.GetId(),
	}
	
	return i.client.CommentDelete(ctx, &req)
}
