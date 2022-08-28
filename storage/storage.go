package storage

type Storage interface {
	SaveNote(note *Note) error
	GetAllNotes(userName string) (*[]Note, error)
	GetNoteById(userName string, id int) (*Note, error)
	GetNotesCount() int
	RemoveNote(id int) error
	IsNoteExists(note *Note)
}

type Note struct {
	Value    string
	UserName string
}
