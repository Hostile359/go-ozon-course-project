//go:generate mockgen -source ../../../../pkg/api/user_grpc.pb.go -destination=./mocks/usergrpc.go -package=mock_usergrpc
package uservalidapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/counter"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrValidationArgs = errors.New("Bad argument")
)

type implementation struct {
	pb.UnimplementedUserServer
	client pb.UserClient
	syncProducer sarama.SyncProducer
	redisClient *redis.Client
}

func New(client pb.UserClient, syncProducer sarama.SyncProducer, redisClient *redis.Client) pb.UserServer {	
	return &implementation{
		client: client,
		syncProducer: syncProducer,
		redisClient: redisClient,
	}
}

func (i *implementation) UserAdd(ctx context.Context, in *pb.UserAddRequest) (*pb.UserAddResponse, error) {
	if err := checkName(in.GetName()); err != nil {
		counter.IncFailReq()
		counter.IncValidErr()
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := checkPassword(in.GetPassword()); err != nil {
		counter.IncFailReq()
		counter.IncValidErr()
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	
	u := user.NewUser(0, in.GetName(), in.GetPassword())
	serializedU, err := json.Marshal(u)
	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	par, off, err := i.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "add_users",
		Value: sarama.ByteEncoder(serializedU),
	})

	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("Sended msg to add_users topic, partition: %v, offset: %v", par, off)

	counter.IncSuccessReq()
	return &pb.UserAddResponse{}, nil
}

func (i *implementation) UserGet(ctx context.Context, in *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	newCtx, span := otel.Tracer(userapp.TracerName).Start(ctx, "userValidApi/UserGet")
	defer span.End()

	span.SetAttributes(attribute.String("id", strconv.FormatUint(in.GetId(), 10)))

	log.Infof("Get user with id=%v", in.GetId())
	res := i.redisClient.Get(strconv.FormatUint(in.GetId(), 10))
	if res.Err() != nil {
		log.Error(res.Err())
		counter.IncCacheMiss()

		return i.requestAddUserToRedis(newCtx, in.GetId())
	}

	counter.IncCacheHit()
	u := user.User{}
	err := res.Scan(&u)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.UserGetResponse{
		User: &pb.UserGetResponse_User{
				Id: uint64(u.GetId()),
				Name: u.GetName(),
				Password: u.GetPassword(),
			},
	}
	log.Infof("User: %v", u)
	counter.IncSuccessReq()
	return resp, nil
}

func (i *implementation) requestAddUserToRedis(ctx context.Context, id uint64, ) (*pb.UserGetResponse, error) {
	req := pb.UserGetRequest{
		Id: id,
	}
	_, err := i.client.UserGet(ctx, &req)
	if err != nil {
		counter.IncFailReq()
		if !errors.Is(err, userapp.ErrUserNotExists) {
			counter.IncInternalErr()
		}
		log.Error(err)
		return nil, err
	}
	
	return nil, status.Error(codes.NotFound, "Loading user into cache. Try again later")
}

func (i *implementation) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	newCtx, span := otel.Tracer(userapp.TracerName).Start(ctx, "userValidApi/UserList")
	defer span.End()

	span.SetAttributes(attribute.String("page", strconv.FormatUint(in.GetPage(), 10)))
	span.SetAttributes(attribute.String("perPage", strconv.FormatUint(in.GetPerPage(), 10)))

	subChannel := fmt.Sprintf("%v_%v_%v", time.Now() ,in.GetPage(), in.GetPerPage())
	newCtx = metadata.AppendToOutgoingContext(newCtx, userapp.PubSubChKey, subChannel)

	log.Infof("Subscribe to redis ch: %v", subChannel)
	pubSub := i.redisClient.Subscribe(subChannel)
	defer pubSub.Close()
	
	req := pb.UserListRequest{
		Page: in.GetPage(),
		PerPage: in.GetPerPage(),
	}
	log.Infof("Get UserList, page=%v, perPage=%v", in.GetPage(), in.GetPerPage())
	_, err := i.client.UserList(newCtx, &req)
	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, err
	}
	
	expectedListSize := in.GetPerPage()
	if expectedListSize == 0 {
		expectedListSize = userapp.PerPageDefault
	}
	res := make([]*pb.UserListResponse_User, 0, expectedListSize)
	
	ch := pubSub.Channel()
	for msg := range ch {
		log.Infof("[%v]: %v", msg.Channel, msg.Payload)
		if msg.Payload == "Done" {
			break
		}
		u := user.User{}
		err := u.UnmarshalBinary([]byte(msg.Payload))
		if err != nil {
			log.Error(err)
			continue
		}
		res = append(res, &pb.UserListResponse_User{
			Id: uint64(u.GetId()),
			Name: u.GetName(),
			Password: u.GetPassword(),
		})
	}

	log.Infof("UserList: %v", res)
	resp := &pb.UserListResponse{
		Users: res,
	}

	counter.IncSuccessReq()
	return resp, nil
}

func (i *implementation) UserUpdate(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	if err := checkName(in.GetName()); err != nil {
		counter.IncFailReq()
		counter.IncValidErr()
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := checkPassword(in.GetPassword()); err != nil {
		counter.IncFailReq()
		counter.IncValidErr()
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	u := user.NewUser(user.UserId(in.GetId()), in.GetName(), in.GetPassword())
	serializedU, err := json.Marshal(u)
	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	par, off, err := i.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "update_users",
		Value: sarama.ByteEncoder(serializedU),
	})

	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("Sended msg to update_users topic, partition: %v, offset: %v", par, off)

	i.deleteUserFromRedis(in.GetId())

	counter.IncSuccessReq()
	return &pb.UserUpdateResponse{}, nil
}

func (i *implementation) UserDelete(ctx context.Context, in *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	u := user.NewUser(user.UserId(in.GetId()), "", "")
	serializedU, err := json.Marshal(u)
	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	par, off, err := i.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "delete_users",
		Value: sarama.ByteEncoder(serializedU),
	})

	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("Sended msg to delete_users topic, partition: %v, offset: %v", par, off)

	i.deleteUserFromRedis(in.GetId())
	
	counter.IncSuccessReq()
	return &pb.UserDeleteResponse{}, nil
}

func (i *implementation) deleteUserFromRedis(id uint64) {
	log.Infof("Deleting user with id [%v] from redis", id)
	res := i.redisClient.Del(strconv.FormatUint(id, 10))
	if res.Err() == redis.Nil {
		log.Infof("User with id [%v] doesn't exist in redis", id)
	} else if res.Err() != nil {
		log.Error(res.Err())
	} else {
		log.Info("Done")
	}
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
