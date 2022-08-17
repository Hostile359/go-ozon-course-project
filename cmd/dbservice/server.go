package main

import (
	"log"
	"net"

	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/commentapi/commentdbapi"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/userapi/userdbapi"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/commentapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
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
