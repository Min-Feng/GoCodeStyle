package viper

import (
	"ddd/pkg/unknown"

	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
)

func NewProjectConfigStore(v *Viper) unknown.ProjectConfigStore {
	return projectConfigStore{v}
}

type projectConfigStore struct {
	*Viper
}

func (store projectConfigStore) Find() *unknown.ProjectConfig {
	cfg := new(unknown.ProjectConfig)

	option := func(c *mapstructure.DecoderConfig) {
		c.TagName = "config"
	}

	err := store.Viper.Unmarshal(cfg, option)
	if err != nil {
		log.Fatal().Msg("Unmarshal ProjectConfig failed: " + err.Error())
		return nil
	}

	return cfg
}
