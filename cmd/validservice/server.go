package main

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/Shopify/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/api/commentapi/commentvalidapi"
	_ "gitlab.ozon.dev/Hostile359/homework-1/internal/counter"
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

	listener, err := net.Listen("tcp", cfg.ValidGrpcPort)
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	err = registerServices(cfg ,conns, grpcServer)
	if err != nil {
		log.Fatalln(err)
	}

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

	http.Handle("/", mux)

	fs := http.FileServer(http.Dir(SwaggerDir))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	if err := http.ListenAndServe(cfg.ValidHttpPort, nil); err != nil {
		log.Fatalln(err)
	}
}

func registerServices(cfg *Config, conns *grpc.ClientConn, grpcServer *grpc.Server) error {
	brokers := strings.Split(cfg.Brokers, ",")
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	userClient := pb.NewUserClient(conns)
	syncProducer, err := sarama.NewSyncProducer(brokers, saramaCfg)
	if err != nil {
		return err
	}
	pb.RegisterUserServer(grpcServer, uservalidapi.New(userClient, syncProducer))

	commentClient := pb.NewCommentClient(conns)
	pb.RegisterCommentServer(grpcServer, commentvalidapi.New(commentClient))

	return nil
}
