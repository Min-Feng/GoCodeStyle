package loghelper

// 此套件的目的, 主要是針對 zerolog 的全域值進行控管
// 若有其他套件需要注入 logger
// 也可以在此輔助套件, 進行實現

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
	DebugLevel = zerolog.DebugLevel // 0
	InfoLevel  = zerolog.InfoLevel  // 1
	ErrorLevel = zerolog.ErrorLevel // 3, for unit test
)

type WriterType string

//noinspection GoUnusedConst
const (
	JSONType  WriterType = "json"
	HumanType WriterType = "human"
)

//noinspection GoUnusedExportedFunction
func GlobalLevel() Level {
	return zerolog.GlobalLevel()
}

// InitAtUnitTest 專門提供給單元測試時使用, 避免 go test 輸出時, 有多餘 log 訊息
func InitAtUnitTest() {
	Init(ErrorLevel, HumanType)
}

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
		Out:          os.Stdout,
		NoColor:      true,
		TimeFormat:   time.RFC3339,
		FormatLevel:  func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("[%v]", i)) },
		FormatCaller: longFileFormatCaller("GoProjectLayout/"),
	}
	return writer
}

// log 執行時, 輸出所在檔案 及 go module 為根目錄的相對路徑目錄
// example: pkg/configs/localProjectConfigStore.go:26
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
// example: configs/localProjectConfigStore.go:26
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
