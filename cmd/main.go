package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"

	"ddd/pkg/configs"
	"ddd/pkg/infra/loghelper"
)

func main() {
	cfg := NewProjectConfig()
	loghelper.Init(cfg.LogLevel, loghelper.HumanType)
	c := spew.Sdump(cfg)
	fmt.Println(c)
}

func NewProjectConfig() *configs.ProjectConfig {
	loghelper.Init(loghelper.DefaultLevel, loghelper.HumanType)
	configFileName := strings.ToLower(os.Getenv("RUN_ENV"))
	if configFileName == "" {
		log.Fatal().Msg("Not found environment variable 'RUN_ENV'")
	}

	store := configs.NewLocalProjectConfigStore(configFileName, "./config", "../config")
	// vp := viper.NewLocalProjectConfigStore("./config", configFileName)
	return store.Find()
}
