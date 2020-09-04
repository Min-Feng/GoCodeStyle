package helperlog

type Level = string

//noinspection GoUnusedConst
const (
	TraceLevel Level = "trace"
	DebugLevel Level = "debug"
	InfoLevel  Level = "info"
	ErrorLevel Level = "error"
	Disabled   Level = "disabled"
)

type WriterKind = string

//noinspection GoUnusedConst
const (
	WriterKindJSON  WriterKind = "json"
	WriterKindHuman WriterKind = "human"
)
