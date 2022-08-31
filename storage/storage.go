package storage

import (
	"crypto/sha1"
	"fmt"
	errors "github.com/slbmax/telegram-pocketbook-bot/lib/custom-errors"
	"io"
)

type Storage interface {
	SaveNote(note *Note) error
	GetAllNotes(userName string) (*[]Note, error)
	RemoveNote(note *Note) error
	IsNoteExists(note *Note) (bool, error)
}

var ErrNoSavedNotes = errors.New("no saved notes")

type Note struct {
	Description string
	Value       string
	UserName    string
}

func (n Note) Hash() (string, error) {
	hash := sha1.New()

	if _, err := io.WriteString(hash, n.Description); err != nil {
		return "", errors.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(hash, n.UserName); err != nil {
		return "", errors.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
