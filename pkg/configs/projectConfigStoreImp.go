package configs

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type LocalProjectConfigStore = projectConfigStore
type ApolloProjectConfigStore = projectConfigStore

type projectConfigStore struct {
	*viper.Viper
}

func (store projectConfigStore) Find() *ProjectConfig {
	cfg := new(ProjectConfig)

	option := func(c *mapstructure.DecoderConfig) { c.TagName = "configs" }
	err := store.Viper.Unmarshal(cfg, option)
	if err != nil {
		log.Fatal().Msg("Unmarshal ProjectConfig failed: " + err.Error())
		return nil
	}

	return cfg
}
