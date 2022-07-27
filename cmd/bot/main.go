package main

import (
	"log"

	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/commander"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/commandhandler"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/storage/userstorage/memoryuserstore"
)

func main() {
	runGRPCServer()
}

func runBot() {
	log.Println("start main")
	userStorage := memoryuserstore.New()
	userApp := userapp.New(userStorage)
	commandHandler := commandhandler.New(*userApp)

	cmd := commander.MustNew(commandHandler)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
