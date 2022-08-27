// +build integration

package tests

import (
	"log"

	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"gitlab.ozon.dev/Hostile359/homework-1/tests/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient pb.UserClient
)

func init() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error while loading config: ", err)
	}

	conns, err := grpc.Dial(cfg.GrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	UserClient = pb.NewUserClient(conns)
}
