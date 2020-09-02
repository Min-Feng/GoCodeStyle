package helperlog

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//nolint:gochecknoglobals
var ModuleDirectory = []string{"GoCodeStyle/"}

// UnitTestSetting 提供 _test.go 使用, 避免執行命令 go test 輸出時, 有多餘 log 訊息
func UnitTestSetting() {
	Init(ErrorLevel, WriterKindHuman)
}

// DevelopSetting 開發期間, 進行單元測試, 查看 log 格式是否符合預期
//
//noinspection GoUnusedExportedFunction
func DevelopSetting() {
	Init(DebugLevel, WriterKindHuman)
}

// Init can set global logLevel = ["debug", "info", "error"]
// wKind = ["json", "human"]
func Init(logLevel Level, wKind WriterKind) {
	switch logLevel {
	case DebugLevel:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case InfoLevel:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case ErrorLevel:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case Disabled:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		panic("Init log package failed: invalid log level")
	}

	switch wKind {
	case WriterKindHuman:
		writer := newConsoleWriter()
		log.Logger = log.Output(writer).With().Caller().Logger()
	case WriterKindJSON:
		log.Logger = log.Output(os.Stdout)
	default:
		panic("Init log package failed: invalid writer kind")
	}
}

func newConsoleWriter() io.Writer {
	writer := &zerolog.ConsoleWriter{
		Out:          os.Stdout,
		NoColor:      true,
		TimeFormat:   "2006-01-02 15:04:05Z07:00",
		FormatLevel:  func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("[%v]", i)) },
		FormatCaller: longFileFormatCaller(ModuleDirectory...),
	}
	return writer
}

// log 執行時, 輸出所在檔案 及 go module 為根目錄的相對路徑目錄
// example: pkg/configs/LocalRepo.go:26
func longFileFormatCaller(moduleDirectory ...string) zerolog.Formatter {
	return func(i interface{}) string {
		filePath := i.(string)
		if len(filePath) == 0 {
			return filePath
		}

		for _, dir := range moduleDirectory {
			if strings.Contains(filePath, dir) {
				path := strings.Split(filePath, dir)
				filePath = path[len(path)-1]
				break
			}
		}

		return filePath
	}
}

// log 執行時, 輸出所在檔案 及 其目錄
// example: configs/LocalRepo.go:26
//
//nolint:deadcode
//noinspection GoUnusedFunction
func shortFileFormatCaller() zerolog.Formatter {
	return func(i interface{}) string {
		filePath := i.(string)
		if len(filePath) == 0 {
			return filePath
		}

		path := strings.SplitAfter(filePath, "/")
		n := len(path)
		b := new(strings.Builder)
		b.WriteString(path[n-2])
		b.WriteString(path[n-1])
		filePath = b.String()

		return filePath
	}
}
