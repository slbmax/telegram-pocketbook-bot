package events

import "github.com/slbmax/telegram-pocketbook-bot/cli"

type EventProcessor struct {
	tg     *cli.Client
	offset int
}
