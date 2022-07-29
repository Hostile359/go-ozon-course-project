.PHONY: run
run:
	go run cmd/bot/main.go cmd/bot/server.go

grpc_client:
	go run client/client.go client/handler.go

build:
	go build -o bin/bot cmd/bot/main.go cmd/bot/server.go

prepare:
	mkdir -p config
	echo 'package config\n\nconst ApiKey = "$(APIKEY)"\n' > config/apikey.go

clean:
	rm -rf bin
	rm config/apikey.go

grpc_api:
	buf generate api

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
