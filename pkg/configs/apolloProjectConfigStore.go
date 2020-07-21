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
	ApolloNamespaces = "applications"
)

func NewApolloProjectConfigStore(remoteAddress string) ProjectConfigStore {
	remote.SetAppID(ApolloAppID)
	remote.SetConfigType("prop")
	remote.SetAgolloOptions(agollo.Cluster(ApolloCluster), agollo.PreloadNamespaces(ApolloNamespaces))

	vp := viper.New()
	vp.SetConfigType("prop")

	if err := vp.AddRemoteProvider("apollo", remoteAddress, ApolloNamespaces); err != nil {
		log.Fatal().
			Str("ApolloRemoteAddress", remoteAddress).
			Msgf("Viper add remote provider failed: %v:", err)
	}

	if err := vp.ReadRemoteConfig(); err != nil {
		log.Fatal().Msg("Reading config: " + err.Error())
	}

	return &ApolloProjectConfigStore{vp}
}
