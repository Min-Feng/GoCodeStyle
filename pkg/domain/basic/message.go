package basic

import "context"

type MessageKind int

const (
	Command MessageKind = -1
	Event   MessageKind = 1
)

type MessageName int

const (
	CreateMemberCommand MessageName = -(iota + 1)
)

const (
	CreatedMemberEvent MessageName = (iota + 1)
)

type Message interface {
	ID() string
	Kind() MessageKind
	Name() MessageName
}

type MessageRegistry struct {
	Command map[MessageName]MessageHandleFunc
	Event   map[MessageName][]MessageHandleFunc
}

type MessageHandleFunc func(ctx context.Context, msg Message)
