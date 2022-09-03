package main

import (
	log "github.com/sirupsen/logrus"
	// "gitlab.ozon.dev/Hostile359/homework-1/internal/app/commentapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/app/userapp"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/commander"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/commandhandler"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal("Error while loading config: ", err)
	}

	// go runBot(userApp)
	go runREST(*cfg)
	runGRPCServer(cfg)
}

func runBot(userApp *userapp.App) {
	log.Println("start bot")
	commandHandler := commandhandler.New(*userApp)

	cmd := commander.MustNew(commandHandler)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
