package viper

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Viper = viper.Viper

func New(configPath string, configFileName string) *Viper {
	if configPath == "" || configFileName == "" {
		log.Fatal().Msg("Not found: configPath or configFileName is empty string")
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)
	v.SetConfigName(configFileName)

	if err := v.ReadInConfig(); err != nil {
		log.Fatal().Msg("Reading config: " + err.Error())
	}

	log.Info().Msg("Using config: " + v.ConfigFileUsed())
	return v
}
