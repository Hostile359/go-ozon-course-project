package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runGRPCServer(userApp *userapp.App) {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, api.New(*userApp))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func runREST() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln(err)
	}
}

