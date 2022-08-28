package main

import (
	"flag"
	"log"
)

func main() {

}

func mustToken() string {
	token := flag.String(
		"telegram-bot-token",
		"",
		"token for telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not provided")
	}

	return *token
}
