package main

import (
	"github.com/slbmax/telegram-pocketbook-bot/cli"
	"github.com/slbmax/telegram-pocketbook-bot/consumer"
	"github.com/slbmax/telegram-pocketbook-bot/events"
	"github.com/slbmax/telegram-pocketbook-bot/storage"
	"log"
	"os"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "./storage/notes"
	batchSize   = 100
)

func main() {
	eventsProcessor := events.New(
		cli.New(tgBotHost, mustToken()),
		storage.New(storagePath),
	)

	log.Println("Starting telegram bot...")

	consumer := consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Service is stopped", err)
	}
}

func mustToken() string {
	//token := flag.String(
	//	"telegram-bot-token",
	//	"",
	//	"token for telegram bot",
	//)
	//
	//flag.Parse()

	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("token is not provided")
	}

	return token
}
