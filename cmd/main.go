package main

import (
	"fmt"
	"os"
	"strings"

	"ddd/pkg/infra/loghelper"
	"ddd/pkg/repository/viper"
	"ddd/pkg/unknown"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := NewProjectConfig()
	loghelper.Init(cfg.LogLevel, loghelper.HumanType)
	c := spew.Sdump(cfg)
	fmt.Println(c)
}

func NewProjectConfig() *unknown.ProjectConfig {
	loghelper.Init(loghelper.InfoLevel, loghelper.HumanType)
	configFileName := strings.ToLower(os.Getenv("RUN_ENV"))
	if configFileName == "" {
		log.Fatal().Msg("Not found environment variable 'RUN_ENV'")
	}

	vp := viper.New("./config", configFileName)
	// vp := viper.New("./config", configFileName)
	store := viper.NewProjectConfigStore(vp)
	return store.Find()
}
