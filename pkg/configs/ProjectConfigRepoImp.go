package configs

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type LocalRepo = projectConfigRepo
type ApolloRepo = projectConfigRepo

type projectConfigRepo struct {
	*viper.Viper
}

func (repo *projectConfigRepo) Find() *ProjectConfig {
	cfg := new(ProjectConfig)

	option := func(c *mapstructure.DecoderConfig) { c.TagName = "configs" }
	if err := repo.Viper.Unmarshal(cfg, option); err != nil {
		log.Fatal().Err(err).Msg("Unmarshal configs.ProjectConfig failed:")
		return nil
	}

	return cfg
}
