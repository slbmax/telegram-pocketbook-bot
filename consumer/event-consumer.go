package consumer

import (
	"github.com/slbmax/telegram-pocketbook-bot/events"
	"log"
	"time"
)

type EventConsumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return EventConsumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (e EventConsumer) Start() error {
	log.Println("Bot started\nListening to updates...")

	for {
		events, err := e.fetcher.Fetch(e.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(events) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := e.handleEvents(events); err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}
	}
}

func (e *EventConsumer) handleEvents(evs []events.Event) error {
	for _, event := range evs {
		log.Printf("[EVENT] got new event: %s", event.Text)

		if err := e.processor.Process(event); err != nil {
			log.Printf("[ERR] can`t handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
