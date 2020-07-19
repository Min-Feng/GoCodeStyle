package loghelper

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Level = zerolog.Level

//noinspection GoUnusedConst
const (
	DebugLevel   = zerolog.DebugLevel // 0
	InfoLevel    = zerolog.InfoLevel  // 1
	DefaultLevel = zerolog.InfoLevel  // 1
	ErrorLevel   = zerolog.ErrorLevel // 3, for unit test
)

type WriterType string

//noinspection GoUnusedConst
const (
	JSONType  WriterType = "json"
	HumanType WriterType = "human"
)

// Init can set global logLevel = [0, 1, 3] = [Debug, Info, Error]
// wType = ["json", "human"]
func Init(logLevel Level, wType WriterType) {
	zerolog.SetGlobalLevel(logLevel)

	switch wType {
	case HumanType:
		writer := newConsoleWriter()
		log.Logger = log.Output(writer).With().Caller().Logger()
	case JSONType:
		log.Logger = log.Output(os.Stdout)
	default:
		panic("Init log package failed")
	}
}

func newConsoleWriter() io.Writer {
	writer := &zerolog.ConsoleWriter{
		Out:         os.Stdout,
		NoColor:     true,
		TimeFormat:  time.RFC3339,
		FormatLevel: func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("[%v]", i)) },
		FormatCaller: func(i interface{}) string {
			filePath := i.(string)
			if len(filePath) == 0 {
				return filePath
			}

			moduleDirectory := []string{"GoProjectLayout/"}
			for _, dir := range moduleDirectory {
				if strings.Contains(filePath, dir) {
					path := strings.Split(filePath, dir)
					filePath = path[len(path)-1]
					break
				}
			}

			// path := strings.SplitAfter(filePath, "/")
			// n := len(path)
			// b := new(strings.Builder)
			// b.WriteString(path[n-2])
			// b.WriteString(path[n-1])
			// filePath = b.String()

			return filePath
		},
	}
	return writer
}
