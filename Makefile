.PHONY: run
run:
	go run cmd/bot/main.go cmd/bot/server.go cmd/bot/pgconfig.go --local

grpc_client:
	go run client/client.go client/handler.go

build:
	go build -o bin/bot cmd/bot/main.go cmd/bot/server.go cmd/bot/pgconfig.go

prepare:
	mkdir -p config
	echo 'package config\n\nconst ApiKey = "$(APIKEY)"\n' > config/apikey.go

clean:
	rm -rf bin
	rm config/apikey.go

grpc_api:
	buf mod update
	buf generate api

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: up-db
up-db:
	docker-compose build
	docker-compose up -d postgres

.PHONY: down-db
down-db:
	docker-compose down

MIGRATIONS_DIR:=./migrations
.PHONY: migration
migration:
	bin/goose -dir=${MIGRATIONS_DIR} create $(NAME) sql 
