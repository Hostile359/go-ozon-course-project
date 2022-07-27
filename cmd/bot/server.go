package main

import (
	"net"

	"gitlab.ozon.dev/Hostile359/homework-1/internal/api"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/storage/userstorage/memoryuserstore"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"google.golang.org/grpc"
)

func runGRPCServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	userStorage := memoryuserstore.New()
	userApp := userapp.New(userStorage)
	pb.RegisterAdminServer(grpcServer, api.New(*userApp))

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
