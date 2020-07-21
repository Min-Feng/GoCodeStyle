package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"ddd/pkg/configs"
	"ddd/pkg/infra/loghelper"
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
	loghelper.Init(DefaultLevel, loghelper.HumanType)
}

func main() {
	configSRC := strings.ToLower(os.Getenv("CONF_SRC"))
	cfg := configs.NewProjectConfig(configSRC)
	loghelper.Init(cfg.LogLevel, loghelper.HumanType)
	c := spew.Sdump(cfg)
	fmt.Println(c)
}
