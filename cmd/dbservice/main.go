package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/commentapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"

	// "gitlab.ozon.dev/Hostile359/homework-1/internal/storage/userstorage/memoryuserstore"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/storage/commentstorage/pgcommentstore"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/storage/userstorage/pguserstore"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := NewConfig()
	if err != nil {
		log.Fatal("Error while loading config: ", err)
	}

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.DBPort, cfg.User, cfg.Password, cfg.DBname)
	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("ping database error", err)
	}

	config := pool.Config()
	config.MaxConnIdleTime = cfg.MaxConnIdleTime
	config.MaxConnLifetime = cfg.MaxConnLifetime
	config.MinConns = cfg.MinConns
	config.MaxConns = cfg.MaxConns
	userStorage := pguserstore.New(pool)
	commentStorage := pgcommentstore.New(pool)
	// userStorage := memoryuserstore.New()

	userApp := userapp.New(userStorage)
	commentApp := commentapp.New(commentStorage, *userApp)
	// go runREST()
	go runConsumer(*cfg, userApp)
	runGRPCServer(cfg, userApp, commentApp)
}
