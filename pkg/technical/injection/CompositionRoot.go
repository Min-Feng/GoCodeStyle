package injection

import (
	"context"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"go.uber.org/fx"

	"ddd/pkg/drivingAdapter/api/operation"
	api "ddd/pkg/drivingAdapter/api/shared"
	"ddd/pkg/infra/part"
	"ddd/pkg/technical/configs"
	"ddd/pkg/technical/logger"
)

func newRouter(lc fx.Lifecycle, logLevel logger.Level) *api.Router {
	router := api.NewRouter(logLevel)
	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go router.Run(":" + "8080")
				return nil
			},
			OnStop: nil,
		},
	)
	return &router
}

func HTTPServer(NewConfig func() *configs.ProjectConfig) {

	ioc := fx.New(
		fx.Provide(
			NewConfig,
			part.NewMySQL,
			newRouter,
			func() *operation.DebugHandler {
				return new(operation.DebugHandler)
			},
		),

		fx.Invoke(RegisterHTTPHandler),
	)
	ioc.Run()
}

func NewConfig() *configs.ProjectConfig {
	var repo configs.ProjectConfigRepoQ

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
	cfg := repo.QueryConfig()
	log.Info().Msg("New Project Config successfully")
	return cfg
}
