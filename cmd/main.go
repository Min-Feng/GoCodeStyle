package main

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"ddd/pkg/api"
	"ddd/pkg/configs"
	"ddd/pkg/helper/helperlog"
	"ddd/pkg/repository/mysql"
)

// 有 init 表示存在全域變數, 非 lib 類型的專案, 全域變數不是好選擇
// 如果要使用 init 盡量不使用隱式 init, 而是顯式 Init
// 且集中放在程式的入口處, 通常會是 cmd/main.go
// 明顯設置 Init 可以讓其他維護者知道要注意全域
//
// 由於還沒讀取到 config 的設定
// 不知道 log level, 所以先使用一個預設等級
func init() {
	helperlog.Init()
	configs.Init()
}

func main() {
	log.Info().Msg("MainStart")

	configSRC := strings.ToLower(os.Getenv("CONF_SRC"))
	cfg := NewConfig(configSRC)
	helperlog.SetGlobal(cfg.LogLevel, helperlog.WriterKindHuman)

	mysql.NewDB(&cfg.MySQL)

	debugHandler := &api.DebugHandler{}
	router := api.NewRouter(cfg.LogLevel)
	router.PUT("debug/logLevel", debugHandler.UpdateLogLevel)

	if err := router.Run(":" + cfg.Port); err != nil {

	}
}

func NewConfig(src string) *configs.ProjectConfig {
	var repo configs.ProjectConfigRepo

	switch src {
	case "local":
		fileName := os.Getenv("FILE_NAME")
		repo = configs.NewLocalRepo(fileName)
	case "apollo":
		ip := os.Getenv("APOLLO_ADDRESS")
		repo = configs.NewApolloRepo(ip)
	default:
		log.Fatal().Str("CONF_SRC", src).Msg("Unexpected environment variable:")
	}

	//noinspection GoNilness
	cfg := repo.Find()
	log.Info().Msg("New Project Config successfully")
	return cfg
}
