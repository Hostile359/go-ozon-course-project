package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/commentapi/commentvalidapi"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/userapi/uservalidapi"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SwaggerDir       = "./swagger"
)

func runGRPCServer(cfg *Config) {
	conns, err := grpc.Dial(cfg.DBGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conns.Close()

	userClient := pb.NewUserClient(conns)
	commentClient := pb.NewCommentClient(conns)

	listener, err := net.Listen("tcp", cfg.ValidGrpcPort)
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, uservalidapi.New(userClient))
	pb.RegisterCommentServer(grpcServer, commentvalidapi.New(commentClient))

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func runREST(cfg Config) {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, cfg.ValidGrpcPort, opts); err != nil {
		log.Fatalln(err)
	}
	if err := pb.RegisterCommentHandlerFromEndpoint(ctx, mux, cfg.ValidGrpcPort, opts); err != nil {
		log.Fatalln(err)
	}

	http_mux := http.NewServeMux()
	http_mux.Handle("/", mux)

	fs := http.FileServer(http.Dir(SwaggerDir))
	http_mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	if err := http.ListenAndServe(cfg.ValidHttpPort, http_mux); err != nil {
		log.Fatalln(err)
	}
}

