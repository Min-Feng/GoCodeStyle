package helperlog

//nolint:gochecknoglobals
var ModuleDirectory []string

func Init() {
	ModuleDirectory = []string{"GoCodeStyle/"}
}

func init() {
	Init()
	UnitTestMode()
}
