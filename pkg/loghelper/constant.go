package loghelper

type Level string

//noinspection GoUnusedConst
const (
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	ErrorLevel Level = "error" // for unit test
)

type WriterKind string

//noinspection GoUnusedConst
const (
	WriterKindJSON  WriterKind = "json"
	WriterKindHuman WriterKind = "human"
)
