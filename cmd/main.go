package main

import (
	"github.com/rs/zerolog/log"

	api2 "ddd/pkg/drivingAdapter/api/shared"
	"ddd/pkg/technical/configs"
	"ddd/pkg/technical/injection"
	"ddd/pkg/technical/logger"
)

// 有 init 表示存在全域變數, 非 lib 類型的專案, 全域變數不是好選擇
// 如果要使用 init 盡量不使用隱式 init, 而是顯式 Init
// 且集中放在程式的入口處, 通常會是 cmd/main.go
// 明顯設置 Init 可以讓其他維護者知道要注意全域
//
// 由於還沒讀取到 config 的設定
// 不知道 log level, 所以先使用一個預設等級
func init() {
	logger.Init()
	logger.DefaultMode()
	configs.Init()
	api2.Init()
}

func main() {
	log.Info().Msg("MainStart")

	// cfg := NewConfig()
	// logger.SetGlobal(cfg.LogLevel, logger.WriterKindHuman)
	//
	// part.NewMySQL(&cfg.MySQL)
	//
	// router := api.NewRouter(":"+cfg.Port, cfg.LogLevel)
	//
	// debugHandler := &operation.DebugHandler{}
	// router.RegisterAPIHandler(debugHandler)
	//
	// err := router.QuicklyStart()
	// if err != nil {
	//
	// }
	injection.HTTPServer(configs.NewConfig)
}
