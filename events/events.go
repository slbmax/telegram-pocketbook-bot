package events

import (
	"github.com/slbmax/telegram-pocketbook-bot/cli"
	errors "github.com/slbmax/telegram-pocketbook-bot/lib/custom-errors"
	"github.com/slbmax/telegram-pocketbook-bot/storage"
)

type EventProcessor struct {
	tg      *cli.Client
	offset  int
	storage storage.Storage
}

var ErrUnknownEventType = errors.New("unknown event type")

func New(client *cli.Client, storage storage.Storage) *EventProcessor {
	return &EventProcessor{
		tg:      client,
		storage: storage,
	}
}

func (p *EventProcessor) Fetch(limit int) ([]Event, error) {
	updates, err := p.tg.Updates(limit, p.offset)
	if err != nil {
		return nil, errors.Wrap("can`t get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]Event, len(updates))

	for i, update := range updates {
		res[i] = updateToEvent(update)
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *EventProcessor) Process(event Event) error {
	switch event.Type {
	case MessageEvent:
		return p.processMessage(event)
	default:
		return errors.Wrap("can`t process message", ErrUnknownEventType)
	}
}

func (p EventProcessor) processMessage(event Event) error {
	if err := p.processCommand(event.Text, event.Username, event.ChatID); err != nil {
		return err
	}

	return nil
}

func updateToEvent(update cli.Update) Event {
	updateType := fetchType(update)

	res := Event{
		Type: updateType,
		Text: fetchText(update),
	}

	if updateType == MessageEvent {
		res.ChatID = update.Message.Chat.ID
		res.Username = update.Message.From.Username
	}

	return res
}

func fetchText(upd cli.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd cli.Update) Type {
	if upd.Message == nil {
		return UnknownEvent
	}

	return MessageEvent
}
