package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// configPath 設置多個, 是因為二進位執行檔, 可能在 go module 根目錄執行 或 cmd 目錄執行
func NewLocalRepo(configFileName string, configPaths ...string) ProjectConfigRepo {
	if configPaths == nil || configFileName == "" {
		log.Fatal().Msg("Not found: configFileName is empty string or configPaths is nil slice")
	}

	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName(configFileName)
	for _, path := range configPaths {
		vp.AddConfigPath(path)
		if err := vp.ReadInConfig(); err == nil {
			break
		}
	}
	if err := vp.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Reading config:")
	}

	log.Info().Msgf("New local config repository from %v successfully", vp.ConfigFileUsed())
	return &LocalRepo{vp}
}
