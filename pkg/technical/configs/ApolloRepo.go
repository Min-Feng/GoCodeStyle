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

func NewApolloRepo(remoteAddress string) ProjectConfigRepoQ {
	return newApolloRepo(remoteAddress, ApolloAppID, ApolloCluster, ApolloNamespaces)
}

func newApolloRepo(remoteAddress string, appID string, clusterName string, namespace string) ProjectConfigRepoQ {
	remote.SetAppID(appID)
	remote.SetConfigType("prop")
	remote.SetAgolloOptions(
		agollo.Cluster(clusterName),
		agollo.DefaultNamespace(namespace),
	)

	vp := viper.New()
	vp.SetConfigType("prop")

	if err := vp.AddRemoteProvider("apollo", remoteAddress, namespace); err != nil {
		log.Fatal().
			Err(err).
			Str("ApolloRemoteAddress", remoteAddress).
			Msg("Viper add remote provider failed:")
	}

	if err := vp.ReadRemoteConfig(); err != nil {
		log.Fatal().Err(err).Msg("Reading apollo remote config failed:")
	}

	log.Info().Msg("New remote config repository from apollo successfully")
	return &ApolloRepo{vp}
}
