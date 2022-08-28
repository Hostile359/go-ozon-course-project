//go:generate mockgen -source ../../../../pkg/api/user_grpc.pb.go -destination=./mocks/usergrpc.go -package=mock_usergrpc
package uservalidapi

import (
	"context"
	"encoding/json"
	
	log "github.com/sirupsen/logrus"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/counter"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/entities/user"
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
	syncProducer sarama.SyncProducer
}

func New(client pb.UserClient, syncProducer sarama.SyncProducer) pb.UserServer {
	return &implementation{
		client: client,
		syncProducer: syncProducer,
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
	req := pb.UserGetRequest{
		Id: in.GetId(),
	}
	log.Infof("Get user with id=%v", in.GetId())
	resp, err := i.client.UserGet(ctx, &req)
	if err != nil {
		counter.IncFailReq()
		if !errors.Is(err, userapp.ErrUserNotExists) {
			counter.IncInternalErr()
		}
		log.Error(err)
		return nil, err
	}
	log.Infof("User: %v", resp.User)

	counter.IncSuccessReq()
	return resp, nil
}

func (i *implementation) UserList(ctx context.Context, in *pb.UserListRequest) (*pb.UserListResponse, error) {
	req := pb.UserListRequest{
		Page: in.GetPage(),
		PerPage: in.GetPerPage(),
	}
	log.Infof("Get UserList, page=%v, perPage=%v", in.GetPage(), in.GetPerPage())
	resp, err := i.client.UserList(ctx, &req)
	if err != nil {
		counter.IncFailReq()
		counter.IncInternalErr()
		log.Error(err)
		return nil, err
	}
	log.Infof("UserList: %v", resp.Users)

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

	counter.IncSuccessReq()
	return &pb.UserDeleteResponse{}, nil
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
