package events

const msgHelp = `I can save and keep you notes. Also I can offer you them to read.

In order to save the note, just send me text/link/etc to it.

In order to get all notes, send me command /show.`

const msgHello = "Hi there! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command 🤔"
	msgNoSavedNotes   = "You have no saved notes 🙊"
	msgSaved          = "Saved! 👌"
	msgAlreadyExists  = "You have already have this note in your list 🤗"
)
