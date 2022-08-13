package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/commander"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/commandhandler"
	// "gitlab.ozon.dev/Hostile359/homework-1/internal/storage/userstorage/memoryuserstore"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/storage/userstorage/pguserstore"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := NewConfig()
	if err != nil {
		log.Fatal("Error while loading config: ", err)
	}

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)
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
	// userStorage := memoryuserstore.New()

	userApp := userapp.New(userStorage)
	go runBot(userApp)
	go runREST()
	runGRPCServer(userApp)
}

func runBot(userApp *userapp.App) {
	log.Println("start bot")
	commandHandler := commandhandler.New(*userApp)

	cmd := commander.MustNew(commandHandler)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
