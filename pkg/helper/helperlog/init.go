package helperlog

//nolint:gochecknoglobals
var ModuleDirectory []string

func Init() {
	ModuleDirectory = []string{"GoCodeStyle/"}
	DefaultLevel := InfoLevel
	SetGlobal(DefaultLevel, WriterKindHuman)
}

func init() {
	Init()
	UnitTestMode()
}
