package events

import (
	"github.com/slbmax/telegram-pocketbook-bot/cli"
	errors "github.com/slbmax/telegram-pocketbook-bot/lib/custom-errors"
	"github.com/slbmax/telegram-pocketbook-bot/storage"
	"log"
	"strings"
)

const (
	GetCommand   = "/show"
	HelpCommand  = "/help"
	StartCommand = "/start"
	SaveCommand  = "/add"
)

func (p EventProcessor) processCommand(text, username string, chatID int) error {
	log.Printf("got new command: '%s' from '%s'", text, username)

	isAdd, text := p.isAddCommand(text)
	if isAdd {
		return p.saveNote(chatID, text, username)
	}

	switch text {
	case GetCommand:
		return p.sendAll(chatID, username)
	case HelpCommand:
		return p.sendHelp(chatID)
	case StartCommand:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p EventProcessor) isAddCommand(text string) (bool, string) {
	values := strings.Split(text, " ")

	if len(values) < 2 {
		return false, text
	}

	if values[0] == SaveCommand {
		return true, strings.Join(values[1:len(values)], " ")
	}

	return false, text
}

func (p EventProcessor) saveNote(chatID int, text, username string) (err error) {
	defer func() { err = errors.WrapIfErr("can`t execute save note command", err) }()

	sendMessage := NewMessageSender(chatID, p.tg)

	note := &storage.Note{
		Value:    text,
		UserName: username,
	}

	isExists, err := p.storage.IsNoteExists(note)
	if err != nil {
		return err
	}

	if isExists {
		return sendMessage(msgAlreadyExists)
	}

	if err = p.storage.SaveNote(note); err != nil {
		return err
	}

	return sendMessage(msgSaved)
}

func (p EventProcessor) sendAll(chatID int, username string) (err error) {
	defer func() { err = errors.WrapIfErr("can't do command: can't send all notes", err) }()

	notes, err := p.storage.GetAllNotes(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedNotes) {
		return p.tg.SendMessage(chatID, msgNoSavedNotes)
	}

	sb := strings.Builder{}
	for _, val := range *notes {
		sb.WriteString(val.Value + "\n")
	}
	message := sb.String()
	message = message[:len(message)-1]

	return p.tg.SendMessage(chatID, message)
}

func (p EventProcessor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p EventProcessor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello)
}

func NewMessageSender(chatID int, tg *cli.Client) func(string) error {
	return func(msg string) error {
		return tg.SendMessage(chatID, msg)
	}
}
