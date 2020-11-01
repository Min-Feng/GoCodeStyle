package basic

type Message interface {
	Kind() MessageKind
}

type MessageKind string

const (
	Event   MessageKind = "event"
	Command MessageKind = "command"
)
