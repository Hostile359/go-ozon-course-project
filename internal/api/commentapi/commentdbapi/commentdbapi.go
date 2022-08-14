package commentdbapi

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/commentapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/comment"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type implementation struct {
	pb.UnimplementedCommentServer
	commentApp commentapp.App
}

func New(commentApp commentapp.App) pb.CommentServer {
	return &implementation{
		commentApp: commentApp,
	}
}

func (i *implementation) CommentAdd(ctx context.Context, in *pb.CommentAddRequest) (*pb.CommentAddResponse, error) {
	c := comment.NewComment(0, in.GetComment(), user.UserId(in.GetUserId()))
	if err := i.commentApp.Add(ctx, c); err != nil {
		if errors.Is(err, userapp.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CommentAddResponse{}, nil
}

func (i *implementation) CommentGet(ctx context.Context, in *pb.CommentGetRequest) (*pb.CommentGetResponse, error) {
	c, err := i.commentApp.Get(ctx, comment.CommentId(in.GetId()))
	if err != nil {
		if errors.Is(err, commentapp.ErrCommentNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	
	return &pb.CommentGetResponse{
		Comment: &pb.CommentGetResponse_Comment{
			Id: uint64(c.GetId()),
			Comment: c.GetComment(),
			UserId: uint64(c.GetUserId()),
		},
	}, nil
}

func (i *implementation) CommentList(ctx context.Context, in *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	commentsList, err := i.commentApp.List(ctx, in.Page, in.PerPage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]*pb.CommentListResponse_Comment, 0, len(commentsList))

	for _, c := range commentsList {
		res = append(res, &pb.CommentListResponse_Comment{
			Id: uint64(c.GetId()),
			Comment: c.GetComment(),
			UserId: uint64(c.GetUserId()),
		})
	}

	return &pb.CommentListResponse{
		Comments: res,
	}, nil
}

func (i *implementation) CommentUpdate(ctx context.Context, in *pb.CommentUpdateRequest) (*pb.CommentUpdateResponse, error) {
	c := comment.NewComment(comment.CommentId(in.GetId()), in.GetComment(), user.UserId(in.GetUserId()))
	if err := i.commentApp.Update(ctx, c); err != nil {
		if errors.Is(err, commentapp.ErrCommentNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CommentUpdateResponse{}, nil
}

func (i *implementation) CommentDelete(ctx context.Context, in *pb.CommentDeleteRequest) (*pb.CommentDeleteResponse, error) {
	if err := i.commentApp.Delete(ctx, comment.CommentId(in.GetId())); err != nil {
		if errors.Is(err, commentapp.ErrCommentNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CommentDeleteResponse{}, nil
}
