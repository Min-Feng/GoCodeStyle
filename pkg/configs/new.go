package configs

import (
	"os"

	"github.com/rs/zerolog/log"
)

type ConfigSource = string

const (
	LOCAL  ConfigSource = "local"
	APOLLO ConfigSource = "apollo"
)

func NewProjectConfig(src ConfigSource) *ProjectConfig {
	var store ProjectConfigStore

	switch src {
	case LOCAL:
		store = NewLocalProjectConfigStore("dev", "./config", "../config")
	case APOLLO:
		ip := os.Getenv("APOLLO_ADDRESS")
		store = NewApolloProjectConfigStore(ip)
	default:
		log.Fatal().Str("CONF_SRC", src).Msg("Unexpected environment variable:")
	}

	//noinspection GoNilness
	return store.Find()
}
