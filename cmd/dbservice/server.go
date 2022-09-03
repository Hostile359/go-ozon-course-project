package main

import (
	"context"
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/commentapi/commentdbapi"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/userapi/userdbapi"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/commentapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/consumer/userconsumer"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"google.golang.org/grpc"
)

func runGRPCServer(cfg *Config, userApp *userapp.App, commentApp *commentapp.App) {
	listener, err := net.Listen("tcp", cfg.DBGrpcPort)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, userdbapi.New(*userApp))
	pb.RegisterCommentServer(grpcServer, commentdbapi.New(*commentApp))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func runConsumer(cfg Config, userApp *userapp.App) {
	brokers := strings.Split(cfg.Brokers, ",")
	saramaCfg := sarama.NewConfig()
	income, err := sarama.NewConsumerGroup(brokers, "dbConsumer", saramaCfg)
	if err != nil {
		log.Fatalln(err)
	}
	userConsumer := userconsumer.New(*userApp)

	userTopics := []string{"add_users", "update_users", "delete_users"}
	for {
		if err := income.Consume(context.Background(), userTopics, userConsumer); err != nil {
			log.Errorf("on consume: <%v>", err)
			time.Sleep(time.Second * 5)
		}
	}
}
