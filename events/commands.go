package events

import (
	"fmt"
	"github.com/slbmax/telegram-pocketbook-bot/cli"
	errors "github.com/slbmax/telegram-pocketbook-bot/lib/custom-errors"
	"github.com/slbmax/telegram-pocketbook-bot/storage"
	"log"
	"strings"
)

const (
	GetCommand    = "/show"
	HelpCommand   = "/help"
	StartCommand  = "/start"
	SaveCommand   = "/add"
	RemoveCommand = "/remove"
)

func (p EventProcessor) processCommand(text, username string, chatID int) error {
	log.Printf("got new command: '%s' from '%s'", text, username)

	isAdd, text := p.isAddCommand(text)
	if isAdd {
		return p.saveNote(chatID, text, username)
	}

	// Command has invalid format
	if strings.HasPrefix(text, SaveCommand) {
		return p.tg.SendMessage(chatID, msgInvalidCommand)
	}

	isRemove, description := p.isRemoveCommand(text)
	if isRemove {
		return p.removeNote(chatID, description, username)
	}

	// Command has invalid format
	if strings.HasPrefix(text, RemoveCommand) {
		return p.tg.SendMessage(chatID, msgInvalidCommand)
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

	if len(values) < 3 {
		return false, text
	}

	if values[0] == SaveCommand {
		return true, strings.Join(values[1:len(values)], " ")
	}

	return false, text
}

func (p EventProcessor) isRemoveCommand(text string) (bool, string) {
	values := strings.Split(text, " ")

	if len(values) != 2 {
		return false, text
	}

	if values[0] == RemoveCommand {
		return true, values[1]
	}

	return false, text
}

func (p EventProcessor) saveNote(chatID int, text, username string) (err error) {
	defer func() { err = errors.WrapIfErr("can`t execute save note command", err) }()

	sendMessage := NewMessageSender(chatID, p.tg)

	values := strings.Split(text, " ")
	if len(values) < 2 {
		return errors.New("Wrong description - note format")
	}

	note := &storage.Note{
		Description: values[0],
		Value:       strings.Join(values[1:len(values)], " "),
		UserName:    username,
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

func (p EventProcessor) removeNote(chatID int, description, username string) (err error) {
	defer func() { err = errors.WrapIfErr("can`t execute remove note command", err) }()

	if err = p.storage.RemoveNote(&storage.Note{
		Description: description,
		UserName:    username,
	}); err != nil {
		return p.tg.SendMessage(chatID, msgNotFound)
	}

	return p.tg.SendMessage(chatID, msgRemoved)

}

func (p EventProcessor) sendAll(chatID int, username string) (err error) {
	defer func() { err = errors.WrapIfErr("can't do command: can't send all notes", err) }()

	notes, err := p.storage.GetAllNotes(username)
	if err != nil && errors.Is(err, storage.ErrNoSavedNotes) {
		return p.tg.SendMessage(chatID, msgNoSavedNotes)
	}

	if err != nil {
		return err
	}

	sb := strings.Builder{}
	for i, note := range *notes {
		row := fmt.Sprintf("***%v.    [--- %s ---]***\n%s\n***%v.    [--- %s ---]***\n", i+1, note.Description, note.Value, i+1, note.Description)
		sb.WriteString(row + "\n\n\r")
	}

	//Telegram ignores last \n
	message := sb.String()

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
