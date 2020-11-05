package configs

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func NewConfig() *ProjectConfig {
	var repo ProjectConfigRepoQ

	src := strings.ToLower(os.Getenv("CONF_SRC"))
	switch src {
	case "local":
		fileName := os.Getenv("FILE_NAME")
		repo = NewLocalRepo(fileName)
	case "apollo":
		ip := os.Getenv("APOLLO_ADDRESS")
		repo = NewApolloRepo(ip)
	default:
		log.Fatal().Str("CONF_SRC", src).Msg("Unexpected environment variable:")
	}

	//noinspection GoNilness
	cfg := repo.QueryConfig()
	log.Info().Msg("New Project Config successfully")
	return cfg
}
