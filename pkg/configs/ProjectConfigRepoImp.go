package configs

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type LocalProjectConfigRepo = projectConfigStore
type ApolloProjectConfigRepo = projectConfigStore

type projectConfigStore struct {
	*viper.Viper
}

func (repo *projectConfigStore) Find() *ProjectConfig {
	cfg := new(ProjectConfig)

	option := func(c *mapstructure.DecoderConfig) { c.TagName = "configs" }
	if err := repo.Viper.Unmarshal(cfg, option); err != nil {
		log.Fatal().Err(err).Msg("Unmarshal ProjectConfig failed:")
		return nil
	}

	log.Info().Msg("ProjectConfigRepo Find config successfully")
	return cfg
}
