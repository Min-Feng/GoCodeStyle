package injection

import (
	"context"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"ddd/pkg/drivingAdapter/api"
	"ddd/pkg/drivingAdapter/api/operation"
	"ddd/pkg/infra/part"
	"ddd/pkg/technical/configs"
	"ddd/pkg/technical/logger"
)

func HTTPServer(NewConfig func() *configs.ProjectConfig) {
	newRouter := func(lc fx.Lifecycle, address string, logLevel logger.Level) *api.Router {
		router := api.NewRouter(address, logLevel)
		lc.Append(
			fx.Hook{
				OnStart: func(context.Context) error {
					go router.QuicklyStart()
					return nil
				},
				OnStop: nil,
			},
		)
		return router
	}

	ioc := fx.New(
		fx.Provide(
			NewConfig,
			part.NewMySQL,
			newRouter,
			func() *operation.DebugHandler {
				return new(operation.DebugHandler)
			},
		),

		fx.Invoke(api.RegisterHandler),
	)
	ioc.Run()
}

func NewConfig() *configs.ProjectConfig {
	var repo configs.ProjectConfigRepo

	src := strings.ToLower(os.Getenv("CONF_SRC"))
	switch src {
	case "local":
		fileName := os.Getenv("FILE_NAME")
		repo = configs.NewLocalRepo(fileName)
	case "apollo":
		ip := os.Getenv("APOLLO_ADDRESS")
		repo = configs.NewApolloRepo(ip)
	default:
		log.Fatal().Str("CONF_SRC", src).Msg("Unexpected environment variable:")
	}

	//noinspection GoNilness
	cfg := repo.Find()
	log.Info().Msg("New Project Config successfully")
	return cfg
}
