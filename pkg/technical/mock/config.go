package mock

import "ddd/pkg/technical/configs"

var Config = configs.NewLocalRepo("dev").QueryConfig()

func NewConfig() *configs.ProjectConfig {
	return configs.NewLocalRepo("dev").QueryConfig()
}
