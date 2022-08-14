package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/commentapi"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/userapi"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/commentapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	pb "gitlab.ozon.dev/Hostile359/homework-1/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SwaggerDir       = "./swagger"
)

func runGRPCServer(userApp *userapp.App, commentApp *commentapp.App) {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, userapi.New(*userApp))
	pb.RegisterCommentServer(grpcServer, commentapi.New(*commentApp))

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
	if err := pb.RegisterUserHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatalln(err)
	}
	if err := pb.RegisterCommentHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatalln(err)
	}

	http_mux := http.NewServeMux()
	http_mux.Handle("/", mux)

	fs := http.FileServer(http.Dir(SwaggerDir))
	http_mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	if err := http.ListenAndServe(":8080", http_mux); err != nil {
		log.Fatalln(err)
	}
}

