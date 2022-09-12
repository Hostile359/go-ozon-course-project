package userdbapi

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type implementation struct {
	pb.UnimplementedUserServer
	userApp userapp.App
	redisClient *redis.Client
}

func New(userApp userapp.App, redisClient *redis.Client) pb.UserServer {
	return &implementation{
		userApp: userApp,
		redisClient: redisClient,
	}
}

func (i *implementation) UserAdd(ctx context.Context, in *pb.UserAddRequest) (*pb.UserAddResponse, error) {
	// u := user.NewUser(0, in.GetName(), in.GetPassword())
	// if err := i.userApp.Add(ctx, u); err != nil {
	// 	if errors.Is(err, userapp.ErrUserExists) {
	// 		return nil, status.Error(codes.AlreadyExists, err.Error())
	// 	}
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	return &pb.UserAddResponse{}, nil
}

func (i *implementation) UserGet(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	newCtx, span := otel.Tracer(userapp.TracerName).Start(ctx, "userDbApi/UserGet")
	defer span.End()

	span.SetAttributes(attribute.String("id", strconv.FormatUint(in.GetId(), 10)))

	log.Infof("Get user with id=%v", in.GetId())
	u, err := i.userApp.Get(newCtx, user.UserId(in.GetId()))
	if err != nil {
		log.Error(err)
		if errors.Is(err, userapp.ErrUserNotExists) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Infof("User: %v", u)

	res := i.redisClient.Set(strconv.FormatUint(in.GetId(), 10), u, 10*time.Second)
	if res.Err() != nil {
		log.Errorf("%v, while set user to redis", res.Err())
		return nil, status.Error(codes.Internal, res.Err().Error())
	}

	return &pb.UserGetResponse{}, nil
}

func (i *implementation) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	newCtx, span := otel.Tracer(userapp.TracerName).Start(ctx, "userDbApi/UserList")
	defer span.End()

	span.SetAttributes(attribute.String("page", strconv.FormatUint(in.GetPage(), 10)))
	span.SetAttributes(attribute.String("perPage", strconv.FormatUint(in.GetPerPage(), 10)))

	var pubChannel string
	md, ok := metadata.FromIncomingContext(newCtx)
	if !ok {
		log.Warn("Metadata is empty")
		pubChannel = fmt.Sprintf("%v_%v", in.GetPage(), in.GetPerPage())
	} else {
		pubChannel = md.Get(userapp.PubSubChKey)[0]
	}

	log.Infof("Get UserList, page=%v, perPage=%v", in.GetPage(), in.GetPerPage())
	usersList, err := i.userApp.List(newCtx, in.Page, in.PerPage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("Publish users to: %v", pubChannel)
	for _, u := range usersList {
		log.Infof("Publishing user %v...", u)
		err := i.redisClient.Publish(pubChannel, u).Err()
		if err != nil {
			log.Error(err)
		}
	}
	err = i.redisClient.Publish(pubChannel, "Done").Err()
	if err != nil {
		log.Error(err)
	}
	log.Infof("Done!")

	return &pb.UserListResponse{}, nil
}

func (i *implementation) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	// u := user.NewUser(user.UserId(in.GetId()), in.GetName(), in.GetPassword())
	// if err := i.userApp.Update(ctx, u); err != nil {
	// 	if errors.Is(err, userapp.ErrUserNotExists) {
	// 		return nil, status.Error(codes.NotFound, err.Error())
	// 	}
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	return &pb.UserUpdateResponse{}, nil
}

func (i *implementation) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	// if err := i.userApp.Delete(ctx, user.UserId(in.GetId())); err != nil {
	// 	if errors.Is(err, userapp.ErrUserNotExists) {
	// 		return nil, status.Error(codes.NotFound, err.Error())
	// 	}
	// 	return nil, status.Error(codes.Internal, err.Error())
	// }

	return &pb.UserDeleteResponse{}, nil
}
