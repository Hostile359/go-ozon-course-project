package main

import (
	"log"

	"gitlab.ozon.dev/Hostile359/homework-1/internal/commander"
)

func main() {
	log.Println("start main")
	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}
}
