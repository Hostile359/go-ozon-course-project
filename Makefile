DB_MAIN:=$(CURDIR)/cmd/dbservice
.PHONY: run-db, build-db
run-db:
	go run $(DB_MAIN)/main.go $(DB_MAIN)/server.go $(DB_MAIN)/pgconfig.go --local

build-db:
	go build -o bin/db-service $(DB_MAIN)/main.go $(DB_MAIN)/server.go $(DB_MAIN)/pgconfig.go

VALID_MAIN:=$(CURDIR)/cmd/validservice
.PHONY: run-valid build-valid
run-valid:
	go run $(VALID_MAIN)/main.go $(VALID_MAIN)/server.go $(VALID_MAIN)/config.go --local

build-valid:
	go build -o bin/valid-service $(VALID_MAIN)/main.go $(VALID_MAIN)/server.go $(VALID_MAIN)/config.go

build: build-db build-valid

grpc_client:
	go run client/client.go client/handler.go

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

.PHONY: .test
.test:
	$(info Running tests...)
	go test -v $$(go list ./... | grep -v -E '/homework-1/pkg/(api)')

.PHONY: cover
cover:
	go test -v $$(go list ./... | grep -v -E '/homework-1/pkg/(api)') -covermode=count -coverprofile=/tmp/c.out
	go tool cover -html=/tmp/c.out
