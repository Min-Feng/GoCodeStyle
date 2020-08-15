package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"

	"ddd/pkg/configs"
	"ddd/pkg/loghelper"
)

// 有 init 表示存在全域變數, 非開發 lib 的專案, 全域變數不是好選擇
// 如果要使用 init 盡量不使用隱式 init, 而是顯式 Init
// 且集中放在程式的入口處, 通常會是 cmd/main.go
// 明顯設置 Init 可以讓其他維護者知道要注意全域
//
// 由於還沒讀取到 config 的設定
// 不知道 log level, 所以先使用一個預設等級
func init() {
	DefaultLevel := loghelper.InfoLevel
	loghelper.Init(DefaultLevel, loghelper.WriterKindHuman)
}

func main() {
	configSRC := strings.ToLower(os.Getenv("CONF_SRC"))
	cfg := NewConfig(configSRC)
	loghelper.Init(cfg.LogLevel, loghelper.WriterKindHuman)
	c := spew.Sdump(cfg)
	fmt.Println(c)
}

func NewConfig(src string) *configs.ProjectConfig {
	var store configs.ProjectConfigRepo

	switch src {
	case "local":
		fileName := os.Getenv("FILE_NAME")
		// 因為不確定 開發者會在什麼地方執行 go run, 專案根目錄 或 cmd 目錄
		store = configs.NewLocalProjectConfigRepo(fileName, "./config", "../config")
	case "apollo":
		ip := os.Getenv("APOLLO_ADDRESS")
		store = configs.NewApolloProjectConfigRepo(ip)
	default:
		log.Fatal().Str("CONF_SRC", src).Msg("Unexpected environment variable:")
	}

	//noinspection GoNilness
	return store.Find()
}
