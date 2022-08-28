package storage

type FileStorage struct {
	basePath string
}

func (f FileStorage) SaveNote(note *Note) error {
	//TODO implement me
	panic("implement me")
}

func (f FileStorage) GetAllNotes(userName string) (*[]Note, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileStorage) GetNoteById(userName string, id int) (*Note, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileStorage) GetNotesCount() int {
	//TODO implement me
	panic("implement me")
}

func (f FileStorage) RemoveNote(id int) error {
	//TODO implement me
	panic("implement me")
}

func (f FileStorage) IsNoteExists(note *Note) {
	//TODO implement me
	panic("implement me")
}

func New(basePath string) Storage {
	return FileStorage{basePath: basePath}
}
