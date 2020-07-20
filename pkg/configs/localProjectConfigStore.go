package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func NewLocalProjectConfigStore(configFileName string, configPath ...string) ProjectConfigStore {
	if configPath == nil || configFileName == "" {
		log.Fatal().Msg("Not found: configPath or configFileName is empty string")
	}

	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName(configFileName)
	for i := range configPath {
		vp.AddConfigPath(configPath[i])
		if err := vp.ReadInConfig(); err == nil {
			break
		}
	}
	if err := vp.ReadInConfig(); err != nil {
		log.Fatal().Msg("Reading config: " + err.Error())
	}

	log.Info().Msg("Using config: " + vp.ConfigFileUsed())
	return &LocalProjectConfigStore{vp}
}
