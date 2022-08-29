package storage

import (
	"encoding/gob"
	"fmt"
	errors "github.com/slbmax/telegram-pocketbook-bot/lib/custom-errors"
	"os"
	"path/filepath"
)

type FileStorage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return FileStorage{basePath: basePath}
}

func (f FileStorage) SaveNote(note *Note) (err error) {
	defer func() { err = errors.WrapIfErr("can't save note", err) }()

	filePath := filepath.Join(f.basePath, note.UserName)

	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}

	fileName, err := fileName(note)
	if err != nil {
		return err
	}

	filePath = filepath.Join(filePath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	if err = gob.NewEncoder(file).Encode(note); err != nil {
		return nil
	}

	return nil
}

func (f FileStorage) GetAllNotes(userName string) (notes *[]Note, err error) {
	defer func() { err = errors.WrapIfErr("can't get notes", err) }()

	filePath := filepath.Join(f.basePath, userName)

	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, ErrNoSavedNotes
	}

	result := make([]Note, len(files))

	for i, file := range files {
		note, err := f.decodeNote(filepath.Join(filePath, file.Name()))
		if err != nil {
			return nil, err
		}

		result[i] = note
	}

	return &result, nil
}

func (f FileStorage) RemoveNote(note *Note) error {
	fileName, err := fileName(note)
	if err != nil {
		return errors.Wrap("can't remove file", err)
	}

	path := filepath.Join(f.basePath, note.UserName, fileName)

	if err := os.Remove(path); err != nil {
		msg := fmt.Sprintf("can't remove file %s", path)
		return errors.Wrap(msg, err)
	}

	return nil
}

func (f FileStorage) IsNoteExists(note *Note) (bool, error) {
	fileName, err := fileName(note)
	if err != nil {
		return false, errors.Wrap("can't check if file exists", err)
	}

	path := filepath.Join(f.basePath, note.UserName, fileName)

	switch _, err = os.Stat(path); {
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	case err != nil:
		msg := fmt.Sprintf("can't check if file %s exists", path)

		return false, errors.Wrap(msg, err)
	}

	return true, nil
}

func fileName(n *Note) (string, error) {
	return n.Hash()
}

func (f FileStorage) decodeNote(filePath string) (Note, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Note{}, errors.Wrap("can`t decode note", err)
	}

	defer file.Close()

	var note Note

	if err := gob.NewDecoder(file).Decode(&note); err != nil {
		return Note{}, errors.Wrap("can`t decode note", err)
	}

	return note, nil
}
