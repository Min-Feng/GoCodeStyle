package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/shima-park/agollo"
	remote "github.com/shima-park/agollo/viper-remote"
	"github.com/spf13/viper"
)

const (
	ApolloAppID      = "hello"
	ApolloCluster    = "default"
	ApolloNamespaces = "application"
)

func NewApolloProjectConfigRepo(remoteAddress string) ProjectConfigRepo {
	remote.SetAppID(ApolloAppID)
	remote.SetConfigType("prop")
	remote.SetAgolloOptions(
		agollo.Cluster(ApolloCluster),
		agollo.DefaultNamespace(ApolloNamespaces),
	)

	vp := viper.New()
	vp.SetConfigType("prop")

	if err := vp.AddRemoteProvider("apollo", remoteAddress, ApolloNamespaces); err != nil {
		log.Fatal().
			Err(err).
			Str("ApolloRemoteAddress", remoteAddress).
			Msg("Viper add remote provider failed:")
	}

	if err := vp.ReadRemoteConfig(); err != nil {
		log.Fatal().Err(err).Msg("Reading apollo remote config failed:")
	}

	log.Info().Msg("New remote config repository from apollo successfully")
	return &ApolloProjectConfigRepo{vp}
}
