package events

const (
	UnknownEvent Type = iota
	MessageEvent
)

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

type Event struct {
	Type     Type
	Text     string
	ChatID   int
	Username string
}
